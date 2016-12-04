#!/bin/bash
if [ $# -lt 1 ]; then
  echo "Argument DAY missing"
  exit 1
fi;
DAY=`printf %02d $1`
if [ ! -d day${DAY}a ]; then
  mkdir day${DAY}a
  mkdir day${DAY}b
  cd day${DAY}a
  code .
  firefox http://adventofcode.com/2015/day/$1
  firefox http://adventofcode.com/2015/day/$1/input
fi;
