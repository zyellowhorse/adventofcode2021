def main():
    print(Calculate_Horizontal_And_Depth_Positions())

def Calculate_Horizontal_And_Depth_Positions():
    horizontalPosition, depthPosition = 0,0
    instructions = ""

    with open('input.txt') as f:
        instructions = f.readlines()
    
    for instruction in instructions:
        movmentCommand = instruction.split(" ")
        
        # only works in python 3.10 and above
        #match movmentCommand[0]:
        #    case "foward":
        #        horizontalPosition += int(movmentCommand[1])

        #    case "up":
        #        depthPosition -= int(movmentCommand[1])

        #    case "down":
        #        depthPosition += int(movmentCommand[1])

        if movmentCommand[0] == "forward":
            horizontalPosition += int(movmentCommand[1])
        elif movmentCommand[0] == "up":
            depthPosition -= int(movmentCommand[1])
        elif movmentCommand[0] == "down":
            depthPosition += int(movmentCommand[1])

    return horizontalPosition * depthPosition

if __name__ == '__main__':
    main()
