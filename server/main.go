package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Tsuyopon-1067/grpc-client-streaming-test/scantext"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	scantext.UnimplementedSenderServer
}

// クライアントから呼び出される関数
func (s *server) SendText(ctx context.Context, in *scantext.ScanText) (*emptypb.Empty, error) {
	log.Printf("Received: %v", in.Content)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080") // 接続を待ち受ける
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()                       // gRPCサーバーを作成
	scantext.RegisterSenderServer(s, &server{}) // サーバーにサービスを登録
	fmt.Println("Server is running on :8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
