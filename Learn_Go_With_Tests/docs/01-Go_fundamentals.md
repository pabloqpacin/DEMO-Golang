# I. Go fundamentals


<details>
<summary>ToC
</summary>

- [I. Go fundamentals](#i-go-fundamentals)
  - [1. Install Go](#1-install-go)
  - [2. Hello, World](#2-hello-world)
    - [First \*\_test.go, go test \& go doc](#first-_testgo-go-test--go-doc)
    - [pkgsite](#pkgsite)
    - [TDD](#tdd)
    - [Refactoring \& subtests](#refactoring--subtests)
  - [3. Integers](#3-integers)
    - [Testable Examples](#testable-examples)
  - [4. Iteration](#4-iteration)
    - [for loops](#for-loops)
    - [Benchmarking](#benchmarking)
    - [Strings Builder](#strings-builder)
    - [Practice exercises](#practice-exercises)

</details>


## 1. Install Go

> [!NOTE]
> [docs/101.install_go.md](/docs/101.install_go.md)

- Go 1.11 introduced [Modules](https://go.dev/wiki/Modules). This approach is the default build mode since Go 1.16, therefore the use of `GOPATH` is not recommended.
- Using Modules is pretty straightforward. Select any directory outside `GOPATH` as the root of your project, and create a new module with the `go mod init` command.

```sh
mkdir src && cd $_
go mod init lgwt/m

# go help mod init
# go help mod
# go help modules
# go help get
```

- Go Linting: [golangci-lint vscode extension](https://golangci-lint.run/welcome/integrations/#visual-studio-code)
- Refactoring and your tooling: ...
- Go tooling ecosystem: [https://awesome-go.com](https://awesome-go.com/)


## 2. Hello, World

### First *_test.go, go test & go doc

- Simple Hello World program and execution:

> - When you write a program in Go, you will have a main package defined with a main func inside it. Packages are ways of grouping up related Go code together.
> - The func keyword defines a function with a name and a body.
> - With import "fmt" we are importing a package which contains the Println function that we use to print.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
```sh
go run hello_world.go
```


- Simple testing

```go
// hello_world.go
```
```go
// hello_world_test.go
```
```sh
go test
  # PASS
  # ok      pabloqpacin.com/LearnGoWithTests_module 0.002s
```

- Misc:
  > - Go's second tool for viewing documentation is the pkgsite command, which powers Go's official package viewing website. You can install pkgsite with `go install golang.org/x/pkgsite/cmd/pkgsite@latest`, then run it with `pkgsite -open .`. Go's install command will download the source files from that repository and build them into an executable binary. For a default installation of Go, that executable will be in `$HOME/go/bin` for Linux and macOS, and `%USERPROFILE%\go\bin` for Windows. If you have not already added those paths to your $PATH var, you might want to do so to make running go-installed commands easier.
```sh
go doc fmt

# # NOTE: not done but interesting
# go install golang.org/x/pkgsite/cmd/pkgsite@latest
# pkgsite -open .
```

### pkgsite

```sh
go install golang.org/x/pkgsite/cmd/pkgsite@latest
pkgsite -open .

open http://localhost:8080
open http://localhost:8080/testing
```

### TDD

> - In the last example, we wrote the test after the code had been written so that you could get an example of how to write a test and declare a function. From this point on, we will be writing tests first. Our next requirement is to let us specify the recipient of the greeting.

```go
// hello_test.go
```
```go
// hello.go
```
```sh
go test
```

### Refactoring & subtests

- Constants
  > - It's worth thinking about creating constants to capture the meaning of values and sometimes to aid performance.
- Empty strings (**subtests**)
  > - The next requirement is when our function is called with an empty string it defaults to printing "Hello, World", rather than "Hello, ".
- Refactor assertion in the test file (make it a function)!!
  > - For helper functions, it's a good idea to accept a `testing.TB` which is an interface that `*testing.T` and `*testing.B` both satisfy, so you can call helper functions from a test, or a benchmark (don't worry if words like "interface" mean nothing to you right now, it will be covered later).
- ...


```go
// hello.go
```

- Refactor final:
  > - In our function signature we have made a named return value (prefix string).
  > - This will create a variable called prefix in your function.
  >   - It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
  >     - You can return whatever it's set to by just calling return rather than return prefix.
  >   - This will display in the Go Doc for your function so it can make the intent of your code clearer.
  > - default in the switch case will be branched to if none of the other case statements match.
  > - **The function name starts with a lowercase letter. In Go, public functions start with a capital letter, and private ones start with a lowercase letter. We don't want the internals of our algorithm exposed to the world, so we made this function private.**
  > - Also, we can group constants in a block instead of declaring them on their own line. For readability, it's a good idea to use a line between sets of related constants.


## 3. Integers

- Simple package:
  > - You will notice that we're using `%d` as our format strings rather than `%q`. That's because we want it to print an integer rather than a string.

```go
// adder_test.go
```
```sh
go test
```
```go
// adder.go
```
```sh
go test
```

### Testable Examples

- Add `ExampleAdd()` to adder_test.go
  > - Example functions are compiled whenever tests are executed. Because such examples are validated by the Go compiler, you can be confident your documentation's examples always reflect current code behavior.
  > - Notice the special format of the comment, `// Output: 6`. While the example will always be compiled, adding this comment means the example will also be executed. Go ahead and temporarily remove the comment `// Output: 6`, then run `go test`, and you will see `ExampleAdd` is no longer executed.
  > - **Examples without output comments are useful for demonstrating code that cannot run as unit tests, such as that which accesses the network, while guaranteeing the example at least compiles.**

```sh
cd **/src
tree
  #  .
  # ├──  01_helloworld
  # │   ├──  hello.go
  # │   └──  hello_test.go
  # ├──  02_integers
  # │   ├──  adder.go
  # │   └──  adder_test.go
  # └──  go.mod

go run ./...
go test ./...

pkgsite -open .
open http://localhost:8080
```


## 4. Iteration

### for loops

- `for` loops only (*in Go there are no `while`, `do`, `until` keywords*)
- Simple TDD:

```go
// repeat_test.go
```
```go
// repeat.go
```
```sh
go test
```

### Benchmarking

- [documentation](https://pkg.go.dev/testing#hdr-Benchmarks)
- Func Structure:

```go
func Benchmark(b *testing.B) {
	//... setup ...
	for b.Loop() {
		//... code to measure ...
	}
	//... cleanup ...
}
```
```sh
go test -bench=.
```

### Strings Builder

- std lib strings.Builder:
  > - Strings in Go are immutable, meaning every concatenation, such as in our `Repeat` function, involves copying memory to accommodate the new string. This impacts performance, particularly during heavy string concatenation.
  > - The standard library provides the `strings.Builder` [stringsBuilder](https://pkg.go.dev/strings#Builder) type which minimizes memory copying. It implements a WriteString method which we can use to concatenate strings:

```sh
go test -bench=. -benchmem
```

### Practice exercises

- [x] Change the test so a caller can specify how many times the character is repeated and then fix the code
- [x] Write `ExampleRepeat` to document your function
- [x] Have a look through the [strings](https://pkg.go.dev/strings) package. Find functions you think could be useful and experiment with them by writing tests like we have here. Investing time learning the standard library will really pay off over time
  - `strings.Contains`
  - `strings.Count`
  - `strings.HasPrefix`





<!-- ## 5. Arrays and slices -->
<!-- ## 6. Structs, methods & interfaces -->
<!-- ## 7. Pointers & errors -->
<!-- ## 8. Maps -->
<!-- ## 9. Dependency Injection -->
<!-- ## 10. Mocking -->
<!-- ## 11. Concurrency -->
<!-- ## 12. Select -->
<!-- ## 13. Reflection -->
<!-- ## 14. Sync -->
<!-- ## 15. Context -->
<!-- ## 16. Intro to property based tests -->
<!-- ## 17. Maths -->
<!-- ## 18. Reading files -->
<!-- ## 19. Templating -->
<!-- ## 20. Generics -->
<!-- ## 21. Revisiting arrays and slices with generics -->

