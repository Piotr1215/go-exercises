# Learning Go

Exercises based on https://www.udemy.com/share/1013gwCUcSdVxRRnQ=/

## Topics

- [x] [Variables, values, types](01-variables-value-types.go)
- [ ] Programming fundamentals
- [x] Control flow
- [x] Grouping data
- [ ] Structs
- [ ] Functions
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
- each file has one metod exposed that in turn calls all other methods
- method names describe type of exercise, sometimes comments are used for clarity
- use VS Code debugger to step through code
  - set brakepoint on the module you want to check out: <kbd>F9</kbd>
  - run with debugger: <kbd>F5</kbd>
- from here you can:
  - step into a function: <kbd>F11</kbd>
  - step over line of code (let it execute): <kbd>F10</kbd>
  - step out from function (useful for stepping out of go code and long loops): <kbd>Shift + F11</kbd>
