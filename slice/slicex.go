package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	array := [5]int{1, 2, 3, 4, 5}
    slice := array[1:4] // slice will be [2 3 4]
	fmt.Println("Array:", array)
	fmt.Println("Slice:", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	animals := []string{"dog", "cat", "fish"}
	fmt.Println("Animals Slice:", animals)

	// Appending to a slice
	animals = append(animals, "bird")
	fmt.Println("After appending bird:", animals)

	// Iterating over a slice
	for index, animal := range animals {
		fmt.Printf("Index: %d, Animal: %s\n", index, animal)
	}

	fmt.Println(len(animals))
	fmt.Println(cap(animals))
}