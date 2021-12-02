package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {

    firstValue, secondValue, IncreaseAmount := 0,0,0

    data, err := os.Open("./input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer data.Close()
    
    scanner := bufio.NewScanner(data)
    for scanner.Scan() {
        firstValue, _ = strconv.Atoi(scanner.Text())
        fmt.Printf("%v > %v\n", firstValue, secondValue)
        if firstValue > secondValue && firstValue != 0 && secondValue != 0 {
            //fmt.Printf("secondValue: %v\n", secondValue)
            IncreaseAmount = IncreaseAmount + 1
        }
        secondValue = firstValue
        fmt.Printf("IncreaseAmount: %v\n", IncreaseAmount)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    fmt.Printf("Increase Amount = %v\n", IncreaseAmount)
}
