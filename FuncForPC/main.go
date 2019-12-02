package main

import (
	"log"
	"reflect"
	"runtime"
)

type Demo struct {
}

func (this *Demo) Hello() {

}

func main() {
	demo := new(Demo)
	pc := reflect.ValueOf(demo.Hello).Pointer()
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())
	log.Println(f.FileLine(pc))
}