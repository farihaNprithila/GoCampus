# Understanding Scoping in Go: Package Level, Function Level, and Block Level

In Go, variables have different scopes depending on where they are declared. This document explains **package-level**, **function-level**, and **block-level** scoping in Go.

## 1. **Package-Level Scope**

A package-level variable is declared outside of any function and is accessible throughout the package. It can be used across different functions and is available for the entire package's scope.

```go
// Package-level variable
var y = "hello scope from package"
```
## 2. **Function-Level Scope**

A function-level variable is declared within a function and is only accessible within that function. The scope of the variable is limited to the function, meaning it cannot be accessed outside the function.

### Example of Function-Level Scope:

```go
// Exported function (accessible outside the package)
func PrintHelloScope() {
    // Initializing a variable at the function level
    var x = "hello scope"

    {
        // Inside this block, we're modifying the variable x
        x = "new hello scope"
        fmt.Println(x) // Output: new hello scope
    }

    // The change inside the block is reflected here because x was reassigned in the same scope
    fmt.Println(x) // Output: new hello scope

    // Accessing the package-level variable 'y'
    fmt.Println(y) // Output: hello scope from package
}
```
### Explanation:

- The variable `x` is declared using `var x = "hello scope"`, which makes it available only within the `PrintHelloScope` function.
- Inside the block, `x` is reassigned with a new value, and this change is reflected throughout the function because both the block and the reassignment happen within the same function scope.
- The function also prints the package-level variable `y`, which is accessible in the function because it is declared outside the function and within the same package.

### Key Points:
- **Function-level variables** are accessible only within the function where they are declared.
- Any change to a variable within a function (even inside a block) will affect the variable throughout the function.
- **Block scoping** does not introduce a new function-level scope; it only creates a nested scope for variables within the block.

## 3. **Block-Level Scope**

A block-level variable is declared inside curly braces `{}` and is only accessible within that specific block. These variables can shadow variables declared in outer scopes, such as the function or package level.

### Example of Block-Level Scope:

```go
// Private function (only accessible within this package)
func printByeScope() {
    // Declaring a new variable inside a block
    x := "hello bye scope"

    {
        // New scope with a new declaration of x (block-level scope)
        x := "new hello bye scope"
        fmt.Println(x) // Output: new hello bye scope
    }

    // The variable x here refers to the function-level x, which has not been modified in the block
    fmt.Println(x) // Output: hello bye scope

    // Accessing the package-level variable 'y'
    fmt.Println(y) // Output: hello scope from package
}
```
### Explanation:

- Inside the `printByeScope` function, a variable `x` is declared with `x := "hello bye scope"`. This is a **function-level variable**.
- Within the block, the variable `x` is shadowed (re-declared) with `x := "new hello bye scope"`. This is a **block-level variable** that only exists within the block.
- The reassignment of `x` inside the block does not affect the `x` in the function scope, as block-level variables shadow but do not modify the variables in outer scopes.
- The **package-level variable** `y` is accessible inside the function and block because it is declared outside the function.

### Key Points:
- **Block-level variables** are confined to the block where they are declared and cannot be accessed outside of that block.
- Block-level variables can **shadow** variables from outer scopes (function or package level) but do not modify the outer variables.
- If a block-level variable has the same name as a variable in an outer scope, the block-level variable takes precedence within that block.
### Code Explanation

- **Package-Level Variable:**
    - `var y = "hello scope from package"` is declared at the package level. It can be accessed from any function within the `secondclass` package.

- **Function-Level Variable:**
    - `var x = "hello scope"` is declared inside the `PrintHelloScope` function, making `x` accessible only within this function.

- **Block-Level Scoping:**
    - Inside the block `{}`, Go allows the use of block-level scoping. In the `printByeScope` function, the variable `x` is shadowed inside the block, meaning that the block declares its own `x`, which doesn't affect the `x` declared in the function.

- **Package-Level Variable Access:**
    - The package-level variable `y` is printed from both functions without modification. It retains its value throughout the program.

### Expected Output

The expected output when running the code would be:

```go
new hello scope
new hello scope
hello scope from package
new hello bye scope
hello bye scope
hello scope from package
```