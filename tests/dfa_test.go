package tests

import (
	"fmt"
	"testing"
)

var auto = dfa.NewDFA(0, false)

func TestCreatedDFA(t *testing.T) {
	if auto.InitState != 0 {
		t.Error("Error! InitState:", auto.InitState)
	}
}

func TestBasicDFA(t *testing.T) {
	var word string

	auto.AddState(1, true)

	auto.AddToAlphabet('a')
	auto.AddToAlphabet('b')

	auto.AddTransition(0, 'a', 1)
	auto.AddTransition(0, 'b', 0)
	auto.AddTransition(1, 'a', 0)
	auto.AddTransition(1, 'b', 1)

	word = "ab"
	if ok := auto.CheckWord(word); !ok {
		t.Error("Expected 'true', got ", ok)
	}

	word = "abbbbbbbbb"
	if ok := auto.CheckWord(word); !ok {
		t.Error("Expected 'true', got ", ok)
	}

	word = "bbbbbbbba"
	if ok := auto.CheckWord(word); !ok {
		t.Error("Expected 'true', got ", ok)
	}

	word = "abbaa"
	if ok := auto.CheckWord(word); !ok {
		t.Error("Expected 'true', got ", ok)
	}
}

func TestDFA(t *testing.T) {

	auto = dfa.NewDFA(0, false)

	auto.AddState(1, true) // f
	auto.AddState(2, false)
	auto.AddState(3, true) // f
	auto.AddState(4, false)
	auto.AddState(5, true) // f
	auto.AddState(6, false)

	auto.AddToAlphabet('a')
	auto.AddToAlphabet('b')
	auto.AddToAlphabet('c')

	// 0
	auto.AddTransition(0, 'a', 0)
	auto.AddTransition(0, 'b', 1)
	auto.AddTransition(0, 'c', 6)

	// 1
	auto.AddTransition(1, 'a', 2)
	auto.AddTransition(1, 'b', 6)
	auto.AddTransition(1, 'c', 3)

	// 2
	auto.AddTransition(2, 'a', 2)
	auto.AddTransition(2, 'b', 1)
	auto.AddTransition(2, 'c', 3)

	// 3
	auto.AddTransition(3, 'a', 6)
	auto.AddTransition(3, 'b', 4)
	auto.AddTransition(3, 'c', 3)

	// 4
	auto.AddTransition(4, 'a', 6)
	auto.AddTransition(4, 'b', 4)
	auto.AddTransition(4, 'c', 5)

	// 5
	auto.AddTransition(5, 'a', 6)
	auto.AddTransition(5, 'b', 4)
	auto.AddTransition(5, 'c', 5)

	// 6
	auto.AddTransition(6, 'a', 6)
	auto.AddTransition(6, 'b', 6)
	auto.AddTransition(6, 'c', 6)

	word := "ab"
	if ok := auto.CheckWord(word); !ok {
		t.Error("Expected 'true', got ", ok)
	} else {
		fmt.Println("++ ACCEPTED ++")
	}

	word = "ababccbc"
	if ok := auto.CheckWord(word); !ok {
		t.Error("Expected 'true', got ", ok)
	} else {
		fmt.Println("++ ACCEPTED ++")
	}
}
