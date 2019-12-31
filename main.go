package main

import (
	"fmt"
	"os"
	"crypto/md5"
	"encoding/hex"
	"sync"
	"time"
	"runtime"
	"bufio"
	"log"
	// "reflect"
)

func compareHash(str string, inputHash string){
	hash := md5.Sum([]byte(str))
	calcHash := string(hex.EncodeToString(hash[:])[:])
	fmt.Println("Checking md5(",str,")against ", inputHash)   
	if calcHash == inputHash{
		fmt.Println("Match found: ", str)
		os.Exit(0);
	}
}

func allStr(s []string, prefix, inputHash string, n,k int, wg *sync.WaitGroup){
	defer wg.Done();
	if k==0{
		compareHash(prefix, inputHash)
		return;
	}
	for i:=0;i<n;i++{
		newPrefix := prefix + s[i]
		wg.Add(1)
		allStr(s, newPrefix, inputHash, n, k-1, wg)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	chars:=[]string{
		"@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", 
		"1","2","3","4","5","6","7","8","9","0" }
	
	var k int
	n:=len(chars)
	var wg sync.WaitGroup
	
	//User input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input hash: ")
	inputHash, _ := reader.ReadString('\n')
	fmt.Print("Length of password: ")
	_, err := fmt.Scanf("%d", &k)
	if err!=nil {
		log.Print("Scan for k failed due to", err)
		return
	}

	start := time.Now()
	for ind := range chars{
		wg.Add(1)
		go allStr(append(chars[ind:],chars[:ind]...),"", inputHash, n, k, &wg)
	}
	wg.Wait()
	
	elapsed := time.Since(start)
	fmt.Printf("Cracked hash in %s", elapsed)
}
