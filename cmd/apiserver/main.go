package main

import (
	"flag"
	"github.com/Traliaa/http-rest-api/internal/app/apiserver"
	api "github.com/Traliaa/http-rest-api/internal/app/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "congig-path", "configs/apiserver.toml", "path to config file")

}

func main() {
	//flag.Parse()
	//
	//config := apiserver.NewConfig()
	//_, err := toml.DecodeFile(configPath, config)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err := apiserver.Start(config); err != nil {
	//	log.Fatal(err)
	//}
	s := grpc.NewServer()
	srv := &apiserver.GRPCServer{}
	api.RegisterLoginServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		logrus.Fatal(err)
	}

}
