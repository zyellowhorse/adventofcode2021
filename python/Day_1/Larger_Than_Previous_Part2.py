#!/usr/bin/env python3

def FindIncreaseAmount():
    sonarReadings = ""
    increaseAmount, firstWindow, secondWindow = 0,0,0

    with open("input.txt") as f:
        sonarReadings = f.readlines()

    sonarReadingsLenght = len(sonarReadings)

    for i in range(sonarReadingsLenght):
        if i >= sonarReadingsLenght-2:
            return increaseAmount
                        
        secondWindow = int(sonarReadings[i]) + int(sonarReadings[i+1]) + int(sonarReadings[i+2])
        #print("~~~~~~~~~~~")
        #print(firstWindow)
        #print(secondWindow)
        if firstWindow < secondWindow and firstWindow != 0:
            increaseAmount += 1

        firstWindow = secondWindow

    return increaseAmount


def main():
    result = FindIncreaseAmount()
    print(result)

if __name__ == '__main__':
    main()
