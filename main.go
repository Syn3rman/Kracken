package main

import (
	"fmt"
	"os"
	"crypto/md5"
	"encoding/hex"
	// "time"
	// "reflect"
)

// MAXCOUNT -> Max number of goroutines
const MAXCOUNT = 1000000
var sem = make(chan int, MAXCOUNT)

func compareHash(str string, inputHash string){
	hash := md5.Sum([]byte(str))
	calcHash := string(hex.EncodeToString(hash[:])[:])
	fmt.Println("Checking md5(", str, ")against ", inputHash)   
	if calcHash == inputHash{
		fmt.Println("--------------------MATCH FOUND--------------------------------", str)
		os.Exit(0);
		// return true;
	}
	<-sem
}

func allStr(s []string, prefix string, n,k int){
	if k==0{
		sem <- 1
		go compareHash(prefix, "26e9262dceb7a2c4e1dfc417b5234370")
		return;
	}
	for i:=0;i<n;i++{
		newPrefix := prefix + s[i]
		allStr(s, newPrefix, n, k-1)
	}
}

func main() {
	// chars:=[]string{"a","b","c","d","e","1","2","3","4"}
	chars:=[]string{
		/*"@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		*/"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", 
		"1","2","3","4","5","6","7","8","9","0" }
	k,n:=7,len(chars)
	allStr(chars,"" , n, k)
}
