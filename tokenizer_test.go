package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		text   string
		tokens []string
	}{
		{
			text:   "",
			tokens: []string{},
		},
		{
			text:   "a",
			tokens: []string{"a"},
		},
		{
			text:   "small wild,cat!",
			tokens: []string{"small", "wild", "cat"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.text, func(st *testing.T) {
			assert.EqualValues(st, tc.tokens, tokenize(tc.text))
		})
	}
}

func TestFileExists(t *testing.T) {
	filename := "doc.json"

	// 使用 os.Stat 检查文件是否存在
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", filename)
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("File %s exists\n", filename)
	}
}
