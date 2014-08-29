package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func For(start int, end int) <-chan int {
	step := 1
	if end < start {
		step = -1
	}

	c := make(chan int)

	go func(start int, end int, c chan int) {
		i := start
		for {
			c <- i
			i += step
			if end < start {
				if i < end {
					close(c)
					return
				}
			} else {
				if i > end {
					close(c)
					return
				}
			}
		}
	}(start, end, c)

	return c
}

func List(items ...interface{}) []interface{} {
	return items
}

func Map(args ...interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	stringType := reflect.TypeOf("")
	for i := 0; i < len(args); i += 2 {
		key := args[i]
		value := args[i+1]

		if reflect.TypeOf(key) != stringType {
			return nil, errors.New("only string can be used as key of map")
		}

		m[key.(string)] = value
	}

	return m, nil
}

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Lowercase(s string) string {
	return strings.ToLower(s)
}

func ToInt(num interface{}) (int, error) {
	switch x := num.(type) {
	case float32:
		return int(x), nil
	case float64:
		return int(x), nil
	case int:
		return x, nil
	case uint:
		return int(x), nil
	case int8:
		return int(x), nil
	case uint8:
		return int(x), nil
	case int16:
		return int(x), nil
	case uint16:
		return int(x), nil
	case int32:
		return int(x), nil
	case uint32:
		return int(x), nil
	default:
		return 0, errors.New(fmt.Sprintf("can't convert value to int (value type: %T, value: %v)", x, x))
	}
}
