package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func partition(arr []int, low int, high int) (int, []int) {
    pivot := arr[high]
    i := low

    for j := low; j < high; j++ {
        if arr[j] < pivot {
            arr[i], arr[j] = arr[j], arr[i]
            i++
        }
    }

    arr[high], arr[i] = arr[i], arr[high]
    return i, arr
}

func quickSort(arr []int, low int, high int) []int {
    if low < high {
        p, arr := partition(arr, low, high)
        quickSort(arr, low, p-1)
        quickSort(arr, p+1, high)
    }
    return arr
}

func abs(a int, b int) int {
    var diff int
    if a > b {
        diff = a - b
    } else {
        diff = b - a
    }
    return diff
}

func taskFirst(firstColumn []int, secondColumn []int) {
    var sum int
    for i := range firstColumn {
        sum += abs(firstColumn[i], secondColumn[i])
    }

    fmt.Println(sum)
}

func taskSecond(firstColumn []int, secondColumn []int) {
    firstCouter := make(map[int]int)

    for _, value := range firstColumn {
        firstCouter[value] += 1
    }
    secondCounter := make(map[int]int)

    for _, value := range secondColumn {
        secondCounter[value] += 1
    }

    var sum int64 = 0
    for key, value := range firstCouter {
        sum += int64(key * value * secondCounter[key])
    }
    fmt.Println(sum)
}


func main() {
    var inputFile = "input.in"

    dat, err := os.ReadFile(inputFile)

    if err != nil {
        panic(err)
    }

    var firstColumn []int
    var secondColumn []int

    for _, line := range strings.Split(string(dat), "\n") {
        a := strings.Split(line, "   ")
        if len(a) == 2 {
            res, err := strconv.Atoi(a[0])
            if err != nil {
                panic(err)
            }
            firstColumn = append(firstColumn, res)

            res, err = strconv.Atoi(a[1])
            if err != nil {
                panic(err)
            }
            secondColumn = append(secondColumn, res)
        }
    }
    firstColumn = quickSort(firstColumn, 0, len(firstColumn)-1)
    secondColumn = quickSort(secondColumn, 0, len(secondColumn)-1)

    // taskFirst(firstColumn, secondColumn)
    taskSecond(firstColumn, secondColumn)
}
