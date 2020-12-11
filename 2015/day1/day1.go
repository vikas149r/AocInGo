package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	f, err := ioutil.ReadFile("C:/Users/sramkrishna/OneDrive - FactSet/Documents/CentralAuthWork/AoC/2015/Day1Input.txt")

	if err != nil {
		log.Fatal(err)
	}

  for
	fmt.Println(string(f))
}
