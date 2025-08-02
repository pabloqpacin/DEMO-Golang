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
  - [5. Arrays and slices](#5-arrays-and-slices)
    - [Sum function](#sum-function)
    - [range](#range)
    - [slices](#slices)
    - [coverage](#coverage)
    - [SumAll function](#sumall-function)
    - [SumAllTails function: pass empty slice for runtime error panic](#sumalltails-function-pass-empty-slice-for-runtime-error-panic)
  - [6. Structs, methods \& interfaces (\& TableDrivenTests)](#6-structs-methods--interfaces--tabledriventests)
  - [7. Pointers \& errors (errcheck)](#7-pointers--errors-errcheck)
  - [8. Maps (CRUD)](#8-maps-crud)

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



## 5. Arrays and slices

### Sum function

- Arrays capacity:
  > - Arrays have a fixed capacity which you define when you declare the variable. We can initialize an array in two ways:
  >   - `[N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}`
  >   - `[...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}`

```go
// sum_test.go
```
```go
// sum.go
```
```sh
go test
```

### range

> [@gobyexample: range (Arrays & For loops)](https://gobyexample.com/range)

```go
// sum.go
```

### slices

- Go has slices which do not encode the size of the collection and instead can have any size.

### coverage

> [Coverage Tool](https://go.dev/blog/cover)

```sh
go test -cover
```

### SumAll function

- We need a new function called `SumAll` which will take a varying number of slices, returning a new slice containing the totals for each slice passed in. Examples:
  - `SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}`
  - `SumAll([]int{1,1,1}) would return []int{3}`
- [`make` function](https://go.dev/tour/moretypes/13) to create slices


### SumAllTails function: pass empty slice for runtime error panic

- ...
- ... eventually, "we're showing a new technique, assigning a function to a variable", `checkSums := func(t testing.TB, got, want []int) {`


## 6. Structs, methods & interfaces (& TableDrivenTests)

- [struct **types**](https://go.dev/ref/spec#Struct_types) <!-- clases -->
- [types **method declarations**](https://go.dev/ref/spec#Method_declarations):
  - "A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type."
  - "Methods are very similar to functions but they are called by invoking them on an instance of a particular type. Where you can just call functions wherever you like, such as `Area(rectangle)` you can only call methods on "things"."
  - "It is a convention in Go to have the receiver variable be the first letter of the type." (`r Rectangle`)
- [interfaces](https://go.dev/ref/spec#Interface_types): (avoid duplication thru [parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism))
  - "are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety."
  - "In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile."
  - **Declouping**: "[...] By declaring an interface, the helper is decoupled from the concrete types and only has the method it needs to do its job. [...] This kind of approach of using interfaces to declare only what you need is very important in software design and will be covered in more detail in later sections."
- [**Table driven tests**](https://go.dev/wiki/TableDrivenTests):


## 7. Pointers & errors (errcheck)

<!-- > manage state... -->

>  [!TIP]
> - When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
> - [Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

- "In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in."
- "Without getting too computer-sciency, when you create a value - like a wallet, it is stored somewhere in memory. You can find out what the *address* of that bit of memory with `&myVal`."
- "*You can see that the addresses of the two balances are different. So when we change the value of the balance inside the code, we are working on a copy of what came from the test. Therefore the balance in the test is unchanged.*"
- [pointers](https://gobyexample.com/pointers)
- Pointers to structs even have their own name: **struct pointers** and they are [automatically dereferenced](https://go.dev/ref/spec#Method_values)
- [Stringer](https://pkg.go.dev/fmt#Stringer): permite extender el output...
- **Error management**:
  - `nil` is synonymous with `null` from other programming languages. Errors can be `nil` because the return type of `Withdraw` will be `error`, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.
  - Like `null` if you try to access a value that is `nil` it will throw a runtime panic. This is bad! You should make sure that you check for `nil`s.
  - `t.Fatal`: "`t.Fatal` which will stop the test if it is called. This is because we don't want to make any more assertions on the error returned if there isn't one around. Without this the test would carry on to the next step and panic because of a nil pointer."


## 8. Maps (CRUD)

- In *arrays & slices*, we saw how to store values in order. Now, we will look at a way to store items by a key and look them up quickly.
- **Declaration**:
  - Declaring a Map is somewhat similar to an array. Except, it starts with the `map` keyword and requires two types. The first is the key type, which is written inside the `[]`. The second is the value type, which goes right after the `[]`.
  - The key type is special. It can only be a comparable type because without the ability to tell if 2 keys are equal, we have no way to ensure that we are getting the correct value. Comparable types are explained in depth in the [language spec](https://go.dev/ref/spec#Comparison_operators).
- *Pointers, copies, et al*...
  - <!-- So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data. -->
  A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. You can read more about maps here.
  - Therefore, you should never initialize a nil map variable:
```go
var m map[string]string
```
- - Instead, you can initialize an empty map or use the make keyword to create a map for you:
```go
var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
```
- - Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime panic.
- (Case of repetitions:) Map will not throw an error if the value already exists. Instead, they will go ahead and overwrite the value with the newly provided value. This can be convenient in practice, but makes our function name less than accurate. Add should not modify existing values. It should only add new words to our dictionary.
- ...
- [`error` interface & constants...](https://dave.cheney.net/2016/04/07/constant-errors)






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

