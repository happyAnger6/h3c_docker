package main

import (
	"fmt"
	
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
)

func cmdAdd(args *skel.CmdArgs) error {
	fmt.Println("cmdAdd")
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	fmt.Println("cmdDel")
	return nil
}

func main(){
	skel.PluginMain(cmdAdd, cmdDel, version.All)
}
