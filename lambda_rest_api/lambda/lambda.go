package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("in Shanghai with seconds:", time.Now().Format("2006-01-02"))
	fmt.Println("in Shanghai with seconds:", time.Now().Format("2006-01-02") > "2022-10-19")

	// lambda.Start(entrypoint.HandleRequest)
}
