package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = "Bonjour, "
	case "Spanish":
		prefix = "Hola, "
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "Sakfja"))
}
