package main

import (
	"testing"
)

func TestResultLenVal(t *testing.T) {
	str1 := "abcdefghijklmnoprstuvyz"
	str2 := "rvyz"

	ret := Rollinghash(str2, str1, 3)

	t.Errorf("len of result slice is incorrect actual :%d , expected:%d", len(ret), 1)

}

func TestFastAfile(t *testing.T) {
	// content, err := ioutil.ReadFile("fasta")
	// Convert []byte to string and print to screen
	// str1 := string(content)
	str2 := "KORCANCCAACCA"
	str3 := "CTTCCCAKORCANCCAACCAACTENDTTCGATCTCTTGTAGATCTGTTCTCTAAACGAACTTTAAAATCT"

	// if err != nil {
	// 	log.Fatal(err)
	// }

	ret := Rollinghash(str2, str3, len(str2))
	t.Errorf("len of result slice is incorrect actual :%d , expected:%d", len(ret), 1)

}
