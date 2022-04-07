package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishhHelloPrefix = "Holla, "
const frenchhHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchhHelloPrefix
	case "Spanish":
		prefix = spanishhHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
