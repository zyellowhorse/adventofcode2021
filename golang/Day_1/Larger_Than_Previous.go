package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func FindNumberOfIncrease() (int, error) {

    firstValue, secondValue, IncreaseAmount := 0,0,0
    itemArray := []int{}

    data, err := os.Open("./input.txt")
    if err != nil {
        return 0, err
    }
    defer data.Close()
    
    scanner := bufio.NewScanner(data)
    for scanner.Scan(){
        item, _ := strconv.Atoi(scanner.Text())
        itemArray = append(itemArray, item)
    }


    for i, item := range itemArray {
        if len(itemArray) - i <= 2 {
            fmt.Printf("length: %v\n", len(itemArray))
            fmt.Printf("Index: %v\n", i)
            return IncreaseAmount, nil
        }

        firstValue = item + itemArray[i+1] + itemArray[i+2] 
        fmt.Printf("%v > %v\n", firstValue, secondValue)
        if firstValue > secondValue && secondValue != 0 {
            IncreaseAmount = IncreaseAmount + 1
        }
        secondValue = firstValue
        fmt.Printf("IncreaseAmount: %v\n", IncreaseAmount)
    }

    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return IncreaseAmount, nil
}

func main() {
    amount, _:= FindNumberOfIncrease()
    fmt.Printf("Amount = %v\n", amount)
}
