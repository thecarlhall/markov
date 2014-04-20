package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
)

const (
	filename = "./grimms.txt"
)

func readFile() []string {
	content, _ := ioutil.ReadFile(filename)
	tokens := strings.FieldsFunc(string(content), func(r rune) bool {
		return r == ' ' || r == '\n'
	})
	return tokens
}

func parse(tokens []string) (roots []string, states map[string][]string) {
	rootSet := make(map[string]struct{})
	states = make(map[string][]string)

	reg, _ := regexp.Compile("[A-Za-z0-9.,:;!'\"]+")
	for i := 1; i < len(tokens); i++ {
		currentToken := reg.FindString(tokens[i])

		prevToken := tokens[i-1]
		nextStates := append(states[prevToken], currentToken)
		states[prevToken] = nextStates

		if strings.Title(prevToken) == prevToken {
			rootSet[prevToken] = struct{}{}
		}
	}

	roots = []string{}
	for root, _ := range rootSet {
		roots = append(roots, root)
	}

	return
}

func generate(roots []string, states map[string][]string) {
	word := roots[rand.Intn(len(roots)-1)]

	for i := 0; i < 1000; i++ {
		fmt.Print(word, " ")
		nextStates := states[word]
		if len(nextStates) == 0 {
			word = roots[rand.Intn(len(roots)-1)]
		} else if len(nextStates) == 1 {
			word = nextStates[0]
		} else {
			word = nextStates[rand.Intn(len(nextStates)-1)]
		}
	}
	fmt.Println("")
}

func main() {
	tokens := readFile()
	roots, states := parse(tokens)
	generate(roots, states)
}
