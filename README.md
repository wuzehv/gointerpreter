# gointerpreter

本项目用来存放 [《Writing An Interpreter In Go》](https://interpreterbook.com/) 中所有的代码

### 支持的语法
```
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);

// array
let myArray = [1, 2, 3, 4, 5];

// map
let thorsten = {"name": "Thorsten", "age": 28};

myArray[0] // => 1
thorsten["name"] // => "Thorsten"

let add = fn(a, b) { return a + b; };

let add = fn(a, b) { a + b; }; // implicit return

// recursive
let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

// higher order functions
let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};

twice(addTwo, 2); // => 6
```