package main

import (
	"fmt"
	// "time"
	"strings"
	// "reflect"
	"crypto/md5"
	"encoding/hex"
	"sync"
)

var res [][]string

func subset(numbers, data []string, n, r, index, i int){
	if index==r{
		temp:=make([]string, r, r)
		for j:=0;j<r;j++{
			temp=append(temp,data[j])
		}
		temp = temp[r:]
		res=append(res, temp)
		return;
	}
	if i>=n{
		return;
	}
	data[index] = numbers[i]
	subset(numbers, data, n, r, index+1, i+1)
	subset(numbers, data, n, r, index, i+1)
}

func compareHash(str string, inputHash string) bool{
	hash := md5.Sum([]byte(str))
	calcHash := string(hex.EncodeToString(hash[:])[:])
	fmt.Println("Checking ", str, "against ", inputHash)   
	if calcHash == inputHash{
		return true;
	}
	return false;
}

func perm(wg *sync.WaitGroup, a []string, size int) bool{

	defer wg.Done();
	if size==1 {
		if strings.Join(a,"") == "abc123"{
			fmt.Println("Hell yeah")
		}
		// fmt.Println(strings.Join(a,""))
		if compareHash(strings.Join(a,""), "9CDFB439C7876E703E307864C9167A15"){
			fmt.Println("--------------------MATCH FOUND--------------------------------")
			return true;
		}
		return false;
	}
	for i:=0; i<size; i++ {
		wg.Add(1)
		perm(wg, a, size-1)
		if size%2==1 {
			a[0],a[size-1]=a[size-1],a[0]
		} else{
			a[i],a[size-1]=a[size-1],a[i]
		}
	}
	return false
}

func main() {
	// numbers:=[]string{'a','b','d','e','f','g','h'}
	numbers:=[]string{"@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1","2","3","4","5","6","7","8","9","0","#"}
	// numbers := []int{1,2,3,4,5}
	r,n:=6,len(numbers)
	data:=make([]string,r)
	subset(numbers, data, n, r, 0, 0)
	// fmt.Println(res)
	// Worker group
	var wg sync.WaitGroup
	for i:=0;i<len(res);i++{
		wg.Add(1)
		go perm(&wg, res[i],len(res[i]))
	}	
	wg.Wait()
	fmt.Println("Main done")
}
