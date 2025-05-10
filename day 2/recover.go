// package main

// import "fmt"
// func logging()  {
// 	fmt.Println("Selesai meamnggil function")
// 	message := recover()
// 	fmt.Println("Terjadi error",message)
// }

// func runApp(error bool)  {
// 	defer logging()
// 	fmt.Println("Run App")
// 	if error {
// 		panic("Uppss Error")
// 	}
// }

// func main()  {
// 	runApp(true)
// }