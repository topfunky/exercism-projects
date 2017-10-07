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
		var err error
		err = interpretWord(word, &i, lines, stk, userDefinedVars)
		if err != nil {
			return stk, err
		}
	}
	return stk, nil
}

func interpretWord(word string, i *int, lines []string, stk *stack.Stack, userDefinedVars map[string][]string) error {
	word = strings.ToLower(word)
	if num, err := strconv.Atoi(word); err == nil {
		// an int
		stk.Push(num)
	} else if statements, ok := userDefinedVars[word]; ok {
		// user-defined words
		for _, stmt := range statements {
			interpretWord(stmt, i, lines, stk, userDefinedVars)
		}
	} else {
		// built-in keywords and operators
		switch word {
		case "+":
			if stk.Len() == 2 {
				i2 := stk.Pop().(int)
				i1 := stk.Pop().(int)
				stk.Push(i1 + i2)
			} else {
				return errors.New("found a single '+', did you mean to prepend some numbers?")
			}
		case "-":
			if stk.Len() == 2 {
				i2 := stk.Pop().(int)
				i1 := stk.Pop().(int)
				stk.Push(i1 - i2)
			} else {
				return errors.New("found a single '-', did you mean to prepend some numbers?")
			}
		case "*":
			if stk.Len() == 2 {
				i2 := stk.Pop().(int)
				i1 := stk.Pop().(int)
				stk.Push(i1 * i2)
			} else {
				return errors.New("found a single '*', did you mean to prepend some numbers?")
			}
		case "/":
			if stk.Len() == 2 {
				i2 := stk.Pop().(int)
				i1 := stk.Pop().(int)
				if i2 != 0 {
					stk.Push(i1 / i2)
				} else {
					return errors.New("can't divide by zero")
				}
			} else {
				return errors.New("found a single '/', did you mean to prepend some numbers?")
			}
		case "dup":
			if topValue := stk.Peek(); topValue != nil {
				stk.Push(topValue)
			} else {
				return errors.New("can't dup without an argument")
			}
		case "drop":
			if stk.Len() > 0 {
				stk.Pop()
			} else {
				return errors.New("can't drop if there is no argument")
			}
		case "swap":
			if stk.Len() >= 2 {
				top := stk.Pop()
				next := stk.Pop()
				stk.Push(top)
				stk.Push(next)
			} else if stk.Len() < 2 {
				return errors.New("can't swap unless there are at least two values")
			}
		case "over":
			if stk.Len() >= 2 {
				top := stk.Pop()
				next := stk.Pop()
				stk.Push(next)
				stk.Push(top)
				stk.Push(next)
			} else {
				return errors.New("can't copy with over if there are no arguments")
			}
		case ":":
			if err = assignStmt(userDefinedVars, i, lines); err != nil {
				return err
			}
		default:
			return errors.New(word + " is not a built-in or recognized user-defined word")
		}
	}
	return nil
}

// Parse `: var-name value ;`
func assignStmt(userDefinedVars map[string][]string, index *int, lines []string) error {
	stmtEndIndex := 0
	for i := *index; i < len(lines); i++ {
		if lines[i] == ";" {
			stmtEndIndex = i
			break
		}
	}
	wordName := lines[*index+1]
	wordValues := lines[*index+2 : stmtEndIndex]
	if _, err := strconv.Atoi(wordName); err == nil {
		return errors.New("numbers can't be redefined as user-defined words")
	}
	userDefinedVars[wordName] = wordValues
	*index = stmtEndIndex
	return nil
}
