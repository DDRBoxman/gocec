package main

import "github.com/DDRBoxman/gocec"
import "fmt"

func main() {
	var config cec.CECConfiguration
	config.DeviceName = "CECTest"

	if er := cec.Init(config); er != nil {
		fmt.Println(er)
		return	
	}

	adapter, er := cec.GetFirstAdapter()
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Println(adapter)

}
