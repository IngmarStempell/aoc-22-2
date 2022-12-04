package main_

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main_() {

	fmt.Println("Reading rock paper scissors")
	inputHandle, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File not found")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(inputHandle)
	line, err := reader.ReadString('\n')
	i := 0
	sum := 0

	for line != "" {
		if err != nil {
			fmt.Println("Something went wrong reading the input")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		sum += matchValue(split(line))
		i++
		line, err = reader.ReadString('\n')
	}

	fmt.Println("Read " + strconv.Itoa(i) + " lines")
	fmt.Println("Got combined value of: " + strconv.Itoa(sum))
}

func matchValue(left string, right string) int {
	if isValidLeft(left) && isValidRight(right) {
		print("(" + left + "," + right + ") ")
		if isDraw(left, right) {
			print("Draw: ")
			println(getDrawValue(left, right))
			return getDrawValue(left, right)
		}
		if isWin(left, right) {
			print("Win: ")
			println(getWinValue(left, right))
			return getWinValue(left, right)
		}
		print("Loss: ")
		println(getLossValue(left, right))
		return getLossValue(left, right)

	}
	println("Got funky values for left or right")
	return 0
}

func getValueForAction(action string) int {
	switch action {
	case "A":
	case "X":
		return 1
	case "B":
	case "Y":
		return 2
	case "C":
	case "Z":
		return 3
	}
	return 10000000
}

func getWinValue(left, right string) int {
	return getValueForAction(right) + 6

}

func getDrawValue(left, right string) int {
	return getValueForAction(right) + 3

}

func getLossValue(left, right string) int {
	return getValueForAction(right) + 0
}

func isWin(left, right string) bool {
	return left == "A" && right == "Y" ||
		left == "B" && right == "Z" ||
		left == "C" && right == "X"

}

func isDraw(left, right string) bool {
	return left == "A" && right == "X" ||
		left == "B" && right == "Y" ||
		left == "C" && right == "Z"
}

func isValidRight(right string) bool {
	return right == "X" ||
		right == "Y" ||
		right == "Z"
}

func isValidLeft(left string) bool {
	return left == "A" ||
		left == "B" ||
		left == "C"
}

func split(line string) (left string, right string) {
	slices := strings.Split(line, " ")
	return slices[0], strings.Trim(slices[1], "\n")
}
