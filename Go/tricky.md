## Tricky Questions and Answers in Go

### Q1: What is the output of the following code?

```go
package main

import "fmt"

func main() {
    var i int
    for i = 0; i < 10; i++ {
        defer fmt.Printf("%d ", i)
    }
}
```

> A1: 9 8 7 6 5 4 3 2 1 0

### Q2: What is the output of the following code?

```go
package main

import "fmt"

func main() {
    defer func() {
        fmt.Println("1")
    }()
    defer func() {
        fmt.Println("2")
    }()
    defer func() {
        fmt.Println("3")
    }()
}
```

> A2: 3 2 1

### Q3: What is the output of the following code?

```go
package main

import "fmt"

func main() {
    defer func() {
        fmt.Println("1")
    }()
    defer func() {
        fmt.Println("2")
    }()
    defer func() {
        fmt.Println("3")
    }()
    panic("panic")
}
```

> A3: 3 2 1 panic

### Q4: What is the output of the following code?

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Println("my age is", age)
}
```

> A4: my age is 0

### Q5: What is the output of the following code?

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Scan(&age)
    fmt.Println("my age is", age)
}
```

> A5: my age is 0 (if you enter 10, then the output will be my age is 10 else it will be 0 even if you enter a string or any other data type value)