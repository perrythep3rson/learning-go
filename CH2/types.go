package main

import "fmt"
import "math"

func main() {
	
	var i int = 20
	var f float32 = float32(i)
	
	fmt.Println(i)
	fmt.Println(f)

	const value = 5

	var i2 int = value
	var f2 float32 = value

	fmt.Println(i2)
	fmt.Println(f2)

	var smallInt int32 = math.MaxInt32; 
	var bigInt uint64 = math.MaxUint64;
	var byte byte = math.MaxUint8;

	smallInt++
	byte++
	bigInt++

	fmt.Println(smallInt)
	fmt.Println(bigInt)
	fmt.Println(byte)

	

}
