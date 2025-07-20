# I. Go fundamentals


<details>
<summary>ToC
</summary>

- [I. Go fundamentals](#i-go-fundamentals)
  - [1. Install Go](#1-install-go)

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


<!-- ## 2. Hello, World -->
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

