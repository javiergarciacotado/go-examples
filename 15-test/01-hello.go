package main

import "fmt"

const helloEnglish string = "Hello "
const helloSpanish string = "Hola "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}
	if language == "Spanish" {
		return helloSpanish + name
	}
	return helloEnglish + name
}

func main() {
	fmt.Println(Hello("World!", ""))
}
