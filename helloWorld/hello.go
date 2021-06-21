package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("", ""))
}

const spanish = "spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const defaultGreetingsReceiver = "World"

func Hello(name, lang string) string {
	if name == "" {
		name = defaultGreetingsReceiver
	}
	return languagePrefix(lang) + name
}

func languagePrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
