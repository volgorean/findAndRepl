package main

import "fmt"
import "io/ioutil"
// import "os"
import "strings"

func main() {
	toModText := openAndRead(getFilePath("Path to .txt you want to findAndRepl: "))
	modWords := strings.Split(openAndRead(getFilePath("Path to .txt with mod words seperated by (,)'s: ")), ",")

	fmt.Println(pullout(toModText, modWords))
	fmt.Scanf("\n")
}

// Given the base string and a slice of strings to replace pullout will atempt to replace all of the words with * characters
func pullout(p string, words []string) string {
	for _, w := range words {
		p = strings.Replace(p, w, strings.Repeat("*", len(w)), -1)
	}

	return p
}

// Given the the path to a file as a string openAndRead will attempt to return the contents of the file as a string
func openAndRead(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	stringForm := string(bytes)
	return stringForm
}

// Will attempt to get file path and use entered text to prompt user
func getFilePath(toSay string) string{
	var err error
	var pathToText string

	fmt.Print(toSay)
	_, err = fmt.Scanln(&pathToText)
	if err != nil {
	    fmt.Println("Error: ", err)
	}
	return pathToText
}