package main

import "fmt"

const (

	str = "abcdefgihjklmnopqrstuvwxyz"
	wStr = "!@#$%^&*()_+-=/*1236547890{[\\~`';:.<>?,|/]}ABCDEFGIHJKLMNOPQRSTUVWXYZ"
    s1 = "fdgh;aeoi"
	s2 = "rrdddwwwvas"
	s3 = "45laskdjfv;l-"
	s4 = " "

)

var count = 0
//var i int
//var j int



func main() {
    
    var i, j int

	for i:= 0, i<len(s1), i++ {
        for j:=0, j < len(wStr), j++ {
			if s1[i] == wStr[j] {
				fmt.Printf("ERROR", s4)
			} 
		}
		
	}
	for i:=0, i<len(s3), i++ {
        for j:=0, j < len(wStr), j++ {
			if s3[i] == wStr[j] {
				fmt.Printf("ERROR", s4)
			} 
		}
		
	}
	for i:=0, i<len(str), i++ {
        for j:=0, j < len(s2), j++{
			if str[i] == s2[j] {
				count++
			} 
		}
		if count>0 {
			fmt.Print(str[i],count)
		}
		
	}
}