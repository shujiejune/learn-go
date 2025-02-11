## Package and Module

A module is a Go project, consisting of multiple packages.
Run command *go mod init [site/app-name]* to tell Go there is a module. 
Run command *go build* to get the [app-name] executable file and [go.mod].
Run command *./app-name* to run the Go project.
Every module has a [go.mod] file, it tells Go where the module is and the go version.

## Variables and Values

Built-in basic types:
- int
- float64
- string, created via **double quotes** or **backticks**
- bool, *true* or *false*
- uint, unsigned interget, strictly no-negative
- int32, 32-bit signed integer
- uint32, 32-bit unsigned integer
- int64, 64-bit signed integer
- rune, an alias for int32, represents a Unicode point

2 ways to declare a variable:
```go
// explicit
var item float64 = 0.0 

// implicit
item := 0.0
```

input and output:
```go
import "fmt"
//input, cannot read multiple words
var amount float64
fmt.Scan(&amount)

//output
fmt.Printf("Amount: %v", amount)
fmt.Println(amount)

// Sprintf formats according to a format specifier and returns the resulting string
// Sprintf(format string, a ...any) string
// Notice: you must Print out the return string of Sprintf
formattedFV := fmt.Sprintf("Amount: %.1f\n", amount)
fmt.Print(formattedFV)
```

**Notice:** when building multiline strings, use **backtick**
e.g. `This could be a multiline string.`

## Functions

```go
func yourFuncName(var1 type1, var2 type2, ...) (returnType1, ...)
{
    // do stuff
    // ret1 := ...
    return ret1, ...
}

// another return value syntax
func yourFuncName(var1 type1, var2 type2, ...) (ret1 type1, ...)
{
    // do stuff
    // ret1 = ...
    return
    // to be more readable, you can also do:
    // return ret1, ...
}
```

**Use function as parameter:**

```go
// use the customed type of function to code efficiently
type transformFn func(int) int

func main() {
    numbers := []int{1, 2, 3, 4}
    doubled := transformNumbers(&numbers, double)
    tripled := transformNumbers(&numbers, triple)

    fmt.Println(doubled)
    fmt.Println(tripled)

    // anonymous function
    transformed := transformNumbers(&numbers, func(number int) int{
        return number * 2
    })
}

func transformNumbers(numbers *[]int, transformFn) []int {
    tNumbers := []int{}

    for _, val := range *numbers {
        tNumbers = append(tNumbers, transformFn(val))
    }

    return tNumbers
}

// return function as value
func getTransformer() transformFn {
    return double
}

func double(number int) int {
    return number * 2
}

func triple(number int) int {
    return number * 3
}

func createTransformer(factor int) func(int) int {
    return func(number int) int {
        return number * factor
    }
}
```

**Variadic function:**
*...* here will collect the standalone values of the specified type and merge them into a slice.

```go
func main() {
    sum := sumup(1, 10, 15, 40, -5)

    fmt.Println(sum)
}

func sumup(startValue int, numbers ...int) int {
    sum := startValue
    for _, val := range numbers {
        sum += val
    }
    return sum
}
```

## Control Structures

### if-else
```go
var choice int
if choice == 1 {
    // do stuff
} else if choice == 2 {
    // do stuff
    // you can add return here
} else {
    // do stuff
}
```

### switch-case
```go
var choice int
switch choice {
case 1:
  // do stuff
case 2:
  // do stuff
default:
  // do stuff
  // you can use return but not break to get out of the structure
}
```

### for loop
```go
for i := 0; i < N; i++ {
    // do stuff
}

// infinite loop
for {
    // do stuff
    // can use break or return to get out of the loop
    // can also use continue
}

// conditional loop
// equalence of while loop
for someCondition {
    // do stuff
}
```

### handle errors

In Go, errors won't crash the application.
```go
import "errors"

func funcName() (float64, error) {
    data, err := callFunc()
    if err != nil {
        return defaultVal, errors.New("Error message")
        // if you want to exit the program whenerrors happen:
        // panic("Panic message")
    }
    return data, nil
}
```

## Read and Write to Files

```go
import (
    "fmt",
    "os",
    "strconv"
)

func writeBalance(balance float64) {
    // WriteFile(name string, data []byte, perm os.FileMode) error
    balanceText := fmt.Sprint(balance)
    os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func getBalance() float64 {
    // ReadFile(name string) ([]byte, error)
    data, _ := os.ReadFile("balance.txt")
    balanceText := string(data)
    balance, _ := strconv.ParseFloat(balanceText, 64)
    return balance
}
```
**0644:** the owner of the file is able to read/write, while other users can only read it

## Packages

- You can use functions in different files but belonging to the same package.
- Every package must go to its own subfolder of the same name in Go.
- Import your own package with the entire module path, e.g. *example.com/bank/fileops*
- Only functions, variables, etc. with an uppercase first character are available in other packages. They can be imported like *fileops.GetFloatFromFile*
- If you want to use third party packages, require them with module paths and version info in *go.mod*

## Struct Type

You can create a struct type either globally on in a function.

```go
import "time"

type User struct {
    firstName string
    lastName string
    birthDate stirng
    createdAt time.Time
}

// anonymous embedded struct
type Admin struct {
    email string
    password string
    User
}

// attach a method to the struct by adding a receiver argument in front of the method name
func (u, *User) outputUserDetails() {
    fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// constructor function
func newUser(firstName, lastName, birthDate string) *User {
    return &User{
        firstName: firstName,
        lastName: lastName,
        birthDate: birthDate,
        createdAt: time.Now(),
    }
}

func main() {
    var appUser *User
    // instantiate a struct
    appUser = newUser()
    appUser.outputUserDetails()
}
```
If you put the struct in a separate package, you should use uppercase first character for the fields/methods you want to export.

You can also use *type* for creating an alias of a baked type.
```go
type customString string

func main() {
    var name customString
}
```

## Interface

The interface type would automatically scan the passed data to examine if it has the methods and return types as regulated inside the interface.
```go
type saver interface {
    Save() error
}

// embedded interface
type outputtable interface {
    saver
    Display()
}
```
There is an interface that can contain any type: *interface{}*

Since *interface{}* is too broad and operators like *+* would be confused with it, you can use generics.
```go
import "fmt"

func add[T int|float64|string](a, b T) T {
    return a + b
}

func main() {
    result := add(1, 2)
    fmt.Println(result)
}
```

## Data Structure

### Array

```go
func main{
    var prodNames [3]string
    prices := [3]float64{10.99, 45.99, 20.00}
    featuredPrices = prices[1:]
    highlightedPrices = featuredPrices[:1]
    // len(highlightedPrices) = 1, cap(highlightedPrices) = 3
}
```
slicing arrays:
- [include, exclude), 0-index
- cannot use negative slicing index like in python
- can use at most *N* as slicing index for array of size *N*
- can slice other slices
- slice is a window to the original array, when you modify an element in the slice, the original array is also modified
- len() returns the number of elements in the array/slice, while cap() returns how many elements are possible to be selected. capacity counts to the end of the original array, cannot count towards the start. 

```go
func main() {
    prices := []float64{10.99, 8.99}
    prices = append(prices, 45.99)

    discountPrices := float64{12.99, 23.59}
    prices = append(prices, discountPrices...)

    // make(array type, length, capacity)
    userNames := make([]string, 2, 5)
    userNames[0] = "Julie"
    userNames = append(userNames, "George")
    userNames = append(userNames, "Rosemarry")
    // userNames = ["Julie", , "George", "Rosemarry"]

    // can use for loop to iterate through the list
    for index, value := range userNames {
        fmt.Println("Index: ", index)
        fmt.Println("Value: ", value)
    }
}
```
**dynamic array:**
- if you omit the capacity of the array, Go creates a dynamic array (a slice)
- can use built-in *append* method to add new elements to array, this method returns a new slice (occupies extra memory) and does not modify the original array
- you can append as many elements as you want
- if you want to remove elements from the array, use slice
- cannot append another list of elements to the existing list, but can use operator *...* to extract all the elements from the other list and then append them to the existing list
- can use *make* method to allocate memory for the dynamic list in advance, the 2nd argument is the number of empty slots reserved in the list, and the 3rd argument is the capacity

### Map

```go
func main() {
    // map[keyType]valueType{}
    websites := map[string]string{
        "Google": "https://google.com",
        "Amazon Web Services": "https://aws.com",
    }

    // mutation of map
    fmt.Println(map["Google"])
    map["LinkedIn"] = "https://linkedin.com"
    delete(websites, "Amazon Web Services")

    // can also use make() method to pre-allocate memory for map
    courseRatings := make(map[string]float64, 3)

    // use for loop to iterate through the map
    for key, value := range courseRatings {
        fmt.Println("Key: ", key)
        fmt.Println("Value: ", value)
    }
}
```
