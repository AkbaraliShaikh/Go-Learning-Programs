package main

import "fmt"

func main() {
	var mapList map[string]int
	mapList = map[string]int{"staff1": 1, "staff2": 2}

	mapCreated := make(map[string]int)

	mapCreated["student1"] = 1
	mapCreated["student2"] = 2
	mapCreated["student3"] = 3
	mapCreated["student4"] = 4
	mapCreated["student5"] = 5

	fmt.Println("# mapList #")
	for key, val := range mapList {
		fmt.Printf("key = %s val = %d \n", key, val)
	}

	fmt.Println("# examples #")
	fmt.Println(mapList["staff1"])
	fmt.Println(mapList["staff2"])

	fmt.Println("# mapCreated #")
	for key, val := range mapCreated {
		fmt.Printf("key = %s val = %d \n", key, val)
	}
}
