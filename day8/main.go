package main

import (
	"fmt"
	"os"
	"strings"
)


type pos struct {
    row int
    col int
}


func checkIfInside(end, point pos) bool {
    return point.row >= 0 && point.col >= 0 && point.row < end.row && point.col < end.col
}


func calculateAntinodes(a, b pos, width, height int) []pos {
    diff_row := a.row - b.row
    diff_col := a.col - b.col

    res := make([]pos, 0)

    first_an := pos{a.row + diff_row, a.col + diff_col}
    second_an := pos{b.row - diff_row, b.col - diff_col}

    if checkIfInside(pos{height, width}, first_an) {
        res = append(res, first_an)
    }
    if checkIfInside(pos{height, width}, second_an) {
        res = append(res, second_an)
    }
    return res
}


func taskFirst(input []string) {
    antenas := make(map[byte][]pos, 0)
    width := len(input[0])
    height := 0

    for row, line := range input {
        if line == "" {
            continue
        }
        height++
        for col, char := range line {
            if char != '.' {
                if _, ok := antenas[byte(char)]; !ok {
                    antenas[byte(char)] = make([]pos, 0)
                }
                antenas[byte(char)] = append(antenas[byte(char)], pos{row, col})
            }
        }
    }

    unique := make(map[pos]bool, 0)
    for _, value := range antenas {
        for i := range value {
            for j := 0; j < i; j++ {
                res := calculateAntinodes(value[i], value[j], width, height)

                for _, r := range res {
                    unique[r] = true
                }
            }
        }
    }
    fmt.Println(height, width)
    // fmt.Println(antenas)
    fmt.Println(unique)
    fmt.Println(len(unique))
}


func calculateAntinodesGrid(a, b pos, width, height int) []pos {
    diff_row := a.row - b.row
    diff_col := a.col - b.col

    res := make([]pos, 2)

    res[0] = a
    res[1] = b

    first_an := pos{a.row + diff_row, a.col + diff_col}
    second_an := pos{b.row - diff_row, b.col - diff_col}

    for checkIfInside(pos{height, width}, first_an) {
        res = append(res, first_an)
        first_an = pos{first_an.row + diff_row, first_an.col + diff_col}
    }
    for checkIfInside(pos{height, width}, second_an) {
        res = append(res, second_an)
        second_an = pos{second_an.row - diff_row, second_an.col - diff_col}
    }
    return res
}


func taskSecond(input []string) {
    antenas := make(map[byte][]pos, 0)
    width := len(input[0])
    height := 0

    for row, line := range input {
        if line == "" {
            continue
        }
        height++
        for col, char := range line {
            if char != '.' {
                if _, ok := antenas[byte(char)]; !ok {
                    antenas[byte(char)] = make([]pos, 0)
                }
                antenas[byte(char)] = append(antenas[byte(char)], pos{row, col})
            }
        }
    }

    unique := make(map[pos]bool, 0)
    for _, value := range antenas {
        for i := range value {
            for j := 0; j < i; j++ {
                res := calculateAntinodesGrid(value[i], value[j], width, height)

                for _, r := range res {
                    unique[r] = true
                }
            }
        }
    }
    // fmt.Println(height, width)
    // fmt.Println(antenas)
    // fmt.Println(unique)
    fmt.Println(len(unique))
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
