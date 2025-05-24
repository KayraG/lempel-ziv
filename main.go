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
				result = appendPair(result, x, -1)
				break
			}
		}

		index += i + 1
	}

	return result
}

func decompress(res []OrderedPair) string {
	result := []string{}
	syms := bit.NewEmptyList()
	length := len(res)
	var temp OrderedPair
	var r bit.BitList

	for index := 0; index < length; index++ {
		temp = res[index]
		if temp.ref == 0 {
			r = bit.NewBitList()
			if temp.addition != -1 {
				r = append(r, temp.addition)
			}
		} else {
			r = syms.Get(temp.ref - 1)
			if r == nil {
				fmt.Printf("Error: You referenced the %dth item before its definition. \n", temp.ref)
			} else {
				if temp.addition != -1 {
					r = append(r, temp.addition)
				}
			}
		}

		syms = bit.Add(syms, r)
		result = append(result, r.ConvertToString())
	}

	return strings.Join(result, "")
}

func main() {

	strs := []string{}
	input := "01010011011001010001011111010101110100010010111101011010010110111110101000101010101011111100000"
	res := LZ(input)

	for _, value := range res {
		strs = append(strs, fmt.Sprintf("(%d, %d)", value.ref, value.addition))
	}

	fmt.Println(strings.Join(strs, ", "))

	output := decompress(res)

	fmt.Printf("\n\nDecompressed: %s \n Is it same? %t \n", output, output == input)

}
