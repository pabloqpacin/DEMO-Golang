/*
> [@Fireship: Go in 100 Seconds](https://www.youtube.com/watch?v=446E-r0rXHI)

To compile and run program:
$ go run hello.go      # executable binary in memory
$ go build hello.go    # executable in same dir as src file
*/

package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World!"))
}
