package main

import (
	"atlantis/routerapi/api"
	"atlantis/routerapi/zk"
	"github.com/jigish/go-flags"
)

type RouterApi struct {
	*flags.Parser
	opts             *ApiOptions
	APIListenAddress string
	ZkServerAddress  string
	ZkNeedConfig     bool
}

type RouterApiOptions struct {
	APIListenAddress string `short:"A" long:"listen-address" description:"The address for the API to listen on"`
	ZkServerAddress  string `short:"Z" long:"zk-address" description:"The address of the ZK server"`
	InitZK           string `short:"I" long:"init-zk" description:"Initialize zk?"`
}

func NewApi() *Api {
	opts := &ApiOptions{}
	return &Api{Parser: flags.NewParser(opts, flags.Default), opts: opts}
}

func (api *Api) setConfig(options ApiOptions) {

	rapi.APIListenAddress = options.APIListenAddress
	rapi.ZkServerAddress = options.ZkServerAddress
	rapi.ZkNeedConfig = (options.InitZK != "")

}
func (api *Api) loadConfig() {
	api.Parse()

	if api.opts.APIListenAddress != "" {
		api.APIListenAddress = api.opts.APIListenAddress
	} else {
		api.APIListenAddress = "99999"
	}
	if api.opts.ZkServerAddress != "" {
		api.ZkServerAddress = api.opts.ZkServerAddress
	} else {
		api.ZkServerAddress = "28080"
	}
	if api.opts.InitZK == "y" {
		api.ZkNeedConfig = true
	}
}

func main() {

	api := NewApi()
	api.loadConfig()
	api.run()
}

func (api *Api) run() {

	api.Init(api.APIListenAddress)
	zk.Init(api.ZkServerAddress, api.ZkNeedConfig)

	api.Listen()

}
