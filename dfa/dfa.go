package dfa

import "fmt"

// INFO flags if aditional information should be sent to stdout
var INFO = false

type Transition struct {
	state int
	input rune
}

/*
Delta is the table of transitions, in the form delta(srcState, input) = destState in a map

Using a map garantees we'll have unique (state, input) pairs
*/
type Delta map[Transition]int

/*
Alphabet represents an alphabet in a deterministic finite automata.
Using a map we can easly check wether the Alphabet contains or not a given character (rune).
*/
type Alphabet map[rune]bool

/*
DFA represents a deterministic finite automata.

DFA = (Q, A, d, q, F)
* Q: finite set of states
* A: Alphabet
* d: Transition function
* q: initial state
* F: accept states

*/
type DFA struct {
	allStates    []int
	alphabet     Alphabet
	InitState    int
	matrix       Delta
	finalStates  map[int]bool
	currentState int
}

// NewDFA creates a new DFA
func NewDFA(state int, isFinal bool) *DFA {
	d := &DFA{
		matrix:       make(Delta),
		alphabet:     make(map[rune]bool),
		InitState:    state,
		currentState: state,
		finalStates:  make(map[int]bool),
	}
	d.AddState(state, isFinal)
	return d
}

// AddToAlphabet add symbol to alphabet
func (dfa *DFA) AddToAlphabet(r rune) {
	if _, ok := dfa.alphabet[r]; ok {
		return
	}
	dfa.alphabet[r] = true
}

// AddState adds a new state to a DFA
func (dfa *DFA) AddState(state int, isFinal bool) int {
	if state < 0 {
		return -1
	}
	dfa.allStates = append(dfa.allStates, state)
	if isFinal {
		dfa.finalStates[state] = isFinal
	}
	return state
}

// AddTransition adds a new transition from source state to destination state using input
func (dfa *DFA) AddTransition(srcState int, input rune, destState int) int {
	if _, ok := dfa.alphabet[input]; !ok {
		return -1
	}
	delta := Transition{srcState, input}
	if _, ok := dfa.matrix[delta]; ok {
		return -2
	}
	dfa.matrix[delta] = destState
	return destState
}

// CheckWord checks if a word belongs can be produced by the grammar
func (dfa *DFA) CheckWord(word string) bool {
	defer dfa.Reset()
	// lexycal analisys
	if !dfa.Lexycal(word) {
		return false
	}
	if dfa.run(word) {
		if dfa.Accept() {
			if INFO {
				fmt.Printf(" %d é estado de aceitação\n", dfa.currentState)
			}
			return true
		}
		if INFO {
			fmt.Printf(" %d não é estado de aceitação\n", dfa.currentState)
		}
	}
	return false
}

// Lexycal runs a lexycal analises to check wether "word" is composed of only acceptable symbols
func (dfa *DFA) Lexycal(word string) bool {
	for _, v := range word {
		// fmt.Printf("alphabet[%d]\n", v)
		if _, ok := dfa.alphabet[v]; !ok {
			return false
		}
	}
	return true
}

func (dfa *DFA) transitionInput(input rune) int {
	delta := Transition{
		state: dfa.currentState,
		input: input,
	}
	dest, ok := dfa.matrix[delta]
	if INFO {
		fmt.Printf(" delta(%d, %c) = %d\n", delta.state, delta.input, dest)
	}
	if !ok {
		return -1
	}
	dfa.currentState = dest
	return dfa.currentState
}

func (dfa *DFA) run(word string) bool {
	for _, v := range word {
		if res := dfa.transitionInput(v); res == -1 {
			return false
		}
	}
	return true
}

// Accept checks if the last state for the input is an accepting state
func (dfa *DFA) Accept() bool {
	_, ok := dfa.finalStates[dfa.currentState]
	return ok
}

// Reset sets the DFA back to initial state
func (dfa *DFA) Reset() {
	dfa.currentState = dfa.InitState
}
