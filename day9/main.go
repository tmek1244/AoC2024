package main

import (
	"fmt"
	"os"
	"strings"
)


type file struct {
    id int
}


func changeString(line string, amount int) (string, []file) {
    res := make([]file, 0)

    counter := 0

    for counter < amount {
        pos := len(line) - 1
        if len(line) % 2 == 0 {
            pos = len(line) - 2
        }

        lastFile := int(line[pos] - '0')
        if lastFile - (amount - counter) > 0 {
            runes := []rune(line)
            runes[pos] = rune(lastFile - (amount - counter) + '0')
            line = string(runes)
            for i := 0; i < amount - counter; i++ {
                res = append(res, file{(pos+1)/2})
            }
            break
        } else {
            counter += lastFile
            line = line[:pos-1]

            for i := 0; i < lastFile; i++ {
                res = append(res, file{(pos+1)/2})
            }
        }
    }
    return line, res
}

func taskFirst(input []string) {
    line := input[0]
    newLine := line
    // newLine, movedFiles := changeString(line)
    sum := 0
    idCounter := 0
    movedCounter := 0
    position := 0
    size := len(newLine)
    i := 0
    for i < size {
        num := int(newLine[i] - '0')
        if i % 2 == 0 {
            tmp := idCounter * (position * num + (num)*(num-1)/2)
            sum += tmp
            idCounter++
        } else {
            changedLine, movedFiles := changeString(newLine, num)
            newLine = changedLine
            size = len(newLine)
            for j := 0; j < num; j++ {
                sum += (position + j) * movedFiles[j].id
                movedCounter++
            }
        }
        position += num
        i++
    }
    fmt.Println(sum)
}

func tmp(line string, amount int) (string, []file) {
    res := make([]file, 0)

    counter := 0

    for counter < amount {
        pos := len(line) - 1
        if len(line) % 2 == 0 {
            pos = len(line) - 2
        }

        lastFile := int(line[pos] - '0')
        if lastFile - (amount - counter) > 0 {
            runes := []rune(line)
            runes[pos] = rune(lastFile - (amount - counter) + '0')
            line = string(runes)
            for i := 0; i < amount - counter; i++ {
                res = append(res, file{(pos+1)/2})
            }
            break
        } else {
            counter += lastFile
            line = line[:pos-1]

            for i := 0; i < lastFile; i++ {
                res = append(res, file{(pos+1)/2})
            }
        }
    }
    return line, res
}


type Space struct {
    startPos int
    size int
}

type File struct {
    startPos int
    size int
    id int
}


func taskSecond(input []string) {
    line := input[0]
    newLine := line
    sum := 0
    idCounter := 0
    position := 0
    size := len(newLine)
    i := 0

    spaces := make([]Space, 0)
    files := make([]File, 0)
    for i < size {
        num := int(newLine[i] - '0')
        if i % 2 == 0 {
            files = append(files, File{position, num, idCounter})
            idCounter++
        } else {
            spaces = append(spaces, Space{position, num})
        }
        position += num
        i++
    }

    for i = len(files) - 1; i >= 0; i-- {
        for j := 0; j < len(spaces); j++ {
            if files[i].size <= spaces[j].size && files[i].startPos > spaces[j].startPos {
                files[i].startPos = spaces[j].startPos
                spaces[j].size -= files[i].size
                spaces[j].startPos += files[i].size
                break
            }
        }
    }

    for _, f := range files {
        sum += f.id * (f.startPos * f.size + f.size*(f.size-1)/2)
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
