package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type fabric_square struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func check_square(tmp fabric_square, table [][]string) bool {
	result := true
	for ii := tmp.x; ii < tmp.x+tmp.width; ii++ {
		for jj := tmp.y; jj < tmp.y+tmp.height; jj++ {
			if table[ii][jj] == "X" {
				result = false
			}
		}
	}
	return result

}
func count_multi_claims(table [][]string, table_size int) int {
	result := 0
	for ii := 0; ii < table_size; ii++ {
		for jj := 0; jj < table_size; jj++ {
			if table[ii][jj] == "X" {
				result++
			}
		}
	}
	return result
}

func claim_fabric(x int, y int, width int, height int, table [][]string, fabric_id int) bool {
	no_overlap := true
	for ii := x; ii < x+width; ii++ {
		for jj := y; jj < y+height; jj++ {
			if table[ii][jj] == "." {
				table[ii][jj] = strconv.Itoa(fabric_id)
			} else if table[ii][jj] != "." {
				table[ii][jj] = "X"
			}

		}

	}
	return no_overlap
}

func main() {
	args := os.Args
	inputfile := args[1]
	file, err := os.Open(inputfile)
	var squares []fabric_square

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//close the file at the end of this function
	defer file.Close()
	table_size := 1000
	fabric_table := make([][]string, table_size)
	for ii := range fabric_table {
		fabric_table[ii] = make([]string, table_size)
	}

	//init fabric table
	for ii := 0; ii < table_size; ii++ {
		for jj := 0; jj < table_size; jj++ {
			fabric_table[ii][jj] = "."
		}
	}

	scanner := bufio.NewScanner(file)
	var no_overlap bool
	fabric_id := 1
	for scanner.Scan() {
		input_text := scanner.Text()

		fabric_corner := strings.Trim(strings.Split(strings.Split(input_text, ":")[0], "@")[1], " ")

		fabric_corner_x, _ := strconv.Atoi(strings.Split(fabric_corner, ",")[0])
		fabric_corner_y, _ := strconv.Atoi(strings.Split(fabric_corner, ",")[1])

		fabric_dims := strings.Trim(strings.Split(input_text, ":")[1], " ")
		fabric_side1_width, _ := strconv.Atoi(strings.Split(fabric_dims, "x")[0])
		fabric_side2_height, _ := strconv.Atoi(strings.Split(fabric_dims, "x")[1])

		no_overlap = claim_fabric(fabric_corner_x, fabric_corner_y, fabric_side1_width, fabric_side2_height, fabric_table, fabric_id)
		tmp_square := fabric_square{id: fabric_id, x: fabric_corner_x, y: fabric_corner_y, width: fabric_side1_width, height: fabric_side2_height}
		squares = append(squares, tmp_square)

		fabric_id++

	}
	multi_claims := count_multi_claims(fabric_table, table_size)
	fmt.Printf("Number of multi claim squares : %d\n", multi_claims)
	for ii := range squares {
		no_overlap = check_square(squares[ii], fabric_table)
		if no_overlap {
			fmt.Printf("Square # %d has no overlap\n", squares[ii].id)
		}
	}
}
