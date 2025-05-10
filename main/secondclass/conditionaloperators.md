# Go `if` Statement

In Go, the `if` statement is used for conditional execution. Go also supports an optional short statement before the condition for more concise code.

---

## Basic `if`

```go
x := 10

if x > 5 {
    fmt.Println("x is greater than 5")
}
```
## if-else and else if
```go
if x < 5 {
    fmt.Println("x is less than 5")
} else if x == 5 {
    fmt.Println("x is equal to 5")
} else {
    fmt.Println("x is greater than 5")
}
```
## if with Short Statement
Go allows a short statement before the condition in an if. This is commonly used to handle errors inline.
### Example: String to Integer Conversion
```go
import (
    "fmt"
    "strconv"
)

func main() {
    if num, err := strconv.Atoi("123"); err == nil {
        fmt.Println("Converted number:", num)
    } else {
        fmt.Println("Conversion failed:", err)
    }
}
```
### Notes
- The short statement and condition are separated by a semicolon ;. 
- Use short statements to keep code concise and variables scoped locally.

# Loops in Go

Go uses a single `for` keyword for all types of loops — including traditional `for`, `while`-like conditional loops, infinite loops, and iteration over collections.

Below are examples demonstrating all these loop forms.

---

## Traditional Loop (C-style `for`)

```go
// traditional loop
// var i=0 won't work here
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
- This is the most common form. 
- We must declare and initialize i inside the loop declaration. 
- Equivalent to C/Java style loops.

## Conditional Loop (like `while`)
```go
// conditional loop (like while)
//var j int = 0, both will work here
j := 0
for j < 5 {
    fmt.Println(j)
    j++
}
```
- Acts like a while loop in other languages.
- The loop runs while the condition is true.
- Initialization is done outside the loop.

## Infinite Loop
```go
var k int = 0
for {
    fmt.Println("Running...")
    k++
    if k == 5 {
        break //need to break it with logic
    }
}
```
- Runs forever unless explicitly broken with break, return, or other logic. 
- Useful in goroutines, servers, or continuous listeners.
## Loop Over Collection
```go
arr := []string{"a", "b", "c"}
for index, value := range arr {
    fmt.Println(index, value)
}
```
- Uses range to iterate over arrays, slices, maps, strings, or channels. 
- Returns index and value for slices/arrays. 
- Very useful for clean iteration over collections.
# Go `switch` Statement

In Go, the `switch` statement is a clean and powerful alternative to long `if-else` chains. It supports value-based matching, condition evaluation, type assertions, and even short statements.

---

## 🔹 Basic `switch`

```go
x := 2

switch x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
case 3:
    fmt.Println("Three")
default:
    fmt.Println("Other")
}
```
## 🔹 `switch` Without an Expression (Condition-Based)
```go
x := 75

switch {
case x < 50:
    fmt.Println("Less than 50")
case x >= 50 && x <= 100:
    fmt.Println("Between 50 and 100")
default:
    fmt.Println("More than 100")
}
```
- Works like if-else if-else. 
- Each case can be a boolean condition.
## 🔹 `switch` with Short Statement
```go
switch n := 3; n {
case 1:
    fmt.Println("One")
case 2, 3, 4:
    fmt.Println("Two to Four")
default:
    fmt.Println("Other")
}
```
- We can declare variables directly in the switch. 
- Cases can match multiple values using commas.
## 🔹 `fallthrough` Keyword
```go
x := 1

switch x {
case 1:
    fmt.Println("Case 1")
    fallthrough
case 2:
    fmt.Println("Case 2")
}
```
- `fallthrough` forces the next case to execute, regardless of the match. 
- it skips the condition check for the next case.
## 🔹 Type Switch
```go
var i interface{} = "hello"

switch v := i.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
default:
    fmt.Println("Unknown type")
}
```
- Used with interface{} to determine the dynamic type. 
- Only valid inside a switch for type assertions.

# Comparison: Go vs Java `switch` Statement

This table compares the `switch` statement features in **Go** and **Java**.

## ✅ Feature Comparison

| Feature                  | Go Support | Java Support |
|--------------------------|------------|--------------|
| Expression-based switch  | ✅         | ✅           |
| Condition-based switch   | ✅ (with empty expression) | ❌ (only expressions) |
| Multiple case values     | ✅ (e.g., `case 2, 3:`) | ✅ (from Java 7+) |
| Short statement support  | ✅ (`switch x := val; x`) | ❌ (not supported) |
| Type switching           | ✅ (with `interface{}`) | ❌ (not directly, only via `instanceof`) |
| Implicit break           | ✅         | ❌ (requires `break`) |
| Optional fallthrough     | ✅ (`fallthrough`) | ✅ (default unless `break`) |

---

## Notes

- **Java 14+** introduced enhanced `switch` expressions with arrow syntax and return values, closing the gap a bit.
- **Go** prefers simplicity and predictability — it breaks automatically.
- **Java** is more verbose and requires `break` statements to prevent accidental fallthrough.

---