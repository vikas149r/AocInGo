package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// var year2020 = 2020

	f, err := ioutil.ReadFile("C:/Users/sramkrishna/OneDrive - FactSet/Documents/CentralAuthWork/AoC/Day1Input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var line = scanner.Text())


	}
}
