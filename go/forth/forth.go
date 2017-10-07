package forth

import (
	"errors"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

const testVersion = 2

// Forth executes a simple subset of the forth language.
// It returns a slice of integers.
func Forth(codeText []string) ([]int, error) {
	lines := lex(codeText)
	stk, err := interpretLines(lines)

	results := make([]int, stk.Len())
	// Unroll backwards
	for i := stk.Len() - 1; i >= 0; i-- {
		results[i] = stk.Pop().(int)
	}
	return results, err
}

// lex turns an array of textual code lines into individual
// tokens (numbers, operators, etc.).
func lex(s []string) (lines []string) {
	for _, line := range s {
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			lines = append(lines, token)
		}
	}
	return lines
}

// interpretLines takes an array of tokens and executes them as instructions
// (addition of numbers, assignment of variables, etc.).
func interpretLines(lines []string) (*stack.Stack, error) {
	stk := stack.New()
	userDefinedVars := make(map[string][]string)

	for i := 0; i < len(lines); i++ {
		word := lines[i]
		if num, err := strconv.Atoi(word); err == nil {
			// It's an int
			stk.Push(num)
		} else {
			switch strings.ToLower(word) {
			case "+":
				if stk.Len() == 2 {
					i2 := stk.Pop().(int)
					i1 := stk.Pop().(int)
					stk.Push(i1 + i2)
				} else {
					return stk, errors.New("found a single '+', did you mean to prepend some numbers?")
				}
			case "-":
				if stk.Len() == 2 {
					i2 := stk.Pop().(int)
					i1 := stk.Pop().(int)
					stk.Push(i1 - i2)
				} else {
					return stk, errors.New("found a single '-', did you mean to prepend some numbers?")
				}
			case "*":
				if stk.Len() == 2 {
					i2 := stk.Pop().(int)
					i1 := stk.Pop().(int)
					stk.Push(i1 * i2)
				} else {
					return stk, errors.New("found a single '*', did you mean to prepend some numbers?")
				}
			case "/":
				if stk.Len() == 2 {
					i2 := stk.Pop().(int)
					i1 := stk.Pop().(int)
					if i2 != 0 {
						stk.Push(i1 / i2)
					} else {
						return stk, errors.New("can't divide by zero!")
					}
				} else {
					return stk, errors.New("found a single '/', did you mean to prepend some numbers?")
				}
			case "dup":
				if topValue := stk.Peek(); topValue != nil {
					stk.Push(topValue)
				} else {
					return stk, errors.New("can't dup without an argument")
				}
			case "drop":
				if stk.Len() > 0 {
					stk.Pop()
				} else {
					return stk, errors.New("can't drop if there is no argument")
				}
			case "swap":
				if stk.Len() >= 2 {
					top := stk.Pop()
					next := stk.Pop()
					stk.Push(top)
					stk.Push(next)
				} else if stk.Len() < 2 {
					return stk, errors.New("can't swap unless there are at least two values")
				}
			case "over":
				if stk.Len() >= 2 {
					top := stk.Pop()
					next := stk.Pop()
					stk.Push(next)
					stk.Push(top)
					stk.Push(next)
				} else {
					return stk, errors.New("can't copy with over if there are no arguments")
				}
			case ":":
				i, err = assignStmt(userDefinedVars, i, lines)
			default:
				// TODO Look for user-defined variables
				println("Got: " + word)
			}
		}
	}
	return stk, nil
}

// Parse `: var-name value ;`
func assignStmt(userDefinedVars map[string][]string, index int, lines []string) (int, error) {
	stmtEndIndex := 0
	for i := index; i < len(lines); i++ {
		if lines[i] == ";" {
			stmtEndIndex = i
			break
		}
	}
	wordName := lines[index+1]
	wordValues := lines[index+2 : stmtEndIndex]
	userDefinedVars[wordName] = wordValues
	return index + stmtEndIndex, nil
}
