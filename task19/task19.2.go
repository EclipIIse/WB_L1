package main

import "fmt"

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println(reverse("главрыба"))
}
