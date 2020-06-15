package main

import (
	"fmt"
	"reflect"
)

//

func main() {
	sounds()
	greet()
	animal2Test()

	m := map[string]interface{}{}
	m["one"] = 1
	m["two"] = 2.0
	m["three"] = true

	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Printf("%s is an integer\n", k)
		case float64:
			fmt.Printf("%s is a float\n", k)
		default:
			fmt.Printf("%s is %v\n", k, reflect.TypeOf(v))
		}
	}
}
