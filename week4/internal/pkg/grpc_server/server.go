package grpc_server
import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server

	address string
}
// 新生成grpc server
func NewServer(address string) *Server {
	srv := grpc.NewServer()
	return &Server{Server: srv, address: address} // 开启grpcServer 并设置该grpc监听的服务端口
}

// 开启 server并等待
func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	log.Printf("grpc server start: %s", s.address)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
		log.Printf("grpc server gracefull stop")
	}()

	return s.Serve(l)// 将grpc服务与该服务进行绑定
}