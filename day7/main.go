package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func lenOfInt(val int) int {
    counter := 1
    for val >= 10 {
        val /= 10
        counter++
    }
    return counter
}


// func findOperatorsBrute(res int, curValue int, vals []int) bool {
//     if len(vals) == 0 {
//         return curValue == res
//     }
//     if curValue == 0 {
//         return findOperatorsBrute(res, vals[0], vals[1:])
//     }
//     // fmt.Println(res, curValue, vals)
//
//     addRes := curValue + vals[0]
//     mulRes := curValue * vals[0]
//     conRes, _ := strconv.Atoi(fmt.Sprintf("%d%d", curValue, vals[0]))
//
//     return findOperatorsBrute(res, addRes, vals[1:]) ||
//             findOperatorsBrute(res, mulRes, vals[1:]) ||
//             findOperatorsBrute(res, conRes, vals[1:])
// }
//

func findOperators(res int, vals []int, useCon bool) bool {
    if res < 0 {
        return false
    }
    if len(vals) == 0 {
        return res == 0
    }
    if len(vals) == 1 {
        return res == vals[0]
    }

    // fmt.Println(res, vals)
    if res % vals[len(vals)-1] == 0 && findOperators(res/vals[len(vals)-1], vals[:len(vals)-1], useCon) {
        return true
    }
    if findOperators(res - vals[len(vals) - 1], vals[:len(vals)-1], useCon) {
        return true
    }

    if useCon {
        x := int(math.Pow10(lenOfInt(vals[len(vals)-1])))
        if res % x == vals[len(vals)-1] {
            return findOperators(res / x, vals[:len(vals)-1], useCon)
        }
        return false
    } else {
        return false
    }
}


func taskFirst(input []string) {
    sum := 0
    for _, line := range input {
        if line == "" {
            continue
        }
        splited := strings.Split(line, ":")
        res, _ := strconv.Atoi(splited[0])

        vals := make([]int, 0)
        for _, val := range strings.Split(splited[1], " ") {
            if val == "" {
                continue
            }
            newVal, _ := strconv.Atoi(val)
            vals = append(vals, newVal)
        }
        if findOperators(res, vals, false) {
            sum += res
        }
    }
    fmt.Println(sum)

}

func taskSecond(input []string) {
    sum := 0
    for _, line := range input {
        if line == "" {
            continue
        }
        splited := strings.Split(line, ": ")
        res, _ := strconv.Atoi(splited[0])

        vals := make([]int, 0)
        for _, val := range strings.Split(splited[1], " ") {
            if val == "" {
                continue
            }
            newVal, _ := strconv.Atoi(val)
            vals = append(vals, newVal)
        }

        if findOperators(res, vals, true) {
            sum += res
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

