# [用两个栈实现队列](https://www.nowcoder.com/practice/54275ddae22f475981afa2244dd448c6?tpId=13&tqId=11158&rp=1&ru=%2Fta%2Fcoding-interviews&qru=%2Fta%2Fcoding-interviews%2Fquestion-ranking&tab=answerKey)

用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。



### 解题思路

栈：先进后出
队列：先进先进

两个栈，第一个栈进，第二个栈出

push 进入第一个栈， pop 从第二个栈出，首先需要把第一个栈的元素进到第二个栈，然后再从第二个栈出

如果第二站还有元素，pop 就从第二个栈操作，如果为空从第一个栈进到第二个栈，然后从第二个栈出


### Accepted

运行时间：2ms，超过46.73%用Go提交的代码

占用内存：848KB，超过48.49%用Go提交的代码
