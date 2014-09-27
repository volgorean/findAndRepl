package main

import "fmt"
// import "bufio"
// import "io/ioutil"
// import "os"
import "strings"

func main(){
	p := "Rooty tooty pants and frutty."
	
	words := []string{"pants", "frutty"}

	fmt.Println(pullout(p, words))
	fmt.Scanf("\n")
}

func pullout(p string, words []string) string{
	for _, w := range words {
		p = strings.Replace(p, w, strings.Repeat( "*", len(w) ), -1)
	}

	return p	
}