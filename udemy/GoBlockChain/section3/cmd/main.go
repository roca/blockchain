package main

import (
	"fmt"
	"os"
	"strconv"

	"udemy.com/goblockchain/section3/utils"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println(utils.GetHost())
	fmt.Println(utils.FindNeighbors(utils.GetHost(), uint16(port), 0, 5, 5000, 5003))
}
