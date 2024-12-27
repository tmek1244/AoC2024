package main

import (
	"fmt"
	"os"
	"strings"
)

type graph struct {
    rows, cols int
    area []byte
    regions []int
    regionId int
}

func (g *graph) get(row, col int) byte {
    if col >= g.cols || row >= g.rows || col < 0 || row < 0 {
        return 0
    }

    return g.area[row*g.cols + col]
}

func (g *graph) getRegion(row, col int) int {
    if col >= g.cols || row >= g.rows || col < 0 || row < 0 {
        return -1
    }

    return g.regions[row*g.cols + col]
}

func (g *graph) setRegion(row, col, regId int) {
    if col >= g.cols || row >= g.rows || col < 0 || row < 0 {
        return
    }
    g.regions[row*g.cols + col] = regId
}

type tuple struct {
    row, col int
}

func (g *graph) setRegions() {
    queue := make([]tuple, 0)

    for i := 0; i < g.rows; i++ {
        for j := 0; j < g.cols; j++ {
            regionId := g.getRegion(i, j)
            if (regionId != 0) {
                continue
            }
            regionId = g.regionId
            g.regionId++
            queue = append(queue, tuple{i, j})

            for len(queue) > 0 {
                ele := queue[0]
                queue = queue[1:]

                row := ele.row
                col := ele.col
                if (g.getRegion(row, col) != 0) {
                    continue
                }
                g.setRegion(row, col, regionId)
                if (g.get(row, col) == g.get(row+1, col)) {
                    queue = append(queue, tuple{row+1, col})
                }
                if (g.get(row, col) == g.get(row-1, col)) {
                    queue = append(queue, tuple{row-1, col})
                }
                if (g.get(row, col) == g.get(row, col+1)) {
                    queue = append(queue, tuple{row, col+1})
                }
                if (g.get(row, col) == g.get(row, col-1)) {
                    queue = append(queue, tuple{row, col-1})
                }
            }
        }
    }
}


type Price struct {
    size, permeter int
}

func (g *graph) countPrice() int {
    g.setRegions()

    priceByRegion := make(map[int]Price, 0)

    for i := 0; i < g.rows; i++ {
        for j := 0; j < g.cols; j++ {
            regId := g.getRegion(i, j)

            if _, ok := priceByRegion[regId]; !ok {
                priceByRegion[regId] = Price{0, 0}
            }
            if entry, ok := priceByRegion[regId]; ok {
                entry.size++

                if g.getRegion(i-1, j) != regId {
                    entry.permeter++
                }
                if g.getRegion(i+1, j) != regId {
                    entry.permeter++
                }
                if g.getRegion(i, j-1) != regId {
                    entry.permeter++
                }
                if g.getRegion(i, j+1) != regId {
                    entry.permeter++
                }
                priceByRegion[regId] = entry
            }
        }
    }

    sum := 0
    fmt.Println(priceByRegion)
    for _, val := range priceByRegion {
        sum += val.permeter * val.size
    }
    return sum
}


func taskFirst(input []string) {
    cols := len(input[0])
    rows := len(input) - 1

    g := graph{rows, cols, make([]byte, rows*cols), make([]int, rows*cols), 1}

    for i, line := range input {
        for j := range line {
            g.area[i*cols+j] = line[j]
        }
    }
    fmt.Println(g.countPrice())
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
