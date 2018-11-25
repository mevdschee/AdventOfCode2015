import sys

n = 0
for input in sys.stdin:
    for c in input:
        if c == "(":
            n += 1
        elif c == ")":
            n -= 1

print(n)