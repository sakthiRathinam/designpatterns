package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) string() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) addEntry(entry string) int {
	entryCount++
	entry = fmt.Sprintf("%d %s", entryCount, entry)
	j.entries = append(j.entries, entry)
	return entryCount
}

//breaks srp

func (j *Journal) saveToFile() {
	// ...
	panic("breaks srp")
}

type Peristence struct {
	lineSeperator string
}

func (per Peristence) saveToFile(j *Journal) {
	_ = ioutil.WriteFile("sample.txt", []byte(strings.Join(j.entries, per.lineSeperator)), 0644)

}

func main() {
	jour := Journal{}
	jour.addEntry("hello gopher")
	jour.addEntry("hello world")
	per := Peristence{"\n"}
	per.saveToFile(&jour)

}
