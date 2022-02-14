package main

import (
	"fmt"

	"udemy.com/goblockchain/section3/utils"
)

func main() {
	//fmt.Println(utils.GetHost())
	fmt.Println(utils.FindNeighbors(utils.GetHost(), 5000, 0, 5, 5000, 5002))
}
