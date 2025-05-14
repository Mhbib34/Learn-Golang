package main

import (
	"day-4/helper"
	"day-4/time"
	"fmt"
)



func main() {
	result := helper.SayHello("Habib")
	fmt.Println("Hello World")
	fmt.Println(result)

	waktu := time.GetTime()
	fmt.Println(waktu)
}