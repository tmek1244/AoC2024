package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func midIfWorking(rules map[int]map[int]bool, update []int) int {
    for i, v := range update {
        for j := 0; j < i; j++ {
            if _, ok := rules[v][update[j]]; ok {
                return 0
            }
        }
    }
    return update[len(update)/2]
}


func taskFirst(input []string) {
    firstPart := true
    rules := make(map[int]map[int]bool)

    sum := 0

    for _, line := range input {
        if (line == "") {
            firstPart = false
            continue
        }
        if (firstPart) {
            res := strings.Split(line, "|")

            first, _ := strconv.Atoi(res[0])
            second, _ := strconv.Atoi(res[1])

            if _, ok := rules[first]; ok {
                rules[first][second] = true
            } else {
                rules[first] = make(map[int]bool, 0)
                rules[first][second] = true
            }
        } else {
            updates := make([]int, 0)
            for _, val := range strings.Split(line, ",") {
                v, _ := strconv.Atoi(val)
                updates = append(updates, v)
            }
            sum += midIfWorking(rules, updates)
        }
    }
    fmt.Println(sum)
}

func fixAndMid(rules map[int]map[int]bool, update []int) int {
    resArr := make([]int, len(update))
    for i, v := range update {
        finalPos := 0
        for j := 0; j < i; j++ {
            if _, ok := rules[v][resArr[j]]; ok {
                break
            }
            finalPos = j + 1
        }
        if finalPos < i {
            copy(resArr[finalPos+1:], resArr[finalPos:])
        }
        resArr[finalPos] = v
    }

    return resArr[len(resArr)/2]
}


func taskSecond(input []string) {
    firstPart := true
    rules := make(map[int]map[int]bool)

    sum := 0

    for _, line := range input {
        if (line == "") {
            firstPart = false
            continue
        }
        if (firstPart) {
            res := strings.Split(line, "|")

            first, _ := strconv.Atoi(res[0])
            second, _ := strconv.Atoi(res[1])

            if _, ok := rules[first]; ok {
                rules[first][second] = true
            } else {
                rules[first] = make(map[int]bool, 0)
                rules[first][second] = true
            }
        } else {
            updates := make([]int, 0)
            for _, val := range strings.Split(line, ",") {
                v, _ := strconv.Atoi(val)
                updates = append(updates, v)
            }
            if midIfWorking(rules, updates) == 0 {
                sum += fixAndMid(rules, updates)
            }
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
