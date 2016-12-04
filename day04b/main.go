package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	i := 0
	for {
		data := string(input) + strconv.Itoa(i)
		sum := md5.Sum([]byte(data))
		str := hex.EncodeToString(sum[:])
		if str[0:6] == "000000" {
			fmt.Println(i)
			break
		}
		i++
	}
}
