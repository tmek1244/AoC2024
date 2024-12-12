package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(a int, b int) int {
    if (a > b) {
        return a - b
    }
    return b - a
}


func checkIfSafe(values []int) bool {
    if len(values) < 2 {
        return true
    }
    if len(values) == 2 {
        if (abs(values[0], values[1]) >= 1 && abs(values[0], values[1]) <= 3) {
            return true
        }
        return false
    }
    var order string
    if (values[0] > values[1]) {
        order = "dec"
    } else {
        order = "inc"
    }

    prev := values[0]
    for _, v := range values[1:] {
        if (order == "inc" && prev >= v) {
            return false
        }
        if (order == "dec" && prev <= v) {
            return false
        }
        diff := abs(prev, v)
        if (diff < 1 || diff > 3) {
            return false
        }
        prev = v
    }

    return true
}

func taskFirst(input []string) {
    sum := 0
    for _, line := range input {
        if (len(line) == 0) {
            continue
        }

        splited := strings.Split(line, " ")
        a := make([]int, len(splited))

        for i, v := range splited {
            p, err := strconv.Atoi(v)
            if (err != nil) {
                continue
            }
            a[i] = p
        }
        if (checkIfSafe(a)) {
            sum += 1
        }

    }

    fmt.Println(sum)
}


func checkIfSafe2(values []int) bool {
    for i := range values {
        slice := make([]int, i)
        copy(slice, values[:i])
        if (checkIfSafe(append(slice, values[i+1:]...))) {
            return true
        }
    }
    return false;
}

func taskSecond(input []string) {
    sum := 0
    for _, line := range input {
        if (len(line) == 0) {
            continue
        }

        splited := strings.Split(line, " ")
        a := make([]int, len(splited))

        for i, v := range splited {
            p, err := strconv.Atoi(v)
            if (err != nil) {
                continue
            }
            a[i] = p
        }
        if (checkIfSafe2(a)) {
            sum += 1
        }

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
