package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestResultLenVal(t *testing.T) {
	haystack := "abcdefghijklmnoprstuvyz"
	needle := "rvyz"

	ret := Rollinghash(needle, haystack, 3)
	//test the index of occurance
	if ret["vyz"][0] != 20 {
		t.Errorf("index of substring is incorrect actual  :%d , expected:%d", ret[needle], 7)
	}

}

func TestFindLongSubstring(t *testing.T) {

	needle := "KORCANCCAACCAACTEND"
	haystack := "CTTCCCAKORCANCCAACCAACTENDTTCGATCTCTTGTAGATCTGTTCTCTAAACGAACTTTAAAATCT"
	ret := Rollinghash(needle, haystack, len(needle))
	//test number of results
	if len(ret) != 1 {
		t.Errorf("len of result map is incorrect actual :%d , expected:%d", len(ret), 1)
	}
	//test number of occurrance
	if len(ret[needle]) != 1 {
		t.Errorf("index of substring is incorrect actual  :%d , expected:%d", ret[needle], 1)
	}
	//test the index of occurance
	if ret[needle][0] != 7 {
		t.Errorf("index of substring is incorrect actual  :%d , expected:%d", ret[needle], 7)
	}

}
func TestFindLineInFastAFile(t *testing.T) {
	content, err := ioutil.ReadFile("sars_cov2_fasta.txt")
	//Convert []byte to string and print to screen
	haystack := string(content)
	needle := "CCTTCAAAATTATGGTGATAGTGCAACATTACCTAAAGGCATAATGATGAATGTCGCAAAATATACTCAA"

	if err != nil {
		log.Fatal(err)
	}

	ret := Rollinghash(needle, haystack, len(needle))
	//test number of results
	if len(ret) != 1 {
		t.Errorf("len of result map is incorrect actual :%d , expected:%d", len(ret), 1)
	}
	//test number of occurrance
	if len(ret[needle]) != 1 {
		t.Errorf("index of substring is incorrect actual  :%d , expected:%d", ret[needle], 1)
	}
	//test the index of occurance
	if ret[needle][0] != 20946 {
		t.Errorf("index of substring is incorrect actual  :%d , expected:%d", ret[needle], 7)
	}

}
