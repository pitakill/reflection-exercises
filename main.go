package main

import (
	"fmt"
	"reflect"
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

func (e *example) TestWithParamInnerStruct(i *exampleInner) *example {
	return &example{
		message: e.message + " " + i.message,
	}
}

// Examples
func callTestNormally(e *example) {
	e.Test()
}

func callTestWithReflection(e *example) {
	m := "Test"

	// Get the type of the instance
	// Also get the method by name
	element := reflect.ValueOf(e).MethodByName(m)

	// Call the method
	// Works well with a function without parameters and without return of the
	// method
	element.Call([]reflect.Value{})
}

func callTestWithParamsNormally(e *example, s string) {
	e.TestWithParam(s)
}

func callTestWithParamsWithReflection(e *example, s string) {
	m := "TestWithParam"

	element := reflect.ValueOf(e).MethodByName(m)
	element.Call([]reflect.Value{reflect.ValueOf(s)})
}

func callTestWithParamAndReturnNormally(e *example, s string) string {
	return e.TestWithParamAndReturn(s)
}

func callTestWithParamAndReturnWithReflection(e *example, s string) string {
	m := "TestWithParamAndReturn"

	element := reflect.ValueOf(e).MethodByName(m)
	called := element.Call([]reflect.Value{reflect.ValueOf(s)})

	// We know how many elements returns the function from the string m and we
	// know the type. So we can expect only one element of type string
	return called[0].Interface().(string)
}

func callTestWithInnerStructAndParamAndReturnNormally(e *example, s string) *exampleInner {
	return e.TestWithInnerStructAndParamAndReturn(s)
}

func callTestWithInnerStructAndParamAndReturnWithReflection(e *example, s string) *exampleInner {
	m := "TestWithInnerStructAndParamAndReturn"

	element := reflect.ValueOf(e).MethodByName(m)
	called := element.Call([]reflect.Value{reflect.ValueOf(s)})

	if casted, ok := called[0].Interface().(*exampleInner); ok {
		return casted
	}

	return &exampleInner{}
}

func callTestWithParamInnerStructNormally(e *example, i *exampleInner) *example {
	return e.TestWithParamInnerStruct(i)
}

func callTestWithParamInnerStructWithReflection(e *example, i *exampleInner) *example {
	m := "TestWithParamInnerStruct"

	element := reflect.ValueOf(e).MethodByName(m)
	called := element.Call([]reflect.Value{reflect.ValueOf(i)})

	if casted, ok := called[0].Interface().(*example); ok {
		return casted
	}

	return &example{}
}

func main() {
	//e := new(example)

	//callTestNormally(e)
	//callTestWithReflection(e)
	//callTestWithParamsNormally(e, "Hello World!")
	//callTestWithParamsWithReflection(e, "Hello World!")
	//t := callTestWithParamAndReturnNormally(e, "HELLO MY FRIEND!")
	//fmt.Println(t)
	//t := callTestWithParamAndReturnWithReflection(e, "HELLO MY FRIEND!")
	//fmt.Println(t)

	e := &example{
		message: "Hello",
	}

	//t := callTestWithInnerStructAndParamAndReturnNormally(e, "World!")
	//fmt.Println(t.message)
	//t := callTestWithInnerStructAndParamAndReturnWithReflection(e, "world!")
	//fmt.Println(t.message)

	i := &exampleInner{
		message: "World!",
	}

	t := callTestWithParamInnerStructWithReflection(e, i)
	fmt.Println(t.message)
}
