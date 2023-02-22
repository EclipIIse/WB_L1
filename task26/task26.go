/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
*/
package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) (string, bool) {
	str = strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return str, false
			}
		}
	}

	return str, true
}

func main() {
	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("aabcd"))
	fmt.Println(checkUnique("aAcbfg"))
}
