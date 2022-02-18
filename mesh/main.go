package main

import (
	"flag"
	"github.com/weibocom/motan-go"
	"github.com/weibocom/motan-go/config"
	vlog "github.com/weibocom/motan-go/log"
	"mesh-for-motan-php/utils"
)

var (
	configFile = flag.String("config", "server.yaml", "mesh config")
)

func main()  {
	done := make(chan struct{})
	defer close(done)

	conf, err := config.NewConfigFromFile(*configFile)
	if err != nil {
		vlog.Infoln(err.Error())
		return
	}
	flag.Parse()
	ext := utils.GetDefaultExt()
	agent := motan.NewAgent(ext)
	agent.StartMotanAgentFromConfig(conf)
	<-done
}
