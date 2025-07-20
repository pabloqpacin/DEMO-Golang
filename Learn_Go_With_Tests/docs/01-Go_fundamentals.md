# I. Go fundamentals


<details>
<summary>ToC
</summary>

- [I. Go fundamentals](#i-go-fundamentals)
  - [1. Install Go](#1-install-go)
  - [2. Hello, World](#2-hello-world)
    - [First \*\_test.go, go test \& go doc](#first-_testgo-go-test--go-doc)
    - [TDD](#tdd)
    - [Refactoring \& subtests](#refactoring--subtests)

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


<!-- ## 3. Integers -->
<!-- ## 4. Iteration -->
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

