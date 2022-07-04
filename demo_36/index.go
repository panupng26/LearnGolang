package main

import (
	"fmt"
	"time"
)

func main() {
	nowtime := time.Now().Unix()
	fmt.Printf("nowtime : %v", nowtime)
}