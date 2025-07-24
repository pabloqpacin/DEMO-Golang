# Learn Go with Tests

> [!IMPORTANT]
> [Learn Go With Tests website](https://quii.gitbook.io/learn-go-with-tests/)


- [Learn Go with Tests](#learn-go-with-tests)
  - [Index](#index)
  - [Takeaways](#takeaways)
    - [TDD cycle](#tdd-cycle)
    - [The TDD process and *why* the tests are important](#the-tdd-process-and-why-the-tests-are-important)
    - [Cost and coverage](#cost-and-coverage)


## Index

```md
> I. Go fundamentals
1. Install Go
2. Hello, World
3. Integers
4. Iteration
5. Arrays and slices
6. Structs, methods & interfaces
7. Pointers & errors
8. Maps
9. Dependency Injection
10. Mocking
11. Concurrency
12. Select
13. Reflection
14. Sync
15. Context
16. Intro to property based tests
17. Maths
18. Reading files
19. Templating
20. Generics
21. Revisiting arrays and slices with generics

> II. Testing fundamentals
22. Introduction to acceptance tests
23. Scaling acceptance tests
24. Working without mocks
25. Refactoring Checklist

> III. Build an application
26. Intro
27. HTTP server (_https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/_)
28. JSON, routing and embedding
29. IO and sorting
30. Command line & package structure
31. Time
32. WebSockets

> IV. Questions and answers
33. OS Exec
34. Error types
35. Context-aware Reader
36. Revisiting HTTP Handlers

> V. Meta
37. Why unit tests and how to make them work for you
38. Anti-patterns
39. Contributing
40. Chapter Template
```

---

## Takeaways

### TDD cycle

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

### The TDD process and *why* the tests are important

- *Write a failing test and see it fail* so we know we have written a *relevant* test for our requirements and seen that it produces an *easy to understand description of the failure*
- Writing the smallest amount of code to make it pass so we know we have working software
- *Then* refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with

In our case, we've gone from `Hello()` to `Hello("name")` and then to `Hello("name", "French")` in small, easy-to-understand steps.

Of course, this is trivial compared to "real-world" software, but the principles still stand. TDD is a skill that needs practice to develop, but by breaking problems down into smaller components that you can test, you will have a much easier time writing software.


### Cost and coverage

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. **Every test has a cost**.

[...] having two tests for this function is redundant. If it works for a slice of one size it's very likely it'll work for a slice of any size (within reason).

```sh
go test -cover
  # PASS
  # coverage: 100.0% of statements
```
