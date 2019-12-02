package main

import (
"fmt"
"runtime"
)

func main() {
	for i := 0 ; i< 10; i++ {
		test(i)
	}
}

func test(skip int) {
	call(skip)
}

func call(skip int) {
	pc, file, line, ok := runtime.Caller(skip)
	pcName := runtime.FuncForPC(pc).Name()  //获取函数名
	fmt.Println(fmt.Sprintf("%d -> %v   %s   %d   %t   %s",skip, pc, file, line, ok, pcName))

}