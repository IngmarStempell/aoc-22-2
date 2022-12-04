package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

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
			neededAction := getDrawAction(left)
			value := getDrawValue(left, neededAction)
			println(value)
			return value
		}
		if isWin(left, right) {
			print("Win: ")
			neededAction := getWinAction(left)
			value := getWinValue(left, neededAction)
			println(value)
			return value
		}
		print("Loss: ")
		neededAction := getLossAction(left)
		value := getLossValue(left, neededAction)
		println(value)
		return value

	}
	println("Got funky values for left or right")
	return 0
}

func getLossAction(action string) string {
	switch action {
	case "A":
		return "Z"
	case "B":
		return "X"
	case "C":
		return "Y"
	}
	panic("unimplemented")
}

func getWinAction(action string) string {
	switch action {
	case "A":
		return "Y"
	case "B":
		return "Z"
	case "C":
		return "X"
	}
	panic("unimplemented")
}

func getDrawAction(action string) string {
	switch action {
	case "A":
		return "X"
	case "B":
		return "Y"
	case "C":
		return "Z"
	}
	panic("unimplemented")
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
	return right == "Z"

}

func isDraw(left, right string) bool {
	return right == "Y"

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
