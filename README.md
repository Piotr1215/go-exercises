# Learning Go

Exercises based on https://www.udemy.com/share/1013gwCUcSdVxRRnQ=/

## Topics

Learning topics with random notes and observations.

- [x] [Variables, values, types](./exercises/01-variables-value-types.go)
  - Go has type inference and compiler was able to infer type of all variables so far
  - <kbd>:=</kbd> [operator](https://golang.org/ref/spec#Short_variable_declarations) creates a variable and assigns value to it in one statement
- [ ] Programming fundamentals
- [x] [Control flow](exercises/03-control-flow.go)
- [x] [Grouping data](exercises/04-grouping-data.go)
- [x] [Structs](exercises/05-structs.go)
- [ ] [Functions](exercises/06-functions.go)
  - functions that have receiver (operate on a type), are called methods and usually mutate state
  - functions that are not methods try to be "pure functions"
  - > there is no enforcement of pure functions concept in Go

- [ ] Pointers
- [ ] Application
- [ ] Concurrency
- [ ] Channels
- [ ] Error handling
- [ ] Writing documentation
- [ ] Testing & benchmarking

## How to run the examples

Best way to check the code is to run in debug mode, stepping into functions and examining variables and program execution flow.

- each section is encapsulated in a separate `go` file
- each file has one meted exposed that in turn calls all other methods
- method names describe type of exercise, sometimes comments are used for clarity
- use VS Code debugger to step through code
  - set breakpoint on the module you want to check out: <kbd>F9</kbd>
  - run with debugger: <kbd>F5</kbd>
- from here you can:
  - step into a function: <kbd>F11</kbd>
  - step over line of code (let it execute): <kbd>F10</kbd>
  - step out from function (useful for stepping out of go code and long loops): <kbd>Shift + F11</kbd>


## Links & Resources

- [Effective Go](https://golang.org/doc/effective_go)