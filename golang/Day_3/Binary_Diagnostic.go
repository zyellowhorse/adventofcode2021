package main

import(
    "fmt"
    "bufio"
    "os"
    "strconv"
    "sort"
)

func main() {

    itemArray := []string{}
    tempArray := []int{}
    index := 0
    result := ""

    data, err := os.Open("./input.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer data.Close()
    
    scanner := bufio.NewScanner(data)
    for scanner.Scan(){
        itemArray = append(itemArray, scanner.Text())
    }

    for index < len(itemArray[0]) {
        for _, item := range itemArray {
            tempValue, _ := strconv.Atoi(item[index:index + 1])
            tempArray = append(tempArray, tempValue)
        }

        sort.Ints(tempArray)

        tempArrayLenght := len(tempArray)

        if tempArrayLenght % 2 == 1 {
            result = result + strconv.Itoa(tempArray[tempArrayLenght/2])
        } else if tempArrayLenght % 2 == 0 {
            firstValue := tempArray[tempArrayLenght/2]
            secondValue := tempArray[tempArrayLenght/2 + 1]
            if firstValue + secondValue == 2 {
                result = result + "1"
            } else {
                result = result + "0"
            }
        }

        tempArray = nil
        index++
    }
    
    reverseResult := ""

    for _, c := range result {
        if string(c) == "0" {
            reverseResult = reverseResult + "1"
        } else {
            reverseResult = reverseResult + "0"
        }
    }

    gamma_rate, _ := strconv.ParseInt(result, 2, 64)
    epsilon_rate, _ := strconv.ParseInt(reverseResult, 2, 64)
    fmt.Printf("Result: %v\n", gamma_rate * epsilon_rate)
    //fmt.Printf("Result: %v\n", result)
    //fmt.Printf("Result: %v\n", reverseResult)
}
