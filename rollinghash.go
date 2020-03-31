package main

import (
	"fmt"
	"math"
)

var myhash int = 0
var hashtooffset map[uint64]int
var MODULAR uint64 = 1000000007
var BASE uint64 = 253

func generatePower(p uint64, stringsize int) uint64 {
	var power uint64 = 1
	var base uint64 = 257
	var overflow_control uint64 = 1000000007
	for i := 0; i < stringsize; i++ {
		power = (power * base) % overflow_control
	}
	return power
}

// find first common string of length L  in two strings A, B

func pow(base uint64, power int) uint64 {
	return (uint64)(math.Pow((float64)(base), (float64)(power)))
}

func computeHash(mystring string, base uint64, modular uint64) uint64 {
	var sum uint64 = 0
	for i := 0; i < len(mystring); i++ {
		sum += (uint64)(mystring[i]) * pow(base, (len(mystring)-1-i))
	}
	return sum % modular
}

func computeRoll(mystring string, rollength int, prevhash uint64, base uint64, firstchar byte, modular uint64) uint64 {
	var hsh uint64 = 0
	if prevhash == 0 {
		prevhash = computeHash(mystring, base, MODULAR)
	}
	lastchar := mystring[len(mystring)-1]
	hsh = base*(prevhash-((uint64)(firstchar)*pow(base, rollength-1))) + (uint64)(lastchar)
	return hsh % modular
}

//iterate one char and calculate hash of each 3 char chunks, add to map
func createHashmap(mystring string, substrsize uint64) {
	fmt.Println("createHash for ", mystring)
	if nil == hashtooffset {
		hashtooffset = make(map[uint64]int)
	}

	for i := 0; i < len(mystring)-(int)(substrsize-1); i++ {
		hsh := computeHash(mystring[i:i+(int)(substrsize)], BASE, MODULAR)
		hashtooffset[hsh] = i
		fmt.Println(mystring[i:i+(int)(substrsize)], "---", hsh, "----", i)
	}

}

func Rollinghash(findthis string, inthis string, k int) map[string][]int {

	str2indexes := make(map[string][]int)

	fmt.Println("rollinghash: trying to find pattern ", findthis, " in ", inthis, " k = ", k)
	if len(inthis) <= k {
		panic("wrong length of string ")
	}
	createHashmap(findthis, (uint64)(k))
	var hsh uint64 = computeHash(inthis[0:k], BASE, MODULAR)
	firstcharinwindow := inthis[0]
	fmt.Println("computed roll hash for ", inthis[0:k], " = ", hsh, "firstchar=", firstcharinwindow)
	for i := 1; i < len(inthis)-k+1; i++ {
		hsh = computeRoll(inthis[i:i+k], k, hsh, BASE, firstcharinwindow, MODULAR)
		fmt.Println("computed roll hash for ", inthis[i:i+k], " = ", hsh, "firstchar=", firstcharinwindow)
		firstcharinwindow = inthis[i]
		if res, ok := hashtooffset[hsh]; ok {
			str2indexes[findthis[res:res+k]] = append(str2indexes[findthis[res:res+k]], i)
			fmt.Println("found  occurrance at ", i, " for pattern ", findthis[res:res+k])
		}
	}
	return str2indexes
}

func main() {

	str1 := "abcdefghijklmrvyznoprstuvyz"
	str2 := "rvyz"

	for str, i := range Rollinghash(str2, str1, 3) {
		fmt.Println("str :", str, " indexes:", i)
	}
}
