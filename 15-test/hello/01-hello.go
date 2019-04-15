package hello

import "fmt"

const helloEnglish string = "Hello "
const helloSpanish string = "Hola "
const helloFrench string = "Bonjour "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}
	if language == "Spanish" {
		return helloSpanish + name
	}
	if language == "French" {
		return helloFrench + name
	}

	return helloEnglish + name
}

func HelloSwitch(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { //named return value

	switch language {
	case "Spanish":
		prefix = helloSpanish
	case "French":
		prefix = helloFrench
	default:
		prefix = helloEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("World!", ""))
}
