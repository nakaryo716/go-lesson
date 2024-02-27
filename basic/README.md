# Basic

## Package
Package contain some function, type, struct.  
In this case of Rust, this is same of crate.

## Public
In Rust code, to sort public or private, I must write `pub`.  
In Go code, the way of separete public, private is letter size.  

## function
In Go code we can write function down below
```go
func functionName(argument1 int{argument type}, argument2, int) int{return type} {
    ...
}
```
Inaddition, we can return maltipul results.Is looks like return tple of rust.
```go
func swap(text1 string, text2 string) (string, string) {
    return text2, text1
}
```
caution: the syntax of return is no need of () like Rust.

## Type conversions
To convert v to T, I can write down below
```go
var pi float = math.PI
// out put: 3.14159....

var intPi int = int(pi)
// out put: 3

```
It is same of rust code ```{value} as i32```