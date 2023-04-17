/*

Implement WordCount.
It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	f := strings.Fields(s)
	//fmt.Printf("Fields: %q \n", f)

	countMap := make(map[string]int)

	for _, v := range f {
		if count, ok := countMap[v]; ok {
			countMap[v] = count + 1
		} else {
			countMap[v] = 1
		}

	}
	//fmt.Println(countMap)
	return countMap
}

func main() {
	//r := WordCount("Hello This is me, is this me?")
	//fmt.Println(r)
	wc.Test(WordCount)
}
