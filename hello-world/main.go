package main

import "fmt"

func main() {
	fmt.Println("Hi world!")
}

/*

how do we run the project?
	1. go build - compiles a bunch of go source code files.
		- will create a file that's runnable depending on the os. = ./main
	2. go run - compiles and executes one or two files. =
	3. go fmt - Formats all the code in each file in the current directory.
	4. go install - Compiles and "installs" a package.
	5. go get - Downloads the raw source code of someone else's package.
	6. go test - Runs any tests associated with the current project.

what does `package main`` mean?
	package == project == workspacke
	packag is collection of common source code files.
	a package can contain multiples files with same package name.
	only requirement of any file inside a package is the first line must be the package name.
	Inside of go there are 2 types of packages
		1. Executable - Generates a file that we can run. eg: main
			the name of the package that determines u r making the executable or dep
			must have a func called 'main'
		2. Reusable - Code used as 'helpers; Good place to put reusable logic
			if the name is any other name it won't create executable file.

what does `import "fmt"` mean?
	give the package main all the functionality contains in fmt library


what's that 'func' thing?
	func is short for function
	just like any other programming langs it's similar to them

How is the main.go file organized?
	package first
	then the import statements for that package
	then body of the file variables functions etc.

*/
