package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func taskFirst(input []string) {
    X := 101
    Y := 103
    moves := 100

    sums := [4]int{0, 0, 0, 0}

    for _, line := range input {
        if line == "" {
            continue
        }
        re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
        res := re.FindStringSubmatch(line)

        px, _ := strconv.Atoi(res[1])
        py, _ := strconv.Atoi(res[2])
        vx, _ := strconv.Atoi(res[3])
        vy, _ := strconv.Atoi(res[4])

        // for i := 0; i < 10; i++ {
        //     fmt.Println(((py+vy*i)%Y+Y)%Y, ((px+vx*i)%X+X)%X)
        // }
        x := ((px+vx*moves)%X+X)%X
        y := ((py+vy*moves)%Y+Y)%Y

        if x < X/2 && y < Y/2 {
            sums[0]++
        } else if x > X/2 && y < Y/2 {
            sums[1]++
        } else if x > X/2 && y > Y/2 {
            sums[2]++
        } else if x < X/2 && y > Y/2 {
            sums[3]++
        }
    }
    fmt.Println(sums[0]*sums[1]*sums[2]*sums[3])
}

func taskSecond(input []string) {

}


func main() {
    file := os.Args[1]

    var inputFile = file + ".in"

    dat, err := os.ReadFile(inputFile)

    if err != nil {
        panic(err)
    }
    taskFirst(strings.Split(string(dat), "\n"))
    // taskSecond(strings.Split(string(dat), "\n"))
}
