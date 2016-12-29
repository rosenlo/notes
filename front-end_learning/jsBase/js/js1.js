// alert('Jason Luo');


/*
Foo()
Bar()
name = 'Rosen'


function Foo() {
	var name = 'Jason';
	console.log(name);
}


function Bar() {
	console.log(name);
}

*/

/*
(function (name) {
	console.log(name);
	// body...
})('Jason')

*/

/*
var array = [1,2,3,4];

array.push('Jason');
console.log(array)
array.unshift('test')
console.log(array)
array.splice(1,0,'ok');
console.log(array)
*/
array = [1,2,3,4]

for (var item in array) {
	console.log(item);
}

for (var i=0; i<array.length;i++) {
	console.log(i);
}