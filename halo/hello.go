package main

import "fmt"

const (
	spanish = "Spanish"
	franch  = "Franch"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	franchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case franch:
		prefix = franchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("World", "Spanish"))
}
