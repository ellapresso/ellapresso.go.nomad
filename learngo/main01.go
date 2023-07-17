package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func multiply2(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	defer fmt.Println("I'm done")
	fmt.Println(words)
}

func superAdd(numbers ...int) (total int) {
	fmt.Println(numbers)
	total = 0
	for _, number := range numbers {
		total += number
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return
}

func main01() {
	name := "hello"
	name = "world"

	fmt.Println(name)

	fmt.Println(multiply(2, 3))

	fmt.Println(lenAndUpper("hello"))
	totalLen, upperName := lenAndUpper("ellapresso")
	totalLen2, _ := lenAndUpper("ellapresso2")
	fmt.Println(totalLen, upperName)
	fmt.Println(totalLen2)

	repeatMe("ellapresso", "ellapresso2", "ellapresso3", "ellapresso4", "ellapresso5")
	len, _ := lenAndUpper2("ellapressoooo")
	fmt.Println(len)

	fmt.Println(superAdd(11, 2, 31, 4, 5, 6))
}
