package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    horizontalPostion, depth, aim := 0,0,0

    data, err := os.Open("./input.txt")
    if err != nil {
        //return 0, err
    }
    defer data.Close()
    
    scanner := bufio.NewScanner(data)
    for scanner.Scan(){
        instructions := strings.Fields(scanner.Text())
        fmt.Println(instructions)
        switch instructions[0] {
        case "forward":
            temp, _ := strconv.Atoi(instructions[1])
            horizontalPostion += temp 
            depth += aim * temp
        case "down":
            temp, _ := strconv.Atoi(instructions[1])
            aim += temp 
        case "up":
            temp, _ := strconv.Atoi(instructions[1])
            aim -= temp 
        }
        fmt.Printf("HP: %v, D: %d\n", horizontalPostion, depth)
    }

    fmt.Printf("Result: %v\n", horizontalPostion * depth)
}


