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

## Program Structure

### Names

    - begins with a letter or an underscore, case sensitive

#### Keywords

- Go has 25 keywords as listed below

<table style="border:none;">
  <tr>
    <td style="border:none;">`break`</td>
    <td style="border:none;">`default`</td>
    <td style="border:none;">`func`</td>
    <td style="border:none;">`interface`</td>
    <td style="border:none;">`select`</td>
  </tr>
  <tr>
    <td style="border:none;">`case`</td>
    <td style="border:none;">`defer`</td>
    <td style="border:none;">`go`</td>
    <td style="border:none;">`map`</td>
    <td style="border:none;">`struct`</td>
  </tr>
  <tr>
    <td style="border:none;">`chan`</td>
    <td style="border:none;">`else`</td>
    <td style="border:none;">`goto`</td>
    <td style="border:none;">`package`</td>
    <td style="border:none;">`switch`</td>
  </tr>
  <tr>
    <td style="border:none;">`const`</td>
    <td style="border:none;">`fallthrough`</td>
    <td style="border:none;">`if`</td>
    <td style="border:none;">`range`</td>
    <td style="border:none;">`type`</td>
  </tr>
  <tr>
    <td style="border:none;">`continue`</td>
    <td style="border:none;">`for`</td>
    <td style="border:none;">`import`</td>
    <td style="border:none;">`return`</td>
    <td style="border:none;">`var`</td>
  </tr>
</table>

- Constants

<table style="border:none;">
  <tr>
    <td style="border:none;">`true`</td>
    <td style="border:none;">`false`</td>
    <td style="border:none;">`iota`</td>
    <td style="border:none;">`nil`</td>
  </tr>
</table>

- Types

<table style="border:none;">
  <tr>
    <td style="border:none;">`int`</td>
    <td style="border:none;">`int8`</td>
    <td style="border:none;">`int16`</td>
    <td style="border:none;">`int32`</td>
    <td style="border:none;">`int64`</td>
  </tr>
  <tr>
    <td style="border:none;">`uint`</td>
    <td style="border:none;">`uint8`</td>
    <td style="border:none;">`uint16`</td>
    <td style="border:none;">`uint32`</td>
    <td style="border:none;">`uint64`</td>
    <td style="border:none;">`uintptr`</td>
  </tr>
  <tr>
    <td style="border:none;">`float32`</td>
    <td style="border:none;">`float64`</td>
    <td style="border:none;">`complex128`</td>
    <td style="border:none;">`complex64`</td>
    
  </tr>
  <tr>
    <td style="border:none;">`bool`</td>
    <td style="border:none;">`byte`</td>
    <td style="border:none;">`rune`</td>
    <td style="border:none;">`string`</td>
    <td style="border:none;">`error`</td>
  </tr>
</table>

- Functions

<table style="border:none;">
  <tr>
    <td style="border:none;">`make`</td>
    <td style="border:none;">`len`</td>
    <td style="border:none;">`cap`</td>
    <td style="border:none;">`new`</td>
    <td style="border:none;">`append`</td>
    <td style="border:none;">`copy`</td>
    <td style="border:none;">`close`</td>
    <td style="border:none;">`delete`</td>
  </tr>
  <tr>
    <td style="border:none;">`complex`</td>
    <td style="border:none;">`real`</td>
    <td style="border:none;">`img`</td>
  </tr>
  <tr>
    <td style="border:none;">`panic`</td>
    <td style="border:none;">`recover`</td>
  </tr>
</table>
