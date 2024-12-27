package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func taskFirst(input []string) {
    values := make([]int, 0)

    for _, valStr := range strings.Split(input[0], " ") {
        v, _ := strconv.Atoi(valStr)
        values = append(values, v)
    }
    fmt.Println(values)

    for i := 0; i < 75; i++ {
        newValues := make([]int, 0)

        for _, val := range values {
            if val == 0 {
                newValues = append(newValues, 1)
            } else {
                valStr := strconv.Itoa(val)
                if len(valStr) % 2 == 0 {
                    leftPart := valStr[:len(valStr)/2]
                    rightPart := valStr[len(valStr)/2:]

                    leftVal, _ := strconv.Atoi(leftPart)
                    rightVal, _ := strconv.Atoi(rightPart)

                    newValues = append(newValues, leftVal)
                    newValues = append(newValues, rightVal)
                } else {
                    newValues = append(newValues, val * 2024)
                }
            }
        }
        values = newValues
    }
    fmt.Println(len(values))
}


type tuple struct {
    start int
    blinks int
}

var cache map[tuple]int


func getNumLength(num int) int {
    res := 1

    for num >= 10 {
        res++
        num /= 10
    }
    return res
}


func getStonesAfter(start, blinks int) int{
    if blinks == 0 {
        return 1
    }
    if _, ok := cache[tuple{start, blinks}]; ok {
        return cache[tuple{start, blinks}]
    }
    if start == 0 {
        return getStonesAfter(1, blinks-1)
    }
    if getNumLength(start) % 2 == 0 {
        valStr := strconv.Itoa(start)
        leftPart := valStr[:len(valStr)/2]
        rightPart := valStr[len(valStr)/2:]

        leftVal, _ := strconv.Atoi(leftPart)
        rightVal, _ := strconv.Atoi(rightPart)

        cache[tuple{start, blinks}] = (getStonesAfter(leftVal, blinks-1) +
                                       getStonesAfter(rightVal, blinks-1))
    } else {
        cache[tuple{start, blinks}] = getStonesAfter(start*2024, blinks-1)
    }
    return cache[tuple{start, blinks}]
}

func taskSecond(input []string) {
    cache = make(map[tuple]int, 0)
    sum := 0
    for _, valStr := range strings.Split(input[0], " ") {
        v, _ := strconv.Atoi(valStr)
        sum += getStonesAfter(v, 75)
    }
    fmt.Println(sum)
    // fmt.Println(cache)
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
