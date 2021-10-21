package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	j = wodJson{
		Word:        "test",
		Definitions: []child{d1, d2},
		Examples:    []child{e1, e2},
		Note:        "test note",
	}

	d1 = child{Text: "def1"}
	d2 = child{Text: "def2"}

	e1 = child{Text: "ex1"}
	e2 = child{Text: "ex2"}
)

func TestString(t *testing.T) {
	expected := strings.Join([]string{
		"test - def1 / def2",
		`"ex1"`,
		`"ex2"`,
		"test note",
	}, "\n")

	assert.Equal(t, j.String(), expected)
}

func TestDefinition(t *testing.T) {
	expected := "def1 / def2"

	assert.Equal(t, j.Definition(), expected)
}

func TestExample(t *testing.T) {
	expected := strings.Join([]string{
		`"ex1"`,
		`"ex2"`,
	}, "\n")

	assert.Equal(t, j.Example(), expected)
}
