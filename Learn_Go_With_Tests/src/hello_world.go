/*
> [@Fireship: Go in 100 Seconds](https://www.youtube.com/watch?v=446E-r0rXHI)

To compile and run program:
$ go run hw.go      # executable binary in memory
$ go build wg.go    # executable in same dir as src file
*/

package main

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
}
