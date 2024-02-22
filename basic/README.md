# Basic

## Package
Package contain some function, type, struct.  
In this case of Rust, this is same of crate.

## Public
In Rust code, to sort public or private, I must write `pub`.  
In Go code, the way of separete public, private is letter size.  

## function
In Go code we can write function down below
```
func functionName(argument1 int{argument type}, argument2, int) int{return type} {
    ...
}
```
Inaddition, we can return maltipul results.Is looks like return tple of rust.
```
func swap(text1 string, text2 string) (string, string) {
    return text2, text1
}
```
caution: the syntax of return is no need of () like Rust.