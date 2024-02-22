package main

import "fmt"
import "log"

func main() {

	ifaces, err := GetCurrIFaces()
	if (err != nil) {
		log.Fatal(err)
	}

	for _, iface := range ifaces {
		fmt.Println(iface.String())
	}
}
