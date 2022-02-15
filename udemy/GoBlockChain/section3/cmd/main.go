package main

import (
	"fmt"
	"os"
	"strconv"

	"udemy.com/goblockchain/section3/utils"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	//fmt.Println(utils.GetHost())
	fmt.Println(utils.FindNeighbors(utils.GetHost(), port, 0, 5, 5000, 5002))
}
