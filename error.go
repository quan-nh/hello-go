package main

import (
	"errors"
	"fmt"
)

// fn return error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

// return error with fmt.Errorf
func f2(arg int) (int, error) {
	if arg >= 42 {
		return -1, fmt.Errorf("can't work with %v", arg)
	}
	return arg + 3, nil
}

// define our error by impl error interface
// type error interface {
// 	Error() string
// }
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f3(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// TODO unwrap

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 43} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// get data in custom error
	_, err := f3(42)
	if e, ok := err.(*argError); ok {
		fmt.Println(e.arg)
		fmt.Println(e.prob)
	}
	// equivalent in go 1.13
	var e *argError
	if errors.As(err, &e) {
		fmt.Println(e.arg)
		fmt.Println(e.prob)
	}
}
