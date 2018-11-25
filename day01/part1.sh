#!/bin/bash
n=0
s=$(<input)
for (( i=0; i<${#s}; i++ )); do
    if [ ${s:i-1:1} = '(' ]; then
        ((n++))
    elif [ ${s:i-1:1} = ')' ]; then
        ((n--))
    fi
done
echo $n