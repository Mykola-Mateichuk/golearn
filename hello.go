package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const fizz = 3
const buzz = 5

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	// Get numbers from console inputs.
	fmt.Print("Enter first number:")
	firstInput, _ := reader.ReadString('\n')
	firstNumber, _ := strconv.ParseInt(strings.TrimSpace(firstInput), 10, 64)

	fmt.Print("Enter second number:")
	secondInput, _ := reader.ReadString('\n')
	secondNumber, _ := strconv.ParseInt(strings.TrimSpace(secondInput), 10, 64)

	// Validate numbers and get result.
	err := validateNumbers(firstNumber, secondNumber)
	if err != nil {
		log.Println("error occurred:", err)
		return
	}

	fmt.Println("Result:", getResultString(firstNumber, secondNumber))
}

// Validate numbers.
func validateNumbers(firstNumber, secondNumber int64) error {
	var err error = nil

	if firstNumber <= 0 || secondNumber <= 0 {
		panic("Wrong or negative number")
	}
	if secondNumber <= firstNumber {
		panic("Second number should be bigger then first")
	}
	if secondNumber - firstNumber > 1000 {
		err = errors.New("Too large range")
	}

	return err
}

// Create result string.
func getResultString(firstNumber, seconNumber int64) string {
	var resultSlice []string

	// Check for fizz, buzz and create result string.
	for i := firstNumber; i <= seconNumber; i++ {
		if i % fizz == 0 && i % buzz == 0 {
			resultSlice = append(resultSlice, "fizzbuzz")
		} else if i % fizz == 0 {
			resultSlice = append(resultSlice, "fizz")
		} else if i % buzz == 0 {
			resultSlice = append(resultSlice, "buzz")
		} else {
			resultSlice = append(resultSlice, strconv.FormatInt(i, 10))
		}
	}

	return strings.Join(resultSlice, ", ")
}
