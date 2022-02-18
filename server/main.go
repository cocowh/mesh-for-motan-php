package main

import (
	"flag"
	"mesh-for-motan-php/utils"
)


var (
	serverType = flag.Int("st", utils.ServerTypeMotan, "server type")
)

func main()  {
	flag.Parse()
	switch *serverType {
	case utils.ServerTypeGrpcPB:
		utils.RunGrpcPBServer()
	default:
		utils.RunMoranServer(*serverType)
	}
}
