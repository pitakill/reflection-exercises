package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type example struct {
	message string
}

type exampleInner struct {
	message string
}

func (e *example) Test() {
	fmt.Println("Test")
}

func (e *example) TestWithParam(s string) {
	fmt.Println(s)
}

func (e *example) TestWithParamAndReturn(s string) string {
	return strings.ToLower(s)
}

func (e *example) TestWithInnerStructAndParamAndReturn(s string) *exampleInner {
	return &exampleInner{
		message: e.message + " " + s,
	}
}

func (e *example) TestWithParamInnerStruct(i exampleInner) *example {
	return &example{
		message: e.message + " " + i.message,
	}
}

type typeRegister map[string]reflect.Type

func (t typeRegister) Set(i interface{}) {
	typ := reflect.TypeOf(i).Elem()
	t[typ.Name()] = typ
}

func (t typeRegister) Get(name string) (interface{}, error) {
	if typ, ok := t[name]; ok {
		return reflect.New(typ).Elem().Interface(), nil
	}

	return nil, errors.New("not valid type registered: " + name)
}

var typeRegistry = make(typeRegister)

func init() {
	typeRegistry.Set(new(exampleInner))
	runtime.GC()
}
