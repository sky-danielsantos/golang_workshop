package main

import (
	"errors"
	"fmt"
	"log"
)

// structured errors

type DivisionByZeroError struct {
	a, b int
}

type DivisionByNegativeError struct {
	a, b int
}

func (e *DivisionByZeroError) getA() int {
	return e.a
}
func (e *DivisionByZeroError) getB() int {
	return e.b
}

func (e *DivisionByZeroError) Error() string {
	return fmt.Sprintf("cannot devide number %d by 0", e.a)
}

func (e *DivisionByNegativeError) getA() int {
	return e.a
}
func (e *DivisionByNegativeError) getB() int {
	return e.b
}

func (e *DivisionByNegativeError) Error() string {
	return fmt.Sprintf("error cannot divide %d by negative number %d", e.a, e.b)
}

func divErrorStructured(a int, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionByZeroError{a: a, b: b}
	}
	if b < 0 {
		return 0, &DivisionByNegativeError{a: a, b: b}
	}
	return a / b, nil
}

func handleError(err error) {
	var divisionNegativeErr *DivisionByNegativeError
	var divisionByZeroErr *DivisionByZeroError
	switch {
	case errors.As(err, &divisionByZeroErr):
		fmt.Println("error division by 0 is not allowed: ", err)
	case errors.As(err, &divisionNegativeErr):
		fmt.Println("error negative division is not allowed: ", err)
	default:
		log.Println("unexpected error on division operation: ", err)
	}
}
func structuredError() {
	if _, err := divErrorStructured(10, 0); err != nil {
		handleError(err)
	}
	if _, err := divErrorStructured(10, -1); err != nil {
		handleError(err)
	}
}

// Error Kinds

var ErrDevideByZero = errors.New("divide by zero")

func divErrorKind(a int, b int) (int, error) {
	if b == 0 {
		return 0, ErrDevideByZero
	}
	return a / b, nil
}

func errorKind() {
	if _, err := divErrorKind(10, 0); err != nil {
		switch {
		case errors.Is(err, ErrDevideByZero):
			log.Println("error kind is devide by zero with zerror message: ", err)
		default:
			log.Println("unexpected error on division operation: ", err)
		}
	}
}

// Simple Error Creation
func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

func simpleErrorCreation() {
	if _, err := div(10, 0); err != nil {
		log.Println("simpleErrorCreation: generated an error as expected: ", err)
	}
	value, _ := div(10, 2)
	log.Println("division value: ", value)
}

func main() {
	//simpleErrorCreation()
	//errorKind()
	//structuredError()
}
