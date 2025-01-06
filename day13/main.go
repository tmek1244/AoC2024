package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func getMoves(line string) (int, int) {
    re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
    res := re.FindStringSubmatch(line)

    x, _ := strconv.Atoi(res[1])
    y, _ := strconv.Atoi(res[2])

    return x, y
}

func getPrice(line string) (int, int) {
    re := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
    res := re.FindStringSubmatch(line)

    x, _ := strconv.Atoi(res[1])
    y, _ := strconv.Atoi(res[2])

    return x, y
}


func getBestSolution(x1, y1, x2, y2, x, y int) (int, int) {
    divider := x1*y2 - y1*x2

    if (x1*y-y1*x) % divider != 0 {
        return -1, -1
    }
    b := (x1*y-y1*x)/divider
    // fmt.Println(b)
    if (y - b*y2) % y1 != 0 {
        return -1, -1
    }
    a := (y - b*y2)/y1

    // fmt.Println(a)

    return a, b
}


func taskFirst(input []string) {
    var sum int = 0
    for i := 0; i < len(input)/4; i++ {
        firstLine := input[4*i]
        secondLine := input[4*i+1]
        thirdLine := input[4*i+2]

        x1, y1 := getMoves(firstLine)
        x2, y2 := getMoves(secondLine)
        x, y := getPrice(thirdLine)

        a, b := getBestSolution(x1, y1, x2, y2, x, y)

        if a == -1 {
            continue
        }
        sum += a*3 + b
    }
    fmt.Println(sum)
}

func taskSecond(input []string) {
    var sum int = 0
    for i := 0; i < len(input)/4; i++ {
        firstLine := input[4*i]
        secondLine := input[4*i+1]
        thirdLine := input[4*i+2]

        x1, y1 := getMoves(firstLine)
        x2, y2 := getMoves(secondLine)
        x, y := getPrice(thirdLine)

        a, b := getBestSolution(x1, y1, x2, y2, x+10000000000000, y+10000000000000)

        if a == -1 {
            continue
        }
        sum += a*3 + b
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
