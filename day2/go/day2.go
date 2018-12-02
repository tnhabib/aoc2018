package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parse_box(box_id string) (is_two_box, is_three_box bool) {
	is_two_box = false
	is_three_box = false
	var char_map map[byte]bool = make(map[byte]bool)
	var box_length = len(box_id)
	for ii := 0; ii < box_length; ii++ {
		char_map[box_id[ii]] = true
	}

	for key := range char_map {

		var char_count int = strings.Count(box_id, string(key))
		//fmt.Printf("Scanned %s for %s, and found %d count\n", box_id, string(key), char_count)
		if char_count == 2 {
			is_two_box = true
		}

		if char_count == 3 {
			is_three_box = true
		}
	}

	return is_two_box, is_three_box

}

func compare_boxes(box1 string, box2 string) (num_char_diff int, diff_idx int) {

	for ii := 0; ii < len(box1); ii++ {
		if box1[ii] != box2[ii] {
			num_char_diff++
			diff_idx = ii
			if num_char_diff > 1 {
				break
			}
		}
	}
	return num_char_diff, diff_idx
}
func find_fabric_boxes(boxes []string) (box1 string, box2 string, diff_idx int) {
	var num_char_diff int

	for ii := 0; ii < len(boxes); ii++ {
		box1 = boxes[ii]
		for jj := 1; jj < len(boxes); jj++ {
			box2 = boxes[jj]
			num_char_diff, diff_idx = compare_boxes(box1, box2)
			if num_char_diff == 1 {
				break
			}
		}
		if num_char_diff == 1 {
			break
		}
	}

	return box1, box2, diff_idx
}
func main() {
	args := os.Args
	inputfile := args[1]
	file, err := os.Open(inputfile)
	fmt.Println("opened file")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//close the file at the end of this function
	defer file.Close()

	var box_list []string
	var two_boxes_found int = 0
	var three_boxes_found int = 0

	var is_two_box bool = false
	var is_three_box bool = false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input_box := scanner.Text()
		box_list = append(box_list, input_box)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		is_two_box, is_three_box = parse_box(input_box)

		if is_two_box {
			two_boxes_found++
		}

		if is_three_box {
			three_boxes_found++
		}

	}
	box1, box2, diff_idx := find_fabric_boxes(box_list)
	fmt.Printf("Number of two boxes : %d\n", two_boxes_found)
	fmt.Printf("Number of three boxes : %d\n", three_boxes_found)
	fmt.Printf("Checksum is %d\n", two_boxes_found*three_boxes_found)
	fmt.Printf("Box1 is %s\n", box1)
	fmt.Printf("Box2 is %s\n", box2)
	fmt.Printf("Diff position is %d\n", diff_idx)

	var common_chars string
	for ii := 0; ii < len(box1); ii++ {
		if ii != diff_idx {
			common_chars = common_chars + string(box1[ii])
		}

	}
	fmt.Printf("common box chars are %s\n", common_chars)
}
