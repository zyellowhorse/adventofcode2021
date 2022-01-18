def main():
    print(Calculate_Life_Support_Rating())

def Get_Oxygen_Generator_Rating(Array):
    tempArray = Array
    onesArray = []
    zerosArray = []
    for i in range(len(Array[0])-1):
        for item in tempArray:
            if len(tempArray) == 1:
                return(tempArray[0])

            if item[i] == "1":
                onesArray.append(item)
            else:
                zerosArray.append(item)

        if len(onesArray) >= len(zerosArray):
            tempArray = onesArray[:]
        else:
            tempArray = zerosArray[:]
        onesArray.clear()
        zerosArray.clear()
    
    return(tempArray[0])

def Get_C02_Scrubber_Rating(Array):
    tempArray = Array
    onesArray = []
    zerosArray = []
    for i in range(len(Array[0])-1):
        for item in tempArray:
            if len(tempArray) == 1:
                return(tempArray[0])

            if item[i] == "1":
                onesArray.append(item)
            else:
                zerosArray.append(item)

        if len(zerosArray) <= len(onesArray):
            tempArray = zerosArray[:]
        else:
            tempArray = onesArray[:]
        onesArray.clear()
        zerosArray.clear()
    
    return(tempArray[0])

def Calculate_Life_Support_Rating():
    diagnosticReport = []

    with open("input.txt") as f:
        for i in f.readlines():
            diagnosticReport.append(i)

    oxygenGeneratorRating = Get_Oxygen_Generator_Rating(diagnosticReport)
    C02ScrubberRatting = Get_C02_Scrubber_Rating(diagnosticReport)
    return int(oxygenGeneratorRating, 2) * int(C02ScrubberRatting, 2)

if __name__ == "__main__":
    main()
