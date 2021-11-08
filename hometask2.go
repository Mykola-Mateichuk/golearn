// Here can be some copyrights.

// Package hometask2 Package main provide functionality for learn purposes. It will contain
// all home task works.
//
// Notice that it also can contain some additional examples of code
// for learning needs.
//
// If you have any suggestion or comment, please feel free to ping me
// directly in teams or by email!
//
// Home task 1 examples
//
// Function validateNumbers check if values in right range.
//  err = validateNumbers(firstNumber, secondNumber)
//	if err != nil {
//		log.Println("error occurred:", err)
//		return
//	}
//
// Function getResultString provide result string contains replacement for numbers
// which can be divided by 3, 5 and 3 and 5.
//  fmt.Println("Result:", getResultString(firstNumber, secondNumber))
//
// Home task 2 examples
//
// You can use HashPassword function next way:
//  func main() {
//      fmt.Println(hasher.HashPassword("123"))
//      // Output: $2a$14$O2lOPk37oLymkT3irkDll.T4zn37aaE84wzqLorT0rcqtSF2xcXQG
//  }
//
// You can use CheckPasswordHash function next way:
//  func main() {
//      fmt.Println(hasher.CheckPasswordHash("123", "$2a$14$O2lOPk37oLymkT3irkDll.T4zn37aaE84wzqLorT0rcqtSF2xcXQG"))
//      // Output: true
//  }
//
// By Mykola Mateichuk
package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/pkg/hasher"
	"log"
	"os"
	"strconv"
	"strings"
)

// Contain entry point for this project.
func hometask2() {
	// Test hasher package.
	hash, _ := hasher.HashPassword("123")
	fmt.Println("Password 123 has next hash:", hash)
	isHashValid := hasher.CheckPasswordHash("123", "$2a$14$O2lOPk37oLymkT3irkDll.T4zn37aaE84wzqLorT0rcqtSF2xcXQG")
	fmt.Println("Hash is valid:", isHashValid)

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

// Validate numbers for correct values.
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
