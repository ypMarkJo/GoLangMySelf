package main

import (
	"fmt"
	"mark/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
	dictionary.Delete(word)
	_, err = dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	}
}
