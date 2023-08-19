package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Entry number, please: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input) // Remove leading and trailing whitespaces

	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	result := calculateFactorial(number)
	fmt.Printf("Factorial numbers %d: %d\n", number, result)
}

func calculateFactorial(n int) int {
	if n <= 0 {
		return 1
	}

	return n * calculateFactorial(n-1)
}
