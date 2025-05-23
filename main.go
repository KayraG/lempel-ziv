package main

import (
	bit "app/bit"
	"fmt"
	"strconv"
	"strings"
)

type OrderedPair struct {
	ref      int
	addition int
}

func appendPair(list []OrderedPair, ref int, add int) []OrderedPair {
	return append(list, OrderedPair{ref, add})
}

func split(s string) []bit.Bit {
	splitted := strings.Split(s, "")
	index := 0
	ln := len(s)
	result := []bit.Bit{}

	for index < ln {
		x, er := strconv.Atoi(splitted[index])
		if er != nil {
			return []bit.Bit{0}
		}
		result = append(result, x)
		index++
	}

	return result
}

func LZ(s string) []OrderedPair {
	result := []OrderedPair{}
	syms := bit.NewEmptyList()
	var list bit.BitList = split(s)
	index := 0
	length := len(s)
	var temp bit.BitList

	for index < length {
		temp = bit.NewBitList()
		i := 0
		x := 0
		_x := 0
		for i+index < length {
			fmt.Printf("Index: %d %d %d \n", index, i, x)
			temp = append(temp, list[index+i])
			_x = x
			x = syms.IndexOf(temp) + 1
			if x == 0 {

				syms = bit.Add(syms, temp)
				result = appendPair(result, _x, list[index+i])
				break

			}
			i++
			if i+index == length {
				result = appendPair(result, _x, -1)
				break
			}
		}

		index += i + 1
	}

	return result
}

func decompress(res []OrderedPair) string {
	result := []string{}

	return strings.Join(result, "")
}

func main() {

	strs := []string{}
	res := LZ("0101001000111010101110100010")

	for _, value := range res {
		strs = append(strs, fmt.Sprintf("(%d, %d)", value.ref, value.addition))
	}

	fmt.Println(strings.Join(strs, ", "))

}
