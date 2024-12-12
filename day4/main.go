package main

import (
	"fmt"
	"os"
	"strings"
)

// func countXMAS(arr [][]byte, i int, j int) {
//     re := regexp.MustCompile(`X.*M.*A.*S`)
//
//     // for _, line := range
//     // fmt.Println(re.FindAll(line, -1))
// }


type position struct {
    row int
    col int
}


func checkX(input [][]byte, pos position) int {
    var sum int = 0
    // forward
    if (pos.col + 3 < len(input[pos.row])) {
        if (string(input[pos.row][pos.col:pos.col + 4]) == "XMAS") {
            sum += 1
        }
    }
    // backward
    if (pos.col - 3 >= 0) {
        if (string(input[pos.row][pos.col - 3:pos.col+1]) == "SAMX") {
            sum += 1
        }
    }
    // up
    if (pos.row - 3 >= 0) {
        if (input[pos.row][pos.col] == 'X' &&
            input[pos.row-1][pos.col] == 'M' &&
            input[pos.row-2][pos.col] == 'A' &&
            input[pos.row-3][pos.col] == 'S') {
            sum += 1
        }
    }

    // down
    if (pos.row + 3 < len(input)) {
        if (input[pos.row][pos.col] == 'X' &&
            input[pos.row+1][pos.col] == 'M' &&
            input[pos.row+2][pos.col] == 'A' &&
            input[pos.row+3][pos.col] == 'S') {
            sum += 1
        }
    }

    // up right
    if (pos.row - 3 >= 0 && pos.col + 3 < len(input[0])) {
        if (input[pos.row][pos.col] == 'X' &&
            input[pos.row-1][pos.col+1] == 'M' &&
            input[pos.row-2][pos.col+2] == 'A' &&
            input[pos.row-3][pos.col+3] == 'S') {
            sum += 1
        }

    }

    // down right
    if (pos.row + 3 < len(input) && pos.col + 3 < len(input[0])) {
        if (input[pos.row][pos.col] == 'X' &&
            input[pos.row+1][pos.col+1] == 'M' &&
            input[pos.row+2][pos.col+2] == 'A' &&
            input[pos.row+3][pos.col+3] == 'S') {
            sum += 1
        }

    }

    // down left
    if (pos.row + 3 < len(input) && pos.col - 3 >= 0) {
        if (input[pos.row][pos.col] == 'X' &&
            input[pos.row+1][pos.col-1] == 'M' &&
            input[pos.row+2][pos.col-2] == 'A' &&
            input[pos.row+3][pos.col-3] == 'S') {
            sum += 1
        }
    }

    // up left
    if (pos.row - 3 >= 0 && pos.col - 3 >= 0) {
        if (input[pos.row][pos.col] == 'X' &&
            input[pos.row-1][pos.col-1] == 'M' &&
            input[pos.row-2][pos.col-2] == 'A' &&
            input[pos.row-3][pos.col-3] == 'S') {
            sum += 1
        }
    }
    return sum
}


func taskFirst(input []string) {
    xs := make([]position, 0)
    height := len(input) - 1

    arr := make([][]byte, height)

    for i, line := range input {
        // fmt.Println(line, len(line))
        if len(line) == 0 {
            continue
        }
        arr[i] = make([]byte, len(line))

        for j, char := range line {
            arr[i][j] = byte(char)

            if (char == 'X') {
                xs = append(xs, position{i, j})
            }
        }
    }

    var sum int = 0
    for _, pos := range xs {
        sum += checkX(arr, pos)
    }

    fmt.Println(sum)
}


func checkA(input [][]byte, pos position) int {
    if (pos.row + 1 < len(input) && pos.row - 1 >= 0 && pos.col + 1 < len(input[0]) && pos.col - 1 >= 0) {
        if (input[pos.row+1][pos.col-1] + input[pos.row-1][pos.col+1] == 160 && input[pos.row-1][pos.col-1] + input[pos.row+1][pos.col+1] == 160) {
            return 1
        }
    }
    return 0
}

func taskSecond(input []string) {
    as := make([]position, 0)
    height := len(input) - 1

    arr := make([][]byte, height)

    for i, line := range input {
        // fmt.Println(line, len(line))
        if len(line) == 0 {
            continue
        }
        arr[i] = make([]byte, len(line))

        for j, char := range line {
            arr[i][j] = byte(char)

            if (char == 'A') {
                as = append(as, position{i, j})
            }
        }
    }

    var sum int = 0
    for _, pos := range as {
        sum += checkA(arr, pos)
    }

    fmt.Println(sum)
}


func main() {
    file := os.Args[1]

    var inputFile = file + ".in"

    dat, err := os.ReadFile(inputFile)

    if err != nil {
        panic(err)
    }
    // taskFirst(strings.Split(string(dat), "\n"))
    taskSecond(strings.Split(string(dat), "\n"))
}
