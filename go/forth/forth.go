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
	// pop all values off the stack and return as slice
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
	userDefinedWords := make(map[string][]string)

	for i := 0; i < len(lines); i++ {
		word := lines[i]
		var err error
		err = interpretWord(word, &i, lines, stk, userDefinedWords)
		if err != nil {
			return stk, err
		}
	}
	return stk, nil
}

// interpretWord reads each token and executes it appropriately.
func interpretWord(word string, i *int, lines []string, stk *stack.Stack, userDefinedWords map[string][]string) error {
	var num int
	var err error
	word = strings.ToLower(word)
	if num, err = strconv.Atoi(word); err == nil {
		// an int
		stk.Push(num)
	} else if statements, ok := userDefinedWords[word]; ok {
		// user-defined words
		for _, stmt := range statements {
			interpretWord(stmt, i, lines, stk, userDefinedWords)
		}
	} else {
		// built-in keywords and operators
		switch word {
		case "+":
			if err = plusOp(stk); err != nil {
				return err
			}
		case "-":
			if err = minusOp(stk); err != nil {
				return err
			}
		case "*":
			if err = multiplyOp(stk); err != nil {
				return err
			}
		case "/":
			if err = divideOp(stk); err != nil {
				return err
			}
		case "dup":
			if err = dupOp(stk); err != nil {
				return err
			}
		case "drop":
			if err = dropOp(stk); err != nil {
				return err
			}
		case "swap":
			if err = swapOp(stk); err != nil {
				return err
			}
		case "over":
			if err = overOp(stk); err != nil {
				return err
			}
		case ":":
			if err = assignStmt(userDefinedWords, i, lines); err != nil {
				return err
			}
		default:
			return errors.New(word + " is not a built-in or recognized user-defined word")
		}
	}
	return nil
}

// assignStmt parses user-defined words in the format `: var-name value ;`
func assignStmt(userDefinedWords map[string][]string, index *int, lines []string) error {
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
	userDefinedWords[wordName] = wordValues
	*index = stmtEndIndex
	return nil
}

func plusOp(stk *stack.Stack) error {
	if stk.Len() == 2 {
		i2 := stk.Pop().(int)
		i1 := stk.Pop().(int)
		stk.Push(i1 + i2)
	} else {
		return errors.New("found a single '+', did you mean to prepend some numbers?")
	}
	return nil
}

func minusOp(stk *stack.Stack) error {
	if stk.Len() == 2 {
		i2 := stk.Pop().(int)
		i1 := stk.Pop().(int)
		stk.Push(i1 - i2)
	} else {
		return errors.New("found a single '-', did you mean to prepend some numbers?")
	}
	return nil
}

func multiplyOp(stk *stack.Stack) error {
	if stk.Len() == 2 {
		i2 := stk.Pop().(int)
		i1 := stk.Pop().(int)
		stk.Push(i1 * i2)
	} else {
		return errors.New("found a single '*', did you mean to prepend some numbers?")
	}
	return nil
}

func divideOp(stk *stack.Stack) error {
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
	return nil
}

func dupOp(stk *stack.Stack) error {
	if topValue := stk.Peek(); topValue != nil {
		stk.Push(topValue)
	} else {
		return errors.New("can't dup without an argument")
	}
	return nil
}

func dropOp(stk *stack.Stack) error {
	if stk.Len() > 0 {
		stk.Pop()
	} else {
		return errors.New("can't drop if there is no argument")
	}
	return nil
}

func swapOp(stk *stack.Stack) error {
	if stk.Len() >= 2 {
		top := stk.Pop()
		next := stk.Pop()
		stk.Push(top)
		stk.Push(next)
	} else if stk.Len() < 2 {
		return errors.New("can't swap unless there are at least two values")
	}
	return nil
}

func overOp(stk *stack.Stack) error {
	if stk.Len() >= 2 {
		top := stk.Pop()
		next := stk.Pop()
		stk.Push(next)
		stk.Push(top)
		stk.Push(next)
	} else {
		return errors.New("can't copy with over if there are no arguments")
	}
	return nil
}
