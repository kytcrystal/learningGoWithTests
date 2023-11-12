package main

// import a package that contains Println function
// import "fmt"

const (
	french  = "French"
	spanish = "Spanish"
	italian = "Italian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	italianHelloPrefix = "Ciao, "
)

func Hello(name string, language string) string {
	//public functions start with uppercase
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	//private functions start with lowercase
	//prefix is a named return value
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
