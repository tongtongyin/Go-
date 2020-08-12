package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes吕图大西瓜!" // UTF-8
	for _, b := range []byte(s) {
		fmt.Printf("%X ",b)
	}
	fmt.Println()
	for i, ch := range s{ // ch is a rune
		fmt.Printf("(%d %X) ",i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	// ch字符，size字符大小

	for len(bytes) > 0{
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}
	fmt.Println()
	for i, ch := range []rune(s){
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()


}
