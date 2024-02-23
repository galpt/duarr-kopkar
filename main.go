package main

import (
	"fmt"
)

func main() {

	startSrv := fmt.Sprintf("Server is running on %v", fmt.Sprintf(":%v", hostPortGin))

	fmt.Println()
	fmt.Println(startSrv)
	fmt.Println()

	server()

}
