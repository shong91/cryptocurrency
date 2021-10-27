package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)


func main() {
	difficulty := 2
	target := strings.Repeat("0", difficulty)
	fmt.Println(target)
	nonce := 1
	for {
		hash := fmt.Sprintf("%x\n",sha256.Sum256([]byte("hello"+ fmt.Sprint(nonce))))
		fmt.Printf("Hash:%s\nTarget:%s\nNonce:%d\n", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			// kill the function 
			return
		} else {
			nonce++
		}
	}
}
