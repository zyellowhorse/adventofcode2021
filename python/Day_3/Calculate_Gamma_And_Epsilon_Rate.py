def main():
    print(Calculate_Submarine_Power_Consumption())

def Calculate_Submarine_Power_Consumption():
    gammaRate, epsilonRate = "",""
    diagnosticReport = ""

    with open('input.txt') as f:
        diagnosticReport = f.readlines()


    mostCommonValueCount = 0
    for i in range(len(diagnosticReport[0])-1):
        for item in diagnosticReport:
            if item[i] == "1":
                mostCommonValueCount += 1

        if mostCommonValueCount > len(diagnosticReport) / 2:
            gammaRate = gammaRate + "1"
        else:
            gammaRate = gammaRate + "0"
        mostCommonValueCount = 0


    for i in gammaRate:
        if i == "0":
            epsilonRate = epsilonRate + "1"
        else:
            epsilonRate = epsilonRate + "0"
    return int(gammaRate, 2) * int(epsilonRate, 2)
    #print("gammaRate: " + gammaRate)
    #print("epsilonRage: " + epsilonRate)
    #print(int(gammaRate,2))
    #print(int(epsilonRate,2))
   


if __name__ == '__main__':
    main()
