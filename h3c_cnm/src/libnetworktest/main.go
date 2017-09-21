package main

import (
	_ "fmt"
	"log"

	"github.com/docker/docker/pkg/reexec"
	"github.com/docker/libnetwork"
	"github.com/docker/libnetwork/options"
	"github.com/docker/libnetwork/netlabel"
	"github.com/docker/libnetwork/config"
)

var controller libnetwork.NetworkController

func createController() error {
	networkType := "bridge"

	driverOptions := options.Generic{}
	genericOption := make(map[string]interface{})
	genericOption[netlabel.GenericData] = driverOptions
	controller, err := libnetwork.New(config.OptionDriverConfig(networkType, genericOption))

	return err;
}

func main(){
	if reexec.Init() {
		return
	}

	networkType := "bridge"

	err := createController()
	if err != nil {
		log.Fatalf("createController:%s", err)
	}

	network, err := controller.NewNetwork(networkType, "network1", "")
	if err != nil {
		log.Fatalf("controller.NewNetwork:%s", err)
	}

	ep, err := network.CreateEndpoint("veth_v9_010_a")
	if err != nil {
		log.Fatalf("CreateEndpoint:%s", err)
	}

	cnt, err := controller.NewSandBox("null_container",
				libnetwork.OptionHostname("test"),
				libnetwork.OptionDomainname("docker.io"),
				libnetwork.OptionExtraHost("web", "192.168.0.1"))
	if err != nil {
		log.Fatalf("NewSandBox:%s", err)
	}

	
	}

}	
