package clientcommon

import (
	"flag"
	"fmt"
)

func CliPort(def uint) uint {
	var port uint
	flag.UintVar(&port, "port", def, "Port listen to.")
	flag.Parse()
	fmt.Println("Port: ", port)
	return port
}
