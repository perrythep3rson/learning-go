package main

import "fmt"

func main() {

	oneHundredSlice := []int{}
	for i:= 0; i <= 100; i++{
		oneHundredSlice = append(oneHundredSlice, i)
	}
	fmt.Println(oneHundredSlice)

	for _, val := range oneHundredSlice {

		if val % 2 == 0 && val % 3 == 0{
			fmt.Println(val)
			fmt.Println("Six!")
			continue
		}
		if val % 2 == 0{
			fmt.Println(val)
			fmt.Println("Two!")
			continue
			
		}
		if val % 3 == 0{
			fmt.Println(val)
			fmt.Println("Three!")
			continue
		}
		if val % 2 != 0 && val % 3 != 0{
			fmt.Println(val)
			fmt.Println("Never mind!")
			continue
		}
	}

	var total int = 0
	for i:=0; i <=100; i++{
		total = total + i
	}
	fmt.Printf("The total is: %d\n", total)
}



