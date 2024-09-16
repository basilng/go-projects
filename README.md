# go-projects

Learning Go Language By Doing Projects.

### Installing Go

go to https://go.dev/dl/

### Getting to know About Packages in GO

go to https://pkg.go.dev/std

### Command Line Arguments in Go

Command line arguments are available to a program in a variable named `Args` that is part of the os package.

os.Args is a slice of strings

- first element of the os.Args, `os.Args[0]` is the name of the command itself.

#### Basic Go Types

| Tyoe    | Example                    |
| ------- | -------------------------- |
| bool    | true, false                |
| string  | "Hi", "Hows it going?"     |
| int     | 0, -100000, 99999          |
| float64 | 10.00002, 0.0007, -100.003 |

#### Datastructures in Go

Array :- Fixed length list of things same type

Slice :- An array that can grow or shrink same type

#### Type Conversion in Go

Type we want(Value we have) eg: []byte("hi")

### Testing in Go

Go resting is not like Rspec, selenium, etc!
to make a test, create a new file ending in \_test.go
To run all tests in a package, run the command
`go test`

### Struct

Data structure, Collection of properties that are related together
embedded struct

### Pointers

by default go is using pass by value language
wherever we pass a value to a function, go will take that value and copy it in another memory location
The slice we are storing in memory in different way it contains the value [length, capacity, ptr to head]
ptr to head is always pointing to the same memory address in pass by value

| Value Types | Reference Types |
| ----------- | --------------- |
| bool        | slices          |
| string      | maps            |
| int         | channels        |
| float       | pointers        |
| struct      | functions       |

### Map

key value pair of same type
it's statically typed

### Map vs Struct

| Map                                                     | Struct                                                           |
| ------------------------------------------------------- | ---------------------------------------------------------------- |
| \* ALl keys must be the same type                       | \* Values van be of different type                               |
| \* Use to represent a collection of releated properties | \* You need to know all the different fields at compile time     |
| \* All values must be the same type                     | \* Keys don't support indexing                                   |
| \* Don't need to know all the keys at compile time      | \* Use to represent a "thing" with a lot of different properties |
| \* Keys are indexed we can iterate over them            | \* Value Type!                                                   |
| \* Reference Type!                                      |

### Interface

- interfaces are not generic types
- interfaces are 'implicit' // no need to manually mention
- interfaces are a contract to help us manage types
- interface are tough.

### Concurrency in Go

- A `goroutine` is a concurrent function execution
- A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine
- The function `main` runs in a `goroutine` and the `go` statement creates another `goroutine`s
- When one `goroutine` attempts a send or receive on a channel, it blocks until another `goroutine` attempts the corresponding receive or send operation, at which point the value is transferred and both goroutines proceeed.
