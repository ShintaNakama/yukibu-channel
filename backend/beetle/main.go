package main

import (
	"github.com/ShintaNakama/yukibu-channel/backend/beetle/pb"
	"github.com/ShintaNakama/yukibu-channel/backend/beetle/rpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"local.rpcs.packages/football"
)

func main() {
	server := rpc.NewRPCServer(&rpc.RPCServerOption{
		Name: "beetle",
		Port: 51000,
	})

	server.Run(func(svr *grpc.Server) {
		//football.RegisterBeetleServer(svr, pb.ProvideBeetleServer(zap.L()))
		football.RegisterBeetleServer(svr, pb.ProvideBeetleServer(zap.L()))
	})
}
