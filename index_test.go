package main

import (
	"fmt"
	"github.com/akrylysov/simplefts/calc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	idx := make(index)

	assert.Nil(t, idx.search("foo"))
	assert.Nil(t, idx.search("donut"))

	idx.add([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	assert.Nil(t, idx.search("a"))
	assert.Equal(t, idx.search("donut"), []int{1})
	assert.Equal(t, idx.search("DoNuts"), []int{1})
	assert.Equal(t, idx.search("glass"), []int{1})

	idx.add([]document{{ID: 2, Text: "donut is a donut"}})
	assert.Nil(t, idx.search("a"))
	assert.Equal(t, idx.search("donut"), []int{1, 2})
	assert.Equal(t, idx.search("DoNuts"), []int{1, 2})
	assert.Equal(t, idx.search("glass"), []int{1})
}

func TestCalcModule(t *testing.T) {
	fmt.Println(calc.Age)
	calc.PrintWelcome()

	person := calc.Person{
		Title:   "title",
		Address: "address",
		Name:    "name",
		ID:      234,
	}
	fmt.Println(person)
}
