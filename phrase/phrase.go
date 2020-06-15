package phrase

import (
	"fmt"
)

type word interface {
	phraseGen(word) string
	getValue() string
}

type adjective struct {
	value string
}

type noun struct {
	value string
}

func (n *noun) getValue() string {
	return n.value
}
func (n *noun) phraseGen(a word) string {
	return n.value + " " + a.getValue()
}

func (a *adjective) getValue() string {
	return a.value
}
func (a *adjective) phraseGen(n word) string {
	return n.getValue() + " " + a.value
}

var first = &noun{"kamaz"}
var second = &adjective{"dirty"}

// PhraseGen generate a right phrase for two words in any positions
func PhraseGen(word1 word, word2 word) string {
	return word1.phraseGen(word2)
}

func main() {
	fmt.Println(PhraseGen(first, second))
}
