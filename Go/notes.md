# Golang Notes
Let's explore Golang and learn about it's concepts.

## Table of Contents
- [Golang Notes](#golang-notes)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [First Program](#first-program)
  - [Golang Compilation Process](#golang-compilation-process)
  - [Input and Output](#input-and-output)
    - [Print](#print)
    - [Println](#println)
    - [Scan](#scan)
    - [Scanln](#scanln)
  - [Comments](#comments)
    - [Single Line Comment](#single-line-comment)
    - [Multi Line Comment](#multi-line-comment)
  - [Variables and Datatypes](#variables-and-datatypes)
    - [A way to declare variables](#a-way-to-declare-variables)
    - [Short hand declaration](#short-hand-declaration)
    - [Zero values](#zero-values)
    - [Type Conversion](#type-conversion)
    - [Type Assertion](#type-assertion)
    - [Type Inference](#type-inference)
  - [Constants](#constants)
  - [Operators](#operators)
    - [Arithmetic Operators](#arithmetic-operators)
    - [Assignment Operators](#assignment-operators)
    - [Comparison Operators](#comparison-operators)
    - [Logical Operators](#logical-operators)
    - [Bitwise Operators](#bitwise-operators)
  - [Control Flow](#control-flow)
    - [If-Else](#if-else)
    - [Switch](#switch)
    - [For](#for)
    - [For-Range](#for-range)
    - [For-Each](#for-each)
    - [Break](#break)
    - [Continue](#continue)
    - [Goto](#goto)
  - [Functions](#functions)
  - [Methods](#methods)
  - [Interfaces](#interfaces)
  - [Structs](#structs)
  - [Arrays](#arrays)
  - [Slices](#slices)
  - [Maps](#maps)
  - [Pointers](#pointers)
  - [Packages](#packages)
  - [Imports](#imports)
  - [Errors](#errors)
  - [Logging](#logging)
  - [Testing](#testing)
  - [Concurrency](#concurrency)
  - [Channels](#channels)
  - [Context](#context)
  - [Files](#files)
  - [HTTP](#http)
  - [Encode and Decode](#encode-and-decode)
    - [JSON](#json)
    - [XML](#xml)
    - [YAML](#yaml)
    - [CSV](#csv)
  - [Connectors](#connectors)
    - [SQL](#sql)
    - [MongoDB](#mongodb)
    - [Redis](#redis)
  - [Web](#web)
  - [Websocket](#websocket)
  - [CLI](#cli)
  - [GUI](#gui)
  - [License](#license)
  - [References](#references)


## Introduction

## First Program

## Golang Compilation Process

## Input and Output

Input and output are the most important part of any programming language. In Golang, we use the `fmt` package to take input and print output. The `fmt` package is a part of the standard library of Golang. To use the `fmt` package, we need to import it in our program. To import the `fmt` package, we use the following syntax:
```go
import "fmt"
```

### Print

To print something in Golang, we use the `Print` function of the `fmt` package. The `Print` function takes a string as an argument and prints it to the console. An example to print something in Golang is as follows:
```go
package main

import "fmt"

func main() {
    fmt.Print("Hello World")
}
```

### Println

To print something in Golang, we use the `Println` function of the `fmt` package. The `Println` function takes a string as an argument and prints it to the console. An example to print something in Golang is as follows:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

### Scan

To take input in Golang, we use the `Scan` function of the `fmt` package. The `Scan` function takes a pointer to a variable as an argument and stores the input in that variable. An example to take input in Golang is as follows:
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Scan(&name)
    fmt.Println("Hello", name)
}
```

### Scanln

To take input in Golang, we use the `Scanln` function of the `fmt` package. The `Scanln` function takes a pointer to a variable as an argument and stores the input in that variable. An example to take input in Golang is as follows:
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Scanln(&name)
    fmt.Println("Hello", name)
}
```

## Comments

Comments are used to explain the code. It is not executed by the compiler. There are two types of comments in Golang.

### Single Line Comment

Single line comments are used to comment a single line. An example of single line comment is as follows:
```go
// This is a single line comment
```

### Multi Line Comment

Multi line comments are used to comment multiple lines. An example of multi line comment is as follows:
```go
/*
This is a multi line comment
This is a multi line comment
This is a multi line comment
*/
```

## Variables and Datatypes

Before diving in, you should note that Golang is a statically typed language. This means that you need to specify the type of the variable when you declare it. Also, you cannot change the type of the variable later on. One more important thing is that you need to use the variable once you declare it. If you don't use it, you will get an error.

**String** is a sequence of characters. An example to define a string datatype variables is as follows:
```go
var name string = "Susmita"
```

**Int** is simply any integer number. An example to define an integer datatype variable is as follows:
```go
var num int = -9
var age int = 21
```

**uint** stands for unsigned integer. It is a non-negative integer.
```go
var age uint = 21
```

**float** is a number with a decimal point. An example to define a float datatype variable is as follows:
```go
var num float32 = 3.14
var num float64 = 3.14
``` 
> Note: 
> 1) float64 means it takes 64 bits of memory and float32 means it takes 32 bits of memory.
> 2) When you define a float number, you need to specify the precision type of the number (64 or 32).
> 3) When you take a float number, by default it is considered as float64.

**bool** is a boolean value which can be either true or false. An example to define a boolean datatype variable is as follows:
```go
var isTrue bool = true
var isFalse bool = false
```

**byte** is an alias for uint8. It is equivalent to uint8 in every way. An example to define a byte datatype variable is as follows:
```go
var b byte = 255
```

**rune** is an alias for int32. It is equivalent to int32 in every way. An example to define a rune datatype variable is as follows:
```go
var r rune = 1024
```

**complex** is a complex number with float32 real and imaginary parts. An example to define a complex datatype variable is as follows:
```go
var c complex64 = 5 + 5i
var c complex128 = 5 + 5i
```

**Arrays** are a collection of elements of the same datatype. An example to define an array datatype variable is as follows:
```go
var arr [3]int = [3]int{1, 2, 3}
```

**Slices** are a collection of elements of the same datatype. An example to define a slice datatype variable is as follows:
```go
var slice []int = []int{1, 2, 3}
```

**Maps** are a collection of key-value pairs. An example to define a map datatype variable is as follows:
```go
var m map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
```

**Structs** are a collection of fields. An example to define a struct datatype variable is as follows:
```go
type Person struct {
    name string
    age int
}
var p Person = Person{"Susmita", 21}
```

**Pointers** are a reference to a memory address. An example to define a pointer datatype variable is as follows:
```go
var p *int = &num
```

**Functions** are a block of code that can be called later. An example to define a function datatype variable is as follows:
```go
func add(a int, b int) int {
    return a + b
}
```

**Interfaces** are a collection of methods. An example to define an interface datatype variable is as follows:
```go
type Shape interface {
    area() float64
}
```

**Channels** are a way for goroutines to communicate with each other. An example to define a channel datatype variable is as follows:
```go
var ch chan int = make(chan int)
```

### A way to declare variables
```go
var (
    name string = "Susmita"
    age int = 21
)
```

### Short hand declaration

In short hand declaration, you don't need to specify the datatype of the variable. Go automatically detects the type and works accordingly. An example to declare a variable using short hand declaration is as follows:

```go
name := "Susmita"
age := 21
```

### Zero values
```go
var name string // ""
var age int // 0
var isTrue bool // false
var isFalse bool // false
var b byte // 0
var r rune // 0
var c complex64 // (0+0i)
var c complex128 // (0+0i)
var arr [3]int // [0 0 0]
var slice []int // []
var m map[string]int // map[]
var p *int // <nil>
var ch chan int // <nil>
```

### Type Conversion

Type conversion is the conversion of one datatype to another. An example to convert one datatype to another is as follows:
```go
var num int = 21
var numFloat float64 = float64(num)
```

### Type Assertion

Type assertion is the conversion of an interface to another datatype. An example to convert an interface to another datatype is as follows:
```go
var i interface{} = 21
var num int = i.(int)
```

### Type Inference

Type inference is the automatic detection of the datatype of a variable. An example to detect the datatype of a variable is as follows:
```go
var name = "Susmita"
var age = 21
```

## Constants

Constants are the variables whose value cannot be changed. An example to define a constant is as follows:
```go
const name string = "Susmita"
```

## Operators

Operators are used to perform operations on variables and values. There are different types of operators in Golang.

### Arithmetic Operators

Arithmetic operators are used to perform arithmetic operations on variables and values. An example to perform arithmetic operations is as follows:
```go
var num1 int = 10
var num2 int = 20
var num3 int = num1 + num2 // 30
var num4 int = num1 - num2 // -10
var num5 int = num1 * num2 // 200
var num6 int = num1 / num2 // 0
var num7 int = num1 % num2 // 10
```

### Assignment Operators

Assignment operators are used to assign values to variables. An example to assign values to variables is as follows:
```go
var num1 int = 10
var num2 int = 20
num1 += num2 // 30
num1 -= num2 // -10
num1 *= num2 // 200
num1 /= num2 // 0
num1 %= num2 // 10
```

### Comparison Operators

Comparison operators are used to compare two values. An example to compare two values is as follows:
```go
var num1 int = 10
var num2 int = 20
var isTrue bool = num1 == num2 // false
var isFalse bool = num1 != num2 // true
var isTrue bool = num1 > num2 // false
var isTrue bool = num1 < num2 // true
var isTrue bool = num1 >= num2 // false
var isTrue bool = num1 <= num2 // true
```

### Logical Operators

Logical operators are used to perform logical operations on variables and values. An example to perform logical operations is as follows:
```go
var num1 int = 10
var num2 int = 20
var isTrue bool = num1 == num2 && num1 != num2 // false

var num1 int = 10
var num2 int = 20
var isTrue bool = num1 == num2 || num1 != num2 // true

var num1 int = 10
var isTrue bool = !(num1 == num2) // true
```

### Bitwise Operators

Bitwise operators are used to perform bitwise operations on variables and values. An example to perform bitwise operations is as follows:
```go
var num1 int = 10
var num2 int = 20
var num3 int = num1 & num2 // 0
var num4 int = num1 | num2 // 30
var num5 int = num1 ^ num2 // 30
var num6 int = num1 << num2 // 10485760
var num7 int = num1 >> num2 // 0
```

## Control Flow

Control flow statements are used to control the flow of execution of the program. There are different types of control flow statements in Golang.

### If-Else

If-Else statements are used to execute a block of code if a condition is true. An example to execute a block of code if a condition is true is as follows:
```go
var num1 int = 10
var num2 int = 20
if num1 == num2 {
    fmt.Println("num1 is equal to num2")
} else {
    fmt.Println("num1 is not equal to num2")
}
```

### Switch

Switch statements are used to execute a block of code based on the value of a variable. An example to execute a block of code based on the value of a variable is as follows:
```go
var num int = 10
switch num {
    case 10:
        fmt.Println("num is equal to 10")
    case 20:
        fmt.Println("num is equal to 20")
    default:
        fmt.Println("num is not equal to 10 or 20")
}
```

### For

For statements are used to execute a block of code repeatedly. An example to execute a block of code repeatedly is as follows:
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### For-Range

For-Range statements are used to iterate over a collection of elements. An example to iterate over a collection of elements is as follows:
```go
var arr [3]int = [3]int{1, 2, 3}
for index, value := range arr {
    fmt.Println(index, value)
}
```

### For-Each

For-Each statements are used to iterate over a collection of elements. An example to iterate over a collection of elements is as follows:
```go
var arr [3]int = [3]int{1, 2, 3}
for _, value := range arr {
    fmt.Println(value)
}
```

### Break   

Break statements are used to break out of a loop. An example to break out of a loop is as follows:
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

### Continue

Continue statements are used to skip the current iteration of a loop. An example to skip the current iteration of a loop is as follows:
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        continue
    }
    fmt.Println(i)
}
```

### Goto

Goto statements are used to jump to a label. An example to jump to a label is as follows:
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        goto label
    }
    fmt.Println(i)
}
label:
fmt.Println("Jumped to label")
```

## Functions

Functions are used to perform a specific task. An example to define a function is as follows:
```go
func add(num1 int, num2 int) int {
    return num1 + num2
}
```

## Methods

Methods are functions that are defined on a type. An example to define a method is as follows:
```go
type Person struct {
    name string
    age int
}

func (p Person) getName() string {
    return p.name
}
```

## Interfaces

Interfaces are used to define a set of methods that a type must implement. An example to define an interface is as follows:
```go
type Person interface {
    getName() string
}
```

## Structs

Structs are used to define a collection of fields. An example to define a struct is as follows:
```go
type Person struct {
    name string
    age int
}
```

## Arrays

Arrays are used to store a collection of elements. An example to define an array is as follows:
```go
var arr [3]int = [3]int{1, 2, 3}
```

## Slices

Slices are used to store a collection of elements. An example to define a slice is as follows:
```go
var arr []int = []int{1, 2, 3}
```

## Maps

Maps are used to store a collection of key-value pairs. An example to define a map is as follows:
```go
var m map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
```

## Pointers

Pointers are used to store the address of a variable. An example to define a pointer is as follows:
```go
var num int = 10
var ptr *int = &num
```

## Packages

Packages are used to organize code into different files and folders. An example to define a package is as follows:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

## Imports

Imports are used to import packages. An example to import a package is as follows:
```go
import "fmt"
```

## Errors

Errors are used to handle errors. An example to handle errors is as follows:
```go
var num1 int = 10
var num2 int = 0
var result int = 0
var err error = nil
if num2 == 0 {
    err = errors.New("Cannot divide by zero")
} else {
    result = num1 / num2
}
```

## Logging

Logging is used to log messages. An example to log messages is as follows:
```go
log.Println("Hello World!")
```

## Testing

Testing is used to test code. An example to test code is as follows:
```go
func TestAdd(t *testing.T) {
    if add(1, 2) != 3 {
        t.Error("Expected 3")
    }
}
```

## Concurrency

Concurrency is used to run multiple tasks concurrently. An example to run multiple tasks concurrently is as follows:
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    fmt.Println("Hello World!")
    wg.Done()
}()
wg.Wait()
```

## Channels

Channels are used to communicate between goroutines. An example to communicate between goroutines is as follows:
```go
var ch chan int = make(chan int)
go func() {
    ch <- 1
}()
var num int = <-ch
```

## Context

Context is used to manage cancellation signals. An example to manage cancellation signals is as follows:
```go
var ctx context.Context = context.Background()
var cancel context.CancelFunc = nil
ctx, cancel = context.WithCancel(ctx)
cancel()
```

## Files

Files are used to read and write files. An example to read and write files is as follows:
```go
var file *os.File = nil
var err error = nil
file, err = os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
var data []byte = make([]byte, 100)
var count int = 0
count, err = file.Read(data)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data[:count]))
```

## HTTP

HTTP is used to make HTTP requests. An example to make HTTP requests is as follows:
```go
var client *http.Client = &http.Client{}
var req *http.Request = nil
var err error = nil
req, err = http.NewRequest("GET", "https://example.com", nil)
if err != nil {
    log.Fatal(err)
}
var resp *http.Response = nil
resp, err = client.Do(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
var data []byte = make([]byte, 100)
var count int = 0
count, err = resp.Body.Read(data)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data[:count]))
```

## Encode and Decode

Encode and decode are used to encode and decode data. Examples to encode and decode data are as follows:

### JSON

JSON is used to encode and decode JSON. An example to encode and decode JSON is as follows:
```go
type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

var person Person = Person{"John Doe", 30}
var data []byte = nil
var err error = nil
data, err = json.Marshal(person)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
var person2 Person = Person{}
err = json.Unmarshal(data, &person2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(person2)
```

### XML

XML is used to encode and decode XML. An example to encode and decode XML is as follows:
```go
type Person struct {
    Name string `xml:"name"`
    Age int `xml:"age"`
}

var person Person = Person{"John Doe", 30}
var data []byte = nil
var err error = nil
data, err = xml.Marshal(person)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
var person2 Person = Person{}
err = xml.Unmarshal(data, &person2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(person2)
```

### YAML

YAML is used to encode and decode YAML. An example to encode and decode YAML is as follows:
```go
type Person struct {
    Name string `yaml:"name"`
    Age int `yaml:"age"`
}

var person Person = Person{"John Doe", 30}
var data []byte = nil

var err error = nil
data, err = yaml.Marshal(person)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))

var person2 Person = Person{}
err = yaml.Unmarshal(data, &person2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(person2)
```

### CSV

CSV is used to encode and decode CSV. An example to encode and decode CSV is as follows:
```go
type Person struct {
    Name string `csv:"name"`
    Age int `csv:"age"`
}

var person Person = Person{"John Doe", 30}
var data []byte = nil
var err error = nil
data, err = csv.Marshal(person)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
var person2 Person = Person{}
err = csv.Unmarshal(data, &person2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(person2)
```

## Connectors

Connectors are used to connect to databases.

### SQL

SQL is used to query SQL databases. An example to query SQL databases is as follows:
```go
var db *sql.DB = nil
var err error = nil
db, err = sql.Open("sqlite3", "file:database.db")
if err != nil {
    log.Fatal(err)
}

var rows *sql.Rows = nil
rows, err = db.Query("SELECT * FROM people")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
    var name string = ""
    var age int = 0
    err = rows.Scan(&name, &age)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name, age)
}
```

### MongoDB

MongoDB is used to query MongoDB databases. An example to query MongoDB databases is as follows:
```go
var client *mongo.Client = nil
var err error = nil
client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
if err != nil {
    log.Fatal(err)
}
err = client.Connect(context.Background())  
if err != nil {
    log.Fatal(err)
}
defer client.Disconnect(context.Background())

var collection *mongo.Collection = client.Database("test").Collection("people")
var ctx context.Context = context.Background()
var cursor *mongo.Cursor = nil
cursor, err = collection.Find(ctx, bson.M{})
if err != nil {
    log.Fatal(err)
}
defer cursor.Close(ctx)
for cursor.Next(ctx) {
    var person Person = Person{}
    err = cursor.Decode(&person)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(person)
}
```

### Redis

Redis is used to query Redis databases. An example to query Redis databases is as follows:
```go
var client *redis.Client = nil
var err error = nil
client = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
})
var ctx context.Context = context.Background()
var keys []string = nil
keys, err = client.Keys(ctx, "*").Result()
if err != nil {
    log.Fatal(err)
}
for _, key := range keys {
    var value string = ""
    value, err = client.Get(ctx, key).Result()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(key, value)
}
```

## Web

Web is used to create web servers. An example to create web servers is as follows:
```go
var mux *http.ServeMux = http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
})
var server *http.Server = &http.Server{
    Addr: ":8080",
    Handler: mux,
}
server.ListenAndServe()
```

## Websocket

Websocket is used to create websocket servers. An example to create websocket servers is as follows:
```go
var mux *http.ServeMux = http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    var conn *websocket.Conn = nil
    var err error = nil
    conn, err = websocket.Upgrade(w, r, w.Header(), 1024, 1024)
    if err != nil {
        log.Fatal(err)
    }
    for {
        var message string = ""
        err = conn.ReadJSON(&message)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(message)
        err = conn.WriteJSON(message)
        if err != nil {
            log.Fatal(err)
        }
    }
})
var server *http.Server = &http.Server{
    Addr: ":8080",
    Handler: mux,
}
server.ListenAndServe()
```

## CLI

CLI is used to create command line interfaces. An example to create command line interfaces is as follows:
```go
var app *cli.App = cli.NewApp()
app.Name = "hello"
app.Usage = "Prints hello world"
app.Action = func(c *cli.Context) error {
    fmt.Println("Hello World!")
    return nil
}
app.Run(os.Args)
```

## GUI

GUI is used to create graphical user interfaces. An example to create graphical user interfaces is as follows:
```go
var window *gtk.Window = nil
gtk.Init(nil)
window, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
window.SetTitle("Hello World")
window.Connect("destroy", func() {
    gtk.MainQuit()
})
window.SetDefaultSize(800, 600)
window.ShowAll()
gtk.Main()
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## References

* [Go](https://golang.org/)
* [GoDoc](https://godoc.org/)
* [Go Report Card](https://goreportcard.com/)
* [Travis CI](https://travis-ci.org/)
* [Codecov](https://codecov.io/)
* [Go by Example](https://gobyexample.com/)
* [Go Web Examples](https://gowebexamples.com/)

---

Made with ❤ by [**@Susmita-Dey**](https://github.com/Susmita-Dey). Feel free to star ⭐ this repository if you like this project. 😊