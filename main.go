package main

import (
	"reflect"
)

//Reverse reverses a slice.
var Reverse func(slice interface{}) = func(slice interface{}) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() == reflect.Ptr {
		sliceValue = sliceValue.Elem()
	}
	switch sliceValue.Kind() {
	case reflect.Slice:
		j := 0
		for i := sliceValue.Len() - 1; i >= sliceValue.Len()/2; i-- {
			a := sliceValue.Index(i)
			b := sliceValue.Index(j)
			switch a.Kind() {
			case reflect.Int:
				av := a.Int()
				bv := b.Int()
				sliceValue.Index(i).SetInt(bv)
				sliceValue.Index(j).SetInt(av)
			case reflect.String:
			case reflect.Float32:
			}
			j += 1
		}
	}
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
