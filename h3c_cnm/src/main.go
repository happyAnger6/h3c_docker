package main

import (
	"fmt"
	
	"github.com/docker/go-plugins-helpers/network"
)

type TestDriver struct{
	network.Driver
}

func (t *TestDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	fmt.Println("testdriver GetCapabilities")
	return &network.CapabilitiesResponse{Scope:network.LocalScope, ConnectivityScope:network.GlobalScope}, nil
}

func (t *TestDriver) CreateNetwork(r *network.CreateNetworkRequest) error {
	fmt.Printf("testdriver CreateNetwork !!!: %v", r)
	return nil
}

func (t *TestDriver) DeleteNetwork(r *network.DeleteNetworkRequest) error {
	fmt.Printf("testdriver DeleteNetwork!!!: %v", r)
	return nil
}

func (t *TestDriver) CreateEndpoint(r *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	fmt.Println("testdriver CreateEndpoint: %v", r)
	return &network.CreateEndpointResponse{}, nil
}

func (t *TestDriver) DeleteEndpoint(r *network.DeleteEndpointRequest) error {
	fmt.Println("testdriver DeleteEndpoint: %v", r)
	return nil
}

func (t *TestDriver) Join(r *network.JoinRequest) (*network.JoinResponse, error) {
	fmt.Println("testdriver Join: %v", r)
	return &network.JoinResponse{}, nil
}

func (t *TestDriver) Leave(r *network.LeaveRequest) error {
	fmt.Println("testdriver Leave: %v", r)
	return nil
}

func main(){

	d := &TestDriver{};
	h := network.NewHandler(d);
	h.ServeTCP("test_network", ":32234", "", nil)
	fmt.Printf("server started\n")
}

