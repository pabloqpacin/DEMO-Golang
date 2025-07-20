/*
> [@Fireship: Go in 100 Seconds](https://www.youtube.com/watch?v=446E-r0rXHI)

To compile and run program:
$ go run hello.go      # executable binary in memory
$ go build hello.go    # executable in same dir as src file
*/

package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	hungarian = "Hungarian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Salut, "
	hungarianHelloPrefix = "Szia, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case hungarian:
		prefix = hungarianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("World!", "English"))
}
