package main

import "fmt"

func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])	//takes the ASCII int
	bucket := letter % buckets
	return bucket
}
