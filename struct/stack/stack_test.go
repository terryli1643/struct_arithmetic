package stack

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestCalculate(t *testing.T) {
	pattern := "3+4/2-1"

	data := NewMyStack(len(pattern))
	symbol := NewMyStack(len(pattern))

	for _, v := range pattern {
		s := string(v)
		if s == "" {
			continue
		}

		if s == "+" || s == "-" || s == "*" || s == "/" {
			handleSymbol(data, symbol, s)
		} else {
			data.Push(s)
		}
	}

	for os := symbol.Pull(); os != nil; os = symbol.Pull() {
		o1 := data.Pull().(string)
		o2 := data.Pull().(string)
		result, err := operate(o1, o2, os.(string))
		if err != nil {
			t.Fatal(err)
		}
		data.Push(result)
	}

	t.Logf("Result is : %s", data.Pull().(string))
}

func handleSymbol(data *MyStack, symbol *MyStack, os string) (err error) {
	for {
		st := symbol.GetTop()
		if shouldCalculate(os, st) {
			o1 := data.Pull().(string)
			o2 := data.Pull().(string)
			st := symbol.Pull().(string)
			result, err := operate(o1, o2, st)
			if err != nil {
				return err
			}
			data.Push(result)
		} else {
			symbol.Push(os)
			break
		}
	}
	return nil
}

func shouldCalculate(op, st interface{}) bool {
	if st == nil {
		return false
	}
	if (op == "*" || op == "/") && (st == "+" || st == "-") {
		return false
	}
	return true
}

func operate(o1 string, o2 string, os string) (result string, err error) {
	v1, _ := strconv.ParseFloat(o1, 64)
	v2, _ := strconv.ParseFloat(o2, 64)

	switch os {
	case "+":
		return fmt.Sprint(v2 + v1), nil
	case "-":
		return fmt.Sprint(v2 - v1), nil
	case "*":
		return fmt.Sprint(v2 * v1), nil
	case "/":
		return fmt.Sprint(v2 / v1), nil
	default:
		return "", errors.New("invalid opration symbol")
	}

}
