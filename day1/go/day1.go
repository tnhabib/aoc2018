package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	inputfile := args[1]
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//close the file at the end of this function
	defer file.Close()
	var input_num int
	var nums []int

	var first_final_frequency int = 0
	var first_fina_frequency_computed = false
	var first_repeat_freq int = 0
	var current_frequency int = 0
	var first_repeat_freq_found bool = false

	for {

		_, err := fmt.Fscanf(file, "%d\n", &input_num) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, input_num)
	}

	var frequency_map map[int]int
	frequency_map = make(map[int]int)
	for !first_repeat_freq_found {
		for _, v := range nums {
			current_frequency += v
			frequency_map[current_frequency]++

			if frequency_map[current_frequency] == 2 && !first_repeat_freq_found {
				first_repeat_freq = current_frequency
				first_repeat_freq_found = true
				fmt.Printf("Found repeat frequency %d \n", current_frequency)
			}
		}
		if !first_fina_frequency_computed {
			first_final_frequency = current_frequency
			first_fina_frequency_computed = true
		}

	}

	fmt.Printf("first Final frequency is %d\n", first_final_frequency)
	fmt.Printf("First repeat frequency is %d\n", first_repeat_freq)
}
