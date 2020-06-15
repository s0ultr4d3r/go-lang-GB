package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	Name string
	num1 string
	num2 string
}

type Contacts []*Contact

func (c Contacts) Len() int      { return len(c) }
func (c Contacts) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

type ByName struct{ Contacts }

func (c ByName) Less(i, j int) bool { return c.Contacts[i].Name < c.Contacts[j].Name }

func main() {
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
}
