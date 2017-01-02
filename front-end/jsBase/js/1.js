


function myfunc() {
	var x = document.getElementById('demo');
	x.innerHTML="Hello JavaScript"; 
	// body...
}

function changeIamge() {
	if (element.src.match("bulbon")) {
		element.src="/i/eg_bulboff.gif";
	}
	else
	{
		element.src="/i/eg_bulbon.gif";
	}
	// body...
}

function myfunc2() {
	var x = document.getElementById("demo2")
	x.style.color="#ff0000"
	// body...
}

function is_num() {
	var x = document.getElementById("demo3").value;
	if (x == "" || isNaN(x)) {
		alert("Not Numeric")
	}
	// body...
}