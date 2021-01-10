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
		t.Log(s)
		if s == "" {
			continue
		}

		if s == "+" || s == "-" || s == "*" || s == "/" {
			handleSymbol(data, symbol, s)
		} else {
			data.Push(s)
		}
	}

	for {
		os := symbol.Pull()
		if os == nil {
			break
		}
		t.Log(os)
		o1 := data.Pull().(string)
		o2 := data.Pull().(string)
		operate(o1, o2, os.(string))
	}
	t.Log(data.Pull().(string))
}

func handleSymbol(data *MyStack, symbol *MyStack, os string) (err error) {
	for {
		st := symbol.GetTop()
		if shouldCalculate(os, st) {
			o1 := data.Pull().(string)
			o2 := data.Pull().(string)
			st := symbol.Pull().(string)
			fmt.Println(st)
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
	if (op == "+" || op == "-") && (st == "*" || st == "/") {
		return true
	}
	return false
}

func operate(o1 string, o2 string, os string) (result string, err error) {
	v1, _ := strconv.ParseFloat(o1, 64)
	v2, _ := strconv.ParseFloat(o2, 64)

	switch os {
	case "+":
		return fmt.Sprint(v1 + v2), nil
	case "-":
		return fmt.Sprint(v1 - v2), nil
	case "*":
		return fmt.Sprint(v1 * v2), nil
	case "/":
		return fmt.Sprint(v1 / v2), nil
	default:
		return "", errors.New("invalid opration symbol")
	}

}