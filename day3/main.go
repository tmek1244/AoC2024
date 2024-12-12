package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func taskFirst(input []string) {
    re := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
    // fmt.Printf("%q\n", re.FindAllStringSubmatch(input, -1))

    var sum int64 = 0

    for _, line := range input {
        for _, m := range re.FindAllStringSubmatch(line, -1) {
            first, _ := strconv.Atoi(m[1])
            second, _ := strconv.Atoi(m[2])

            sum += int64(first) * int64(second)
        }
    }
    fmt.Println(sum)
}

func taskSecond(input []string) {
    re := regexp.MustCompile(`mul\((\d*),(\d*)\)|do\(\)|don't\(\)`)
    // fmt.Printf("%q\n", re.FindAllStringSubmatch(input, -1))

    var sum int64 = 0
    enabled := true
    for _, line := range input {
        for _, m := range re.FindAllStringSubmatch(line, -1) {
            if (m[0] == "do()") {
                enabled = true
            } else if (m[0] == "don't()") {
                enabled = false
            } else if (enabled) {
                first, _ := strconv.Atoi(m[1])
                second, _ := strconv.Atoi(m[2])

                sum += int64(first) * int64(second)
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
