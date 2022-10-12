package main

import "fmt"

func main() {
	//another way --> colors := make(map[string]string)
	colors := map[string]string{
		"blue":  "BLUE",
		"green": "GREEN",
	}
	colors["red"] = "RED"
	delete(colors, "red")
	for key, value := range colors {
		fmt.Printf("%s = %s\n", key, value)
	}
	fmt.Println(colors)
}
