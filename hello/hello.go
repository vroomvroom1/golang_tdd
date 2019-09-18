package main

import "fmt"

const englishHelloPrefix = "Hello, "
const dutchHelloPrefix = "Hallo, "
const frenchHelloPrefix = "Bonjuor, "
const japaneseHelloPrefix = "Kunichiwa, "

const dutch = "Dutch"
const french = "French"
const japanese = "Japanese"

// Hello function to be tested
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case dutch:
		prefix = dutchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Matt", french))
}
