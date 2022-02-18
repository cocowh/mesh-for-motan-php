package utils

import (
	"fmt"
	"github.com/weibocom/motan-go"
	"github.com/weibocom/motan-go/core"
	"github.com/weibocom/motan-go/endpoint"
	"github.com/weibocom/motan-go/filter"
	"github.com/weibocom/motan-go/ha"
	"github.com/weibocom/motan-go/lb"
	vlog "github.com/weibocom/motan-go/log"
	"github.com/weibocom/motan-go/provider"
	"github.com/weibocom/motan-go/registry"
	"github.com/weibocom/motan-go/serialize"
	"github.com/weibocom/motan-go/server"
	"google.golang.org/grpc"
	"mesh-for-motan-php/pb/testmsg"
	"mesh-for-motan-php/server/services"
	"net"
)

const (
	ServerTypeMotan = iota
	ServerTypeBreeze
	ServerTypeGrpcPB
)

func GetDefaultExt() core.ExtensionFactory  {
	ext := &core.DefaultExtensionFactory{}
	ext.Initialize()
	// all default extension
	filter.RegistDefaultFilters(ext)
	ha.RegistDefaultHa(ext)
	lb.RegistDefaultLb(ext)
	endpoint.RegistDefaultEndpoint(ext)
	provider.RegistDefaultProvider(ext)
	registry.RegistDefaultRegistry(ext)
	server.RegistDefaultServers(ext)
	server.RegistDefaultMessageHandlers(ext)
	serialize.RegistDefaultSerializations(ext)
	return ext
}

func RunMoranServer(serverType int)  {
	done := make(chan struct{})
	defer close(done)

	service,config := GetServiceAndConfig(serverType)

	ext := GetDefaultExt()
	msContext := motan.GetMotanServerContext(config)
	err := msContext.RegisterService(service, "")
	if err != nil {
		return
	}
	msContext.Start(ext)
	msContext.ServicesAvailable()
	<- done
}

func RunGrpcPBServer()  {
	lis,err := net.Listen("tcp", ":9528")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	testmsg.RegisterGreeterServer(s, &services.GrpcPBService{})
	vlog.Infoln(fmt.Sprintf("server listening at %v\n", lis.Addr()))
	if err := s.Serve(lis); err != nil  {
		vlog.Infoln(fmt.Sprintf("failed to serve: %v\n", err))
	}
}

func GetServiceAndConfig(serverType int) (services.Service, string) {
	switch serverType {
	case ServerTypeBreeze:
		return &services.BreezeService{}, "breeze.yaml"
	default:
		return &services.MotanService{}, "server.yaml"
	}
}