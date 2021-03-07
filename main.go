package main

import "fmt"

func main() {
	fmt.Println("hello world !")
	fmt.Println(sayWhat("恰如其十大覅为u人情味"))
}

func sayWhat(str string) string {
	return fmt.Sprintf("这句话‘%s’没听懂", str)
}
