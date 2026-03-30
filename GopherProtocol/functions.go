package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToUpperCase(s string) string {
	return strings.TrimSpace(strings.ToUpper(s))
}

func ToLowerCase(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func capsFirstLetter(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
	}
	return strings.Join(words, " ")
}

func titleCase(s string) string {
	minorWords := map[string]bool{
		"a": true, "an": true, "the": true,
		"and": true, "but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true,
	}

	words := strings.Fields(strings.ToLower(s))
	for i, word := range words {
		if i == 0 || !minorWords[word] {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

func snakeCase(s string) string {
	result := strings.FieldsFunc(strings.ToLower(s), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	return strings.Join(result, "_")

}

func reverser(s string) string {
	r := []rune(strings.TrimSpace(s))

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseSentence(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words)-1; i++ {
		words[i] = reverser(words[i])
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ToUpperCase("sentinel is online"))
	fmt.Println(ToLowerCase("ALERT LEVEL FIVE DETECTED"))
	fmt.Println(capsFirstLetter("THREAT LEVEL elevated"))
	fmt.Println(titleCase("a threat in the north"))
	fmt.Println(snakeCase("Alert! Level 5 detected."))
	fmt.Println(reverseSentence("Go is fun"))
}
