package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Tsuyopon-1067/grpc-client-streaming-test/scantext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// grpc.WithInsecure() を指定するとTLSで暗号化を行わずに通信する
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close() // main関数終了時に接続を閉じる

	client := scantext.NewSenderClient(conn) // クライアントを作成

	for {
		var text string                           // 標準入力受取用
		fmt.Println("Enter text: (quit to exit)") // ただの指示
		fmt.Scan(&text)                           // 標準入力を受け取る
		if text == "quit" {                       // quit と入力されたらループを抜けて終了
			break
		}

		scanText := scantext.ScanText{ // 送信するメッセージを作成
			Content: text,
		}

		_, err := client.SendText(context.Background(), &scanText) // メッセージを送信
		if err != nil {                                            // エラーが発生したらログ出力して終了
			log.Fatalf("Error sending message: %v", err)
		}

		fmt.Printf("Sent message: %s\n", text) // ただのログ出力
	}

	fmt.Println("All messages sent successfully")
}
