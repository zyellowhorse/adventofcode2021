#!/usr/bin/env python3

sonarReadings = ""
increaseAmount = 0

with open("input.txt") as f:
    sonarReadings = f.readlines()

for i in range(1, len(sonarReadings)):
    if int(sonarReadings[i]) > int(sonarReadings[i-1]):
        increaseAmount += 1

print(increaseAmount)

