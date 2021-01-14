package recursive

import (
	"runtime"
	"strings"
)

func RecursiveDepthChecker(max int) bool {
	//注意，这里我们跳过了本函数自己的栈，找到父调用的函数名
	caller, _, _, _ := runtime.Caller(1)
	currentFuncName := runtime.FuncForPC(caller).Name()
	stack := make([]byte, 65535*1) //max 1MB stack traceback
	//由于Golang Runtime中很多关于栈的函数未导出，无法使用。因此使用最肮脏的字符串检测方法
	runtime.Stack(stack, false)
	start := 0
	depth := 0
	for {
		count := strings.Index(string(stack[start:]), currentFuncName)
		if count >= 0 {
			start += count + len(currentFuncName)
			depth++
		} else {
			break
		}
	}

	if depth > max {
		return false
	}

	return true
}
