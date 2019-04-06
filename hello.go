package main

import "fmt"
import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {
	var a string
	for i := 0; i < len(sliceToConcat); i++ {
		a += string(sliceToConcat[i])
		fmt.Println(string(sliceToConcat[i]))
		if (i + 1) != len(sliceToConcat) {
			a += string('-')
		}
	}
	return a
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {

	for i := 0; i < len(sliceToEncrypt); i++ {

		sliceToEncrypt[i] = sliceToEncrypt[i] + byte(ceaserCount-48)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

func main() {
	slc := make([]byte, 2)
	slc[0] = 48
	slc[1] = 49
	fmt.Println(ConcatSlice(slc))

}
