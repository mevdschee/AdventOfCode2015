import sys

floor = 0
for input in sys.stdin:
    for c in input:
        if c == "(":
            floor += 1
        elif c == ")":
            floor -= 1

print(floor)