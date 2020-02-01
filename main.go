package main

import (
	"reflect"
)

func MakeMap(fpt interface{}) {
	fnV := reflect.ValueOf(fpt).Elem()
	fnI := reflect.MakeFunc(fnV.Type(), implMap)
	fnV.Set(fnI)
}

//TODO:completes implMap function.
var implMap func([]reflect.Value) []reflect.Value = func(values []reflect.Value) []reflect.Value {
	var output []reflect.Value
	if len(values) <= 1 {
		return nil
	}
	arg0 := values[0]
	arg1 := values[1]
	switch arg1.Kind() {
	case reflect.Map:
		keys := arg1.MapKeys()
		for _, key := range keys {
			v := arg1.MapIndex(key)
			res := arg0.Call([]reflect.Value{v})
			arg1.SetMapIndex(key, res[0])
		}
	case reflect.Slice:
		for i := 0; i < arg1.Len(); i++ {
			ret := arg0.Call([]reflect.Value{arg1.Index(i)})
			arg1.Index(i).Set(ret[0])
		}
	}
	output = append(output, arg1)
	return output
}

func main() {

	println("It is said that Go has no generics.\nHowever we have many other ways to implement a generics like library if less smoothly,one is reflect.MakeFunc.\nUnderscore is a very useful js library,and now let's implement part of it-map,it will help you to understand how reflect works.\nPlease finish the 'implMap' function and pass the test.")
}
