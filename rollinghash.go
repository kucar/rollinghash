package main

import (
	"fmt"
	"math"
)

var myhash int = 0
var hashtooffset map[int64]int
var MODULAR int64 = 928223
var BASE int64 = 10

// find first common string of length L  in two strings A, B

func pow(base int64, power int, modular int64) int64 {
	return (int64)(math.Pow((float64)(base), (float64)(power))) % modular
}

func computeHash(mystring string, base int64, modular int64) int64 {
	var ret int64 = 0
	for i := 0; i < len(mystring); i++ {
		ret = ret*base + (int64)(mystring[i])
		ret %= modular
	}
	return ret
}

func computePower(base int64, power int, modular int64) int64 {
	var ret int64 = 1
	for i := 0; i < power; i++ {
		ret = (ret * base) % modular
	}
	return ret
}
func computeRoll(mystring string, rollength int, prevhash int64, base int64, power int64, firstchar byte, modular int64) int64 {
	var hsh int64 = 0

	lastchar := mystring[len(mystring)-1]

	hsh = base*(prevhash-((int64)(firstchar)*power)%modular) + (int64)(lastchar)

	for hsh < 0 {
		hsh = hsh + modular
	}
	// fmt.Println(base, "*(", prevhash, "-((", firstchar, ")*", power, "))+", lastchar)
	return hsh % modular
}

//iterate one char and calculate hash of each 3 char chunks, add to map
func createHashmap(mystring string, substrsize int64) {
	fmt.Println("createHash for ", mystring)
	if nil == hashtooffset {
		hashtooffset = make(map[int64]int)
	}

	for i := 0; i < len(mystring)-(int)(substrsize-1); i++ {
		hsh := computeHash(mystring[i:i+(int)(substrsize)], BASE, MODULAR)
		hashtooffset[hsh] = i
		fmt.Println(mystring[i:i+(int)(substrsize)], "---", hsh, "----", i)
	}

}

func Rollinghash(findthis string, inthis string, windowlen int) map[string][]int {

	str2indexes := make(map[string][]int)
	p := computePower(BASE, windowlen-1, MODULAR)
	fmt.Println("rollinghash: trying to find pattern ", findthis, " in ", inthis, " windowlen = ", windowlen)
	if len(inthis) <= windowlen {
		panic("wrong length of string ")
	}
	createHashmap(findthis, (int64)(windowlen))
	var hsh int64 = computeHash(inthis[0:windowlen], BASE, MODULAR)
	firstcharinwindow := inthis[0]
	fmt.Println("computed roll hash for ", inthis[0:windowlen], " = ", hsh, "firstchar=", firstcharinwindow)
	for i := 1; i < len(inthis)-windowlen+1; i++ {
		hsh = computeRoll(inthis[i:i+windowlen], windowlen, hsh, BASE, p, firstcharinwindow, MODULAR)
		fmt.Println("computed roll hash for ", inthis[i:i+windowlen], " = ", hsh, "firstchar=", firstcharinwindow)
		firstcharinwindow = inthis[i]
		if res, ok := hashtooffset[hsh]; ok {
			str2indexes[findthis[res:res+windowlen]] = append(str2indexes[findthis[res:res+windowlen]], i)
			fmt.Println("found  occurrance at ", i, " for pattern ", findthis[res:res+windowlen])
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
