# Golang 切片与函数参数“陷阱”

线性结构是计算机最常用的数据结构之一。无论是数组（arrary）还是链表（list），在编程中不可或缺。golang也有数组，不同于别的语言，golang还提供了切片（slice）。切片比数组有更好的灵活性，具有某些动态特性。然而切片又不像动态语言的列表（Python list）。不明白切片的基本实现，写程序的时候容易掉“坑”里。

##slice参数
本来写一个堆排序，使用了golang的slice来做堆，可是发现在pop数据的时候，切片不改变。进而引发了golang函数切片的参数，是传值还是传引用呢？我们知道slice相比array是引用类型。那么直觉上告诉我们如果函数修改了参数的切片，那么外层的切片变量也会变啦。

```go
func main() {

    slice := []int{0, 1, 2, 3}
    fmt.Printf("slice: %v slice addr %p \n", slice, &slice)
    
    ret := changeSlice(slice)
    fmt.Printf("slice: %v ret: %v slice addr %p \n", slice, &slice, ret)
}

func changeSlice(s []int) []int {
    s[1] = 111
    return s
}
```

结果和假设的一样：

```
slice: [0 1 2 3], slice addr: 0xc4200660c0
slice: [0 111 2 3], ret: [0 111 2 3], slice addr: 0xc4200660c0 
```

`changeSlice`函数修改了切片，变量 slice也跟着修改了。可是如果轻易就下结论，切片参数是按照引用传递，那么下面的现象就需要一种说法了：

func changeSlice(s []int) []int {
    fmt.Printf("func: %p \n", &s)
    s[1] = 111
    return s
}
我们在函数中打出参数 s 的地址，可以看见这个地址和main函数中的slice竟然不是同一个。为了了解这个，我们需要了解golang中的slice基本实现。

##slice基本实现
Golang中的slice，是一个看似array却不是array的复合结构。切片顾名思义，就是数组切下来的一个片段。slice结构大致存储了三个部分，第一部分为指向底层数组的指`ptr`，其次是切片的大小`len`和切片的容量`cap`：

      +--------+
      |        |
      |  ptr   |+------------+-------+-----------+
      |        |                     |           |
      +--------+                     |           |
      |        |                     |           |
      |        |                     |           |
      |  len 3 |                     |           |
      |        |                     |           |
      +--------+                     v           v
      |        |             +-----+-----+-----+-----+----+
      |        |             |     |     |     |     |    |
      |  cap 5 |     [5]int  |  0  |  1  |  2  |  3  | 4  |
      |        |             +-----+-----+-----+-----+----+
      +--------+



       slice := arr[1:4]             arr := [5]int{0,1,2,3,4}
       
有一个数组`arr`是一个包含五个int类型的结构，它的切片slice只是从其取了 1到3这几个数字。我们同样可以再生成一个切片 `slice2 := arr[2:5]`, 所取的就是数组后面的连续块。他们共同使用arr作为底层的结构，可以看见共用了数字的第3，4个元素。修改其中任何一个，都能改变两个切片的值。

```go
func main() {

    arr := [5]int{0, 1, 2, 3, 4}
    fmt.Println(arr)

    slice := arr[1:4]
    slice2 := arr[2:5]

    fmt.Printf("arr %v, slice1 %v, slice2 %v, %p %p %p\n", arr, slice, slice2, &arr, &slice, &slice2)
    
    fmt.Printf("arr[2]%p slice[1] %p slice2[0]%p\n", &arr[2], &slice[1], &slice2[0])

    arr[2] = 2222

    fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)


    slice[1] = 1111

    fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)

}
```

输出的值为：

```
[0 1 2 3 4]
arr [0 1 2 3 4], slice1 [1 2 3], slice2 [2 3 4], 0xc42006e0c0 0xc4200660c0 0xc4200660e0
arr[2]0xc42006e0d0 slice[1] 0xc42006e0d0 slice2[0]0xc42006e0d0
arr [0 1 2222 3 4], slice1 [1 2222 3], slice2 [2222 3 4]
arr [0 1 1111 3 4], slice1 [1 1111 3], slice2 [1111 3 4]
```

由此可见，数组的切片，只是从数组上切一段数据下来，不同的切片，其实是共享这些底层的数据数据。不过这些切片本身是不一样的对象，其内存地址都不一样。

从数组中切一块下来形成切片很好理解，有时候我们用make函数创建切片，实际上golang会在底层创建一个匿名的数组。如果从新的slice再切，那么新创建的两个切片都共享这个底层的匿名数组。

```
func main() {

    slice := make([]int, 5)
    for i:=0; i<len(slice);i++{
        slice[i] = i
    }
    fmt.Printf("slice %v \n", slice)

    slice2 := slice[1:4]
    fmt.Printf("slice %v, slice2 %v \n", slice, slice2)

    slice[1] = 1111
    fmt.Printf("slice %v, slice2 %v \n", slice, slice2)
}
```
输出如下：

```
slice [0 1 2 3 4] 
slice [0 1 2 3 4], slice2 [1 2 3] 
slice [0 1111 2 3 4], slice2 [1111 2 3] 
```

##slice的复制
既然slice的创建依赖于数组，有时候新生成的slice会修改，但是又不想修改原来的切片或者数组。此时就需要针对原来的切片进行复制了。

```
func main() {

    slice := []int{0, 1, 2, 3, 4}

    slice2 := slice[1:4]

    slice3 := make([]int, len(slice2))

    for i, e := range slice2 {
        slice3[i] = e
    }

    fmt.Printf("slice %v, slice3 %v \n", slice, slice3)

    slice[1] = 1111

    fmt.Printf("slice %v, slice3 %v \n", slice, slice3)
}
```

输出：

```
slice [0 1 2 3 4], slice3 [1 2 3] 
slice [0 1111 2 3 4], slice3 [1 2 3] 
```

由此可见，新创建的slice3，不会因为slice和slice2的修改而改变slice3。复制很有用，因此golang实现了一个内建的函数`copy`， copy有两个参数，第一个参数是复制后的对象，第二个是复制前的数组切片对象。

```
func main() {

    slice := []int{0, 1, 2, 3, 4}
    slice2 := slice[1:4]

    slice4 := make([]int, len(slice2))
    
    copy(slice4, slice2)
    
    fmt.Printf("slice %v, slice4 %v \n", slice, slice4)
    slice[1] = 1111
    fmt.Printf("slice %v, slice4 %v \n", slice, slice4)
}
```

slice4是从slice2中copy生成，slice和slice4底层的匿名数组是不一样的。因此修改他们不会影响彼此。

##slice 追加
###append 简介
创建复制切片都是常用的操作，还有一个追加元素或者追加数组也是很常用的功能。golang提供了`append`函数用于给切片追加元素。append第一个参数为原切片，随后是一些可变参数，用于将要追加的元素或多个元素。

```go
func main() {

    slice := make([]int, 1, 2)
    slice[0] = 111

    fmt.Printf("slice %v, slice addr %p, len %d, cap %d \n", slice, &slice, len(slice), cap(slice))

    slice = append(slice, 222)
    fmt.Printf("slice %v, slice addr %p, len %d, cap %d \n", slice, &slice, len(slice), cap(slice))

    slice = append(slice, 333)
    fmt.Printf("slice %v, slice addr %p, len %d, cap %d \n", slice, &slice, len(slice), cap(slice))

}
```

输出结果为：

```
slice [111], slice addr 0xc4200660c0, len 1, cap 2 
slice [111 222], slice addr 0xc4200660c0, len 2, cap 2 
slice [111 222 333], slice addr 0xc4200660c0, len 3, cap 4 
```

###切片容量
无论数组还是切片，都有长度限制。也就是追加切片的时候，如果元素正好在切片的容量范围内，直接在尾部追加一个元素即可。如果超出了最大容量，再追加元素就需要针对底层的数组进行复制和扩容操作了。

这里有一个切片容量的概念，从数组中切数据，切片的容量应该是切片的最后一个数据，和数组剩下元素的大小，再加上现有切片的大小。

数组 [0, 1, 2, 3, 4] 中，数组有5个元素。如果切片 s = [1, 2, 3]，那么3在数组的索引为3，也就是数组还剩最后一个元素的大小，加上s已经有3个元素，因此最后s的容量为 1 + 3 = 4。如果切片是
s1 = [4]，4的索引再数组中是最大的了，数组空余的元素为0，那么s1的容量为 0 + 1 = 1。具体如下表：

|切片	|切片字面量	|数组剩下空间|	长度|	容量|
|------|------|--------|------|-----|
|s[1:3]|	[1 2] | 2|	2|	4|
|s[1:1]|	[]|    4 |	0|	4|
|s[4:4]|	[]|    1 |	0|	1|
|s[4:5]|	[4]|	0 |	1|	1|

尽管上面的第二个和第三个切片的长度一样，但是他们的容量不一样。容量与最终append的策略有关系。

###append简单实现
我们已经知道，切片都依赖底层的数组结构，即使是直接创建的切片，也会生成一个匿名的数组。使用append时候，本质上是针对底层依赖的数组进行操作。如果切片的容量大于长度，给切片追加元素其实是修改底层数中，切片元素后面的元素。如果容量满了，就不能在原来的数组上修改，而是要创建一个新的数组，当然golang是通过创建一个新的切片实现的，因为新切片必然也有一个新的数组，并且这个数组的长度是原来的2倍，使用动态规划算法的简单实现。

```go
func main() {

    arr := [3]int{0, 1, 2}

    slice := arr[1:2]

    fmt.Printf("arr %v len %d, slice %v  len %d, cap %d, \n", arr, len(arr), slice, len(slice), cap(slice))

    slice[0] = 333

    fmt.Printf("arr %v len %d, slice %v  len %d, cap %d, \n", arr, len(arr), slice, len(slice), cap(slice))

    slice = append(slice, 4444)

    fmt.Printf("arr %v len %d, slice %v  len %d, cap %d, \n", arr, len(arr), slice, len(slice), cap(slice))

    slice = append(slice, 5555)

    fmt.Printf("arr %v len %d, slice %v  len %d, cap %d, \n", arr, len(arr), slice, len(slice), cap(slice))

    slice[0] = 333

    fmt.Printf("arr %v len %d, slice %v  len %d, cap %d, \n", arr, len(arr), slice, len(slice), cap(slice))
}
```
输出：

```
arr [0 1 2] len 3, slice [1]  len 1, cap 2, 
arr [0 333 2] len 3, slice [333]  len 1, cap 2, 
arr [0 333 444] len 3, slice [333 444]  len 2, cap 2, 
arr [0 333 444] len 3, slice [333 444 555]  len 3, cap 4, 
arr [0 333 444] len 3, slice [333 444 555]  len 3, cap 4, 
```

###小于容量的append
重输出，我们来画一下这个动态过程的图示：

```
+----+----+----+                           +----+----+----+                             +----+----+----+
       |    |    |    |                           |    |    |    |                             |    |    |    |
 arr   | 0  |  1 |  2 |                     arr   | 0  |333 | 2  |                       arr   | 0  |333 |444 |
       +----+----+----+                           +----+----+----+                             +----+----+----+
               ^                                          ^                                            ^    ^
               |                                          |                                            |    |
               |                                          |                                            |    |
               |                 slic0] = 333             |              slice = append(slice, 444)    +----+
               |               +----------------->        |                +---------------->          |
               |                                          |                                            |
            +--+--+----+----+                          +--+--+----+----+                            +--+--+----+----+
            |     |    |    |                          |     |    |    |                            |     |    |    |
            | p   | 1  | 2  |                          | p   | 1  | 2  |                            | p   | 2  | 2  |
            +-----+----+----+                          +-----+----+----+                            +-----+----+----+

            slice :=arr[1:2]                           slice :=arr[1:2]                             slice :=arr[1:2]

```


            

arr 是一个含有三个元素的数组，slice从arr中切了一个元素，由于切片的最后一个元素1是数组的索引是1，距离数组的最大长度还是1，因此slice的容量为2。当修改slice的第一个元素，由于slice底层是arr数组，因此arr的第二个元素也相应被修改。使用append方法给slice追加元素的时候，由于slice的容量还未满，因此等同于扩展了slice指向数组的内容，可以理解为重新切了一个数组内容附给slice，同时修改了数组的内容。

超出容量的append
如果接着append一个元素，那么数组肯定越界。此时append的原理大致如下：

创建一个新的临时切片t，t的长度和slice切片的长度一样，但是t的容量是slice切片的2倍，一个动态规划的方式。新建切片的时候，底层也创建了一个匿名的数组，数组的长度和切片容量一样。
复制s里面的元素到t里，即填入匿名数组中。然后把t赋值给slice，现在slice的指向了底层的匿名数组。
转变成小于容量的append方法。

```
       +----+----+----+                                +----+----+----+----+----+----+
       |    |    |    |                                |    |    |    |    |    |    |
 arr   | 0  |333 |444 |                                | 333| 444|    |    |    |    |
       +----+----+----+                                +----+----+----+----+----+----+
               ^    ^                                    ^     ^
               |    |                                    |     |
               |    |                                    +-----+
               +----+           +--------------->        |
               |                                         |
               |                                         +
            +--+--+----+----+                          +-----+-----+-----+
            |     |    |    |                          |     |     |     |
            | p   | 2  | 2  |                          | p   |  2  |  6  |
            +-----+----+----+                          +-----+-----+-----+

            slice :=arr[1:2]                          t := make([]int, len=2, cap=6)

                                                                  +
                                                                  |
                                                                  |
                                                                  |
                                                                  |
                                                                  v


                                                       +----+----+----+----+----+----+
                                                       |    |    |    |    |    |    |
                                                       | 333| 444|555 |    |    |    |
                                                       +----+----+----+----+----+----+
                                                         ^         ^
                                                         |         |
                                                         +----+----+
                                                         |
                                                         |
                                                         +
                                                       +-----+-----+-----+
                                                       |     |     |     |
                                                       | p   |  3  |  6  |
                                                       +-----+-----+-----+

                                                        slice = t
```                                                        

 
上面的图示描述了大于容量的时候append的操作原理。新生成的切片其依赖的数组和原来的数组就没有关系了，因此在修改新的切片元素，旧的数组也不会有关系。至于临时的切片t，将会被golang的gc回收。当然arr或它衍生的切片都没有应用的时候，也会被gc所回收。

>slice和array的关系十分密切，通过两者的合理构建，既能实现动态灵活的线性结构，也能提供访问元素的高效性能。当然，这种结构也不是完美无暇，共用底层数组，在部分修改操作的时候，可能带来副作用，同时如果一个很大的数组，那怕只有一个元素被切片应用，那么剩下的数组都不会被垃圾回收，这往往也会带来额外的问题。

###作为函数参数的切片
####直接改变切片
回到最开始的问题，当函数的参数是切片的时候，到底是传值还是传引用？从changeSlice函数中打出的参数s的地址，可以看出肯定不是传引用，毕竟引用都是一个地址才对。然而changeSlice函数内改变了s的值，也改变了原始变量slice的值，这个看起来像引用的现象，实际上正是我们前面讨论的切片共享底层数组的实现。

即切片传递的时候，传的是数组的值，等效于从原始切片中再切了一次。原始切片slice和参数s切片的底层数组是一样的。因此修改函数内的切片，也就修改了数组。


                                                +-----+----+-----+
                                                |     |    |     |
                 +-----------------------------+| p   |  3 |  3  |
                 |          +                   +-----+----+-----+
                 |          |
                 |          |                     s
                 |          |
                 |          |
                 v          v
               +----+----+-----+
               |    |    |     |
       arr     | 0  |  1 |  2  |
               +----+----+-----+
                 ^           ^
                 |           |
                 |           |
                 +-----------+
                 |
                 |
               +-+--+----+-----+
               |    |    |     |
               |  p |  3 |  3  |
               +----+----+-----+


                  slice
                  
                  
例如下面的代码：

```go
    slice := make([]int, 2, 3)
    for i := 0; i < len(slice); i++ {
        slice[i] = i
    }

    fmt.Printf("slice %v %p \n", slice, &slice)

    ret := changeSlice(slice)
    fmt.Printf("slice %v %p, ret %v \n", slice, &slice, ret)

    ret[1] = 1111

    fmt.Printf("slice %v %p, ret %v \n", slice, &slice, ret)
}

func changeSlice(s []int) []int {
    fmt.Printf("func s %v %p \n", s, &s)
    s = append(s, 3)
    return s
}
```


输出：

```
slice [0 1] 0xc42000a1e0 
func s [0 1] 0xc42000a260 
slice [0 1] 0xc42000a1e0, ret [0 1 3] 
slice [0 1111] 0xc42000a1e0, ret [0 1111 3] 
```

从输出可以看出，当slice传递给函数的时候，新建了切片s。在函数中给s进行了append一个元素，由于此时s的容量足够到，并没有生成新的底层数组。当修改返回的ret的时候，ret也共用了底层的数组，因此修改ret的原始，相应的也看到了slice的改变。

####append 操作
如果在函数内，append操作超过了原始切片的容量，将会有一个新建底层数组的过程，那么此时再修改函数返回切片，应该不会再影响原始切片。例如下面代码：

```go
func main() {
   slice := make([]int, 2, 2)
   for i := 0; i < len(slice); i++ {
       slice[i] = i
   }

   fmt.Printf("slice %v %p \n", slice, &slice)

   ret := changeSlice(slice)
   fmt.Printf("slice %v %p, ret %v \n", slice, &slice, ret)

   ret[1] = -1111

   fmt.Printf("slice %v %p, ret %v \n", slice, &slice, ret)
}

func changeSlice(s []int) []int {
   fmt.Printf("func s %v %p \n", s, &s)
   s[0] = -1
   s = append(s, 3)
   s[1] =  1111
   return s
}
```

输出：

```
slice [0 1] 0xc42000a1a0 
func s [0 1] 0xc42000a200 
slice [-1 1] 0xc42000a1a0, ret [-1 1111 3] 
slice [-1 1] 0xc42000a1a0, ret [-1 -1111 3] 
```

从输出可以很清楚的看到了我们的猜想。 即函数中先改变s第一个元素的值，由于slice和s都共用了底层数组，因此无论原始切片slice还是ret，第一个元素都是-1.然后append操作之后，因为超出了s的容量，因此会新建底层数组，虽然s变量没变，但是他的底层数组变了，此时修改s第一个元素，并不会影响原始的slice切片。也就是slice[1]还是1，而ret[1]则是-1。最后在外面修改ret[1]为 -1111，也不会影响原始的切片slice。

通过上面的分析，我们大致可以下结论，slice或者array作为函数参数传递的时候，本质是传值而不是传引用。传值的过程复制一个新的切片，这个切片也指向原始变量的底层数组。（个人感觉称之为传切片可能比传值的表述更准确）。函数中无论是直接修改切片，还是append创建新的切片，都是基于共享切片底层数组的情况作为基础。也就是最外面的原始切片是否改变，取决于函数内的操作和切片本身容量。

####传引用方式
array和slice作为参数传递的过程基本上是一样的，即传递他们切片。有时候我们需要处理传递引用的形式。golang提供了指针很方便实现类似的功能。

```
func main() {
    slice := []int{0, 1}
    fmt.Printf("slice %v %p \n", slice, &slice)

    changeSlice(&slice)
    fmt.Printf("slice %v %p \n", slice, &slice)

    slice[1] = -1111

    fmt.Printf("slice %v %p \n", slice, &slice)
}

func changeSlice(s *[]int) {
    fmt.Printf("func s %v %p \n", *s, s)
    (*s)[0] = -1
    *s = append(*s, 3)
    (*s)[1] =  1111
}
```

输出如下：

```
slice [0 1] 0xc42000a1e0 
func s [0 1] 0xc42000a1e0 
slice [-1 1111 3] 0xc42000a1e0 
slice [-1 -1111 3] 0xc42000a1e0 

```

从输出可以看到，传递给函数的是slice的指针，函数内对对s的操作本质上都是对slice的操作。并且也可以从函数内打出的s地址看到，至始至终就只有一个切片。虽然在append过程中会出现临时的切片或数组。

##总结
golang提供了array和slice两种序列结构。其中array是值类型。slice则是复合类型。slice是基于array实现的。slice的第一个内容为指向数组的指针，然后是其长度和容量。通过array的切片可以切出slice，也可以使用make创建slice，此时golang会生成一个匿名的数组。

因为slice依赖其底层的array，修改slice本质是修改array，而array又是有大小限制，当超过slice的容量，即数组越界的时候，需要通过动态规划的方式创建一个新的数组块。把原有的数据复制到新数组，这个新的array则为slice新的底层依赖。

数组还是切片，在函数中传递的不是引用，是另外一种值类型，即通过原始变量进行切片传入。函数内的操作即对切片的修改操作了。当然，如果为了修改原始变量，可以指定参数的类型为指针类型。传递的就是slice的内存地址。函数内的操作都是根据内存地址找到变量本身


>作者：人世间   
>链接：https://www.jianshu.com/p/7dbce907767a   
>來源：简书
