package main

import "fmt"

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }

	// if age < 18 {
	// 	return false
	// }
	// return true

	// switch age {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }
	// return false

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

	// switch {
	// case age < 18:
	// 	return false
	// case age == 18:
	// 	return true
	// case age > 50:
	// 	return false
	// }
	// return false
}

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	fmt.Println(canIDrink(14))

	a := 2
	b := &a
	*b = 20
	fmt.Println(a, *b)

	array := [5]string{"a", "b", "c", "d"}
	array[2] = "z"
	array[4] = "e"

	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")

	fmt.Println(array)
	fmt.Println(names)

	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}

	favFood := []string{"kimchi", "ramen"}
	nico2 := person{"nico", 18, favFood}
	fmt.Println(nico2)
	fmt.Println(nico2.name)

	nico3 := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico3)
}
