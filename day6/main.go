package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type guardSimulation interface {
    run(start pos)
    getResult() int
}

type pos struct {
    row int
    col int
    dir int
}

type lab struct {
    area [][]byte
    height int
    width int
}

func newLab(width int, height int, obstacles []pos) *lab {
    area := make([][]byte, height)

    for i := 0; i < height; i++ {
        area[i] = make([]byte, width)
    }

    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            area[i][j] = '.'
        }
    }

    for _, o := range obstacles {
        area[o.row][o.col] = '#'
    }

    l := lab{height: height, width: width, area: area}
    return &l
}


func (l lab) run(start pos) {
    fmt.Println(start)

    cur := pos{start.row, start.col, start.dir}

    for cur.row >= 0 && cur.row < l.height && cur.col >= 0 && cur.col < l.width {
        colChange := (2-cur.dir)%2
        rowChange := (cur.dir-1)%2
        if l.area[cur.row][cur.col] == '#' {
            cur.dir = (cur.dir+1)%4
            cur.col -= colChange
            cur.row -= rowChange
            continue
        }
        // fmt.Println(cur.row, cur.col, cur.dir, l.height, l.width)
        l.area[cur.row][cur.col] = 'X'
        cur.row += rowChange
        cur.col += colChange
    }
    // fmt.Println(l)
}

func (l lab) getResult() int {
    sum := 0
    for i := 0; i < l.height; i++ {
        for j := 0; j < l.width; j++ {
            if l.area[i][j] == 'X' {
                sum++
            }
        }
    }
    return sum
}

func taskFirst(input []string) {
    width := len(input[0])
    heigh := len(input) - 1

    re := regexp.MustCompile(`#`)
    var start pos
    obstacles := make([]pos, 0)
    for row, line := range input {
        if val := strings.Index(line, "^"); val != -1 {
            start = pos{row, val, 0}
        }
        for _, obsPos := range re.FindAllStringIndex(line, -1) {
            obstacles = append(obstacles, pos{row, obsPos[0], 0})
        }
    }

    lab := newLab(width, heigh, obstacles)

    lab.run(start)

    fmt.Println(lab.getResult())
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