package main

import (
	// "Go_Projects/Back/back"
	"fmt"
)

func calc() {
	// const name, age = "Kim", 22
	var aa byte = 0xAA

	fmt.Println(aa)


	var x int = 99
	var y int = 1
	var total int = sum(x, y)
	fmt.Printf("Hello, World %d",total)

	name,age,adult := getCv("石上静香",31,true)
	fmt.Printf("\nCV: %s 歳: %d R18: %t", adult,age,name)

}

func sum(a int, b int) int {
	return a + b

}

func getCv(name string,age int,adult bool) (bool,int,string) {
	return  adult,age,name
}