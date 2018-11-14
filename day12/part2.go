package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func count(val interface{}) float64 {
	sum := 0.0
	switch reflect.TypeOf(val).String() {
	case "float64":
		sum += val.(float64)
	case "[]interface {}":
		array := val.([]interface{})
		for i := 0; i < len(array); i++ {
			sum += count(array[i])
		}
	case "map[string]interface {}":
		object := val.(map[string]interface{})
		skip := false
		for _, v := range object {
			if v == "red" {
				skip = true
				break
			}
		}
		if !skip {
			for _, v := range object {
				sum += count(v)
			}
		}
	}
	return sum
}

func main() {
	input, _ := ioutil.ReadFile("input")
	var data interface{}
	json.Unmarshal(input, &data)
	fmt.Println(count(data))
}
