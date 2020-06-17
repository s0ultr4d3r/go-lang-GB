package main

import (
	"fmt"
	"sort"
	"chess"
	// "sortingaddress"
	"calculator"
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

// PhraseGen generate a right phrase for two words in any positions.
func PhraseGen(word1 word, word2 word) string {
	return word1.phraseGen(word2)
}

// Struct for contacts.
type Contact struct {
	Name string
	num1 string
	num2 string
}

type Contacts []*Contact

// Sortint realize.
func (c Contacts) Len() int      { return len(c) }
func (c Contacts) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// Byname realize sort in alphabet.
type ByName struct{ Contacts }

func (c ByName) Less(i, j int) bool { return c.Contacts[i].Name < c.Contacts[j].Name }

func main() { 

//Задача на варианты хода коня
	var x string
	var y int
	fmt.Print("Введите букву и цифру стартовой позиции: ")
    fmt.Scan(&x, &y)
	chess.KnightVars(x,y)

	// Задача на калькулятор
	input := ""
    for {
        fmt.Print("> ")
        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println(err)
            continue
        }

        if input == "exit" {
            break
        }
		
        if res, err := calculator.Calculate(input); err == nil {
            fmt.Printf("Результат: %v\n", res)
        } else {
            fmt.Println("Не удалось произвести вычисление")
		}
		if input == "help"{
			fmt.Println("Простой калькулятор, введите пример в строку без пробелов")
			
		}
	}
	


// Задание на сортировку

c := []*Contact{
	{"Bob", "89167243812", "-"},
	{"Alex", "89996543210", "89155243627"},
}

sort.Sort(ByName{c})
fmt.Println("Contacts by name:")
printContacts(c)
}

func printContacts(c []*Contact) {
for _, p := range c {
	fmt.Println("\n", p.Name, p.num1, p.num2)
}

//Задача на интерфейс и метод

fmt.Println(PhraseGen(first, second))

}