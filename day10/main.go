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

func BFS(arr [][]int, start pos, overlapCheck bool) int {
    sum := 0

    q := make([]pos, 0)
    visited := make(map[pos]bool, 0)
    q = append(q, start)

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if _, ok := visited[node]; ok && overlapCheck {
            continue
        }
        visited[node] = true

        if arr[node.row][node.col] == 9 {
            sum += 1
        } else {
            if (node.row > 0 &&
                    arr[node.row - 1][node.col] == arr[node.row][node.col] + 1) {
                q = append(q, pos{node.row-1, node.col})
            }
            if (node.col > 0 &&
                    arr[node.row][node.col - 1] == arr[node.row][node.col] + 1) {
                q = append(q, pos{node.row, node.col-1})
            }
            if (node.row < len(arr) - 1 &&
                    arr[node.row + 1][node.col] == arr[node.row][node.col] + 1) {
                q = append(q, pos{node.row+1, node.col})
            }
            if (node.col < len(arr[0]) - 1 &&
                    arr[node.row][node.col + 1] == arr[node.row][node.col] + 1) {
                q = append(q, pos{node.row, node.col+1})
            }
        }
    }
    return sum
}


func taskFirst(input []string) {
    height := len(input) - 1

    arr := make([][]int, height)
    zeros := make([]pos, 0)

    for i, line := range input {
        if line == "" {
            break
        }
        arr[i] = make([]int, len(line))
        for j := 0; j < len(line); j++ {
            arr[i][j] = int(line[j]-'0')
            if arr[i][j] == 0 {
                zeros = append(zeros, pos{i, j})
            }
        }
    }

    sum := 0
    for _, zero := range zeros {
        sum += BFS(arr, zero, true)
    }

    fmt.Println(sum)
}

func taskSecond(input []string) {
    height := len(input) - 1

    arr := make([][]int, height)
    zeros := make([]pos, 0)

    for i, line := range input {
        if line == "" {
            break
        }
        arr[i] = make([]int, len(line))
        for j := 0; j < len(line); j++ {
            arr[i][j] = int(line[j]-'0')
            if arr[i][j] == 0 {
                zeros = append(zeros, pos{i, j})
            }
        }
    }

    sum := 0
    for _, zero := range zeros {
        sum += BFS(arr, zero, false)
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
