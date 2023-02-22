/* К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.
*/

package main

import "fmt"

func someFunc() {
	v := createHugeString(1 << 10)
	justString := v[:100]
	//fmt.Println(justString)
	fmt.Printf("SomeString lenght is %d\n", len(v))
	fmt.Printf("justString lenght is %d\n", len(justString))
	//fmt.Println(v)
}

func createHugeString(size int) string {
	var someString string
	for i := 0; i < size; i++ {
		someString += "k"
	}
	return someString
}

func main() {
	someFunc()
}
