package main

// import a package that contains Println function
import "fmt"

const englishHelloPrefix = "Hello, "

func Hello() string {
	return englishHelloPrefix + "world"
}

func HelloName(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello())
}
