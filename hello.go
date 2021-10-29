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

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	// Get numbers from console inputs.
	fmt.Print("Enter first number:")

	firstInput, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error occurred:", err)
		return
	}

	firstNumber, err := strconv.ParseInt(strings.TrimSpace(firstInput), 10, 64)
	if err != nil {
		log.Println("error occurred:", err)
		return
	}

	fmt.Print("Enter second number:")
	secondInput, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error occurred:", err)
		return
	}

	secondNumber, err := strconv.ParseInt(strings.TrimSpace(secondInput), 10, 64)
	if err != nil {
		log.Println("error occurred:", err)
		return
	}

	// Validate numbers and get result.
	err = validateNumbers(firstNumber, secondNumber)
	if err != nil {
		log.Println("error occurred:", err)
		return
	}

	fmt.Println("Result:", getResultString(firstNumber, secondNumber))
}

// Validate numbers.
func validateNumbers(firstNumber, secondNumber int64) error {
	if firstNumber <= 0 || secondNumber <= 0 {
		return errors.New("wrong or negative number")
	}
	if secondNumber <= firstNumber {
		return errors.New("second number should be bigger then first")
	}
	if secondNumber - firstNumber > 1000 {
		return errors.New("too large range")
	}

	return nil
}

// Create result string.
func getResultString(firstNumber, seconNumber int64) string {
	var resultSlice []string
	const fizz = 3
	const buzz = 5

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
