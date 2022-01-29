package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("valor 1:")
	input1, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	number1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		panic(err)
	}

	fmt.Print("valor 2:")
	input2, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	number2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		panic(err)
	}

	sum := number1 + number2
	rsum := math.Round(sum*100) / 100
	fmt.Println("Suma:", rsum)

}
