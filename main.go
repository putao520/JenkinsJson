package main

import (
	"ScriptHelper/pkg"
	"fmt"
)

func main() {
	fmt.Println("hello, world")

	name, version, err := ScriptHelper.Cargo()
	if err != nil {
		return
	}
	ScriptHelper.WriteScript(name, version)

	fmt.Printf("update file success")
}
