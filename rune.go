package main

//type rune = int32

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    r := 'a'
    
    //Print Size: 4
    fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
    
    //Print Type: int32
    fmt.Printf("Type: %s\n", reflect.TypeOf(r))
    
    //Print Unicode CodePoint: U+0061
    fmt.Printf("Unicode CodePoint: %U\n", r)
    
    //Print Character: a
	fmt.Printf("Character: %c\n", r)
	
	// rune array to string
    runeArray := []rune{'a', 'b', '£'}
    s := string(runeArray)
	fmt.Println(s)
	
	// string to rune array
	ss := "ab£"
    rr := []rune(ss)
    fmt.Printf("%U\n", rr)
}
