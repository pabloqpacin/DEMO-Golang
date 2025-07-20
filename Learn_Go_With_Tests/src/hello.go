/*
> [@Fireship: Go in 100 Seconds](https://www.youtube.com/watch?v=446E-r0rXHI)

To compile and run program:
$ go run hello.go      # executable binary in memory
$ go build hello.go    # executable in same dir as src file
*/

package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const hungarian = "Hungarian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Salut, "
const hungarianHelloPrefix = "Szia, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case hungarian:
		prefix = hungarianHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World!", "English"))
}
