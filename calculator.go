package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go calculator!")
	cmd := readLine("Enter command: [a]tambah, [s]ubtract, [m]ultiply, [d]ivide: ")
	fmt.Print(cmd)
	if cmd == "a" {
		num1, num2 := getUserNumbers()
		result := num1 + num2
		fmt.Println(fmt.Sprintf("%d", result))
	} else if cmd == "s" {
		num1, num2 := getUserNumbers()
		result := num1 - num2
		fmt.Println(fmt.Sprintf("%d", result))
	} else if cmd == "m" {
		num1, num2 := getUserNumbers()
		result := num1 * num2
		fmt.Println(fmt.Sprintf("%d", result))
	} else if cmd == "d" {
		num1, num2 := getUserNumbers()
		result := float32(num1) / float32(num2)
		fmt.Println(fmt.Sprintf("%f", result))
	} else {
		fmt.Println("Invalid input")
	}
	readLine("press 'enter' to exit")
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	num1String := readLine("angka pertama: ")
	num1, _ := strconv.Atoi(num1String)
	num2String := readLine("angka kedua: ")
	num2, _ := strconv.Atoi(num2String)
	return num1, num2
}
