package main

import "fmt"

// loop a map only can use 'for range'

func main() {

	// 一維
	myMap := make(map[string]string)

	myMap["n1"] = "N1"
	myMap["n2"] = "N2"

	for index, value := range myMap {
		fmt.Println(index, value)
	}

	// 二維
	studentMap := make(map[string]map[string]string)
	studentMap["s_001"] = make(map[string]string, 2)
	studentMap["s_001"]["name"] = "Tom"
	studentMap["s_001"]["sex"] = "Boy"
	studentMap["s_002"] = make(map[string]string, 2)
	studentMap["s_002"]["name"] = "Marry"
	studentMap["s_002"]["sex"] = "Girl"

	for index1, value := range studentMap {
		fmt.Println("index1 = ", index1)
		for index2, value2 := range value {
			fmt.Printf("\t index2 = %v, value2 = %v\n", index2, value2)
		}
	}
}
