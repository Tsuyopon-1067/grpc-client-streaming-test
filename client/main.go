package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Tsuyopon-1067/grpc-client-streaming-test/scantext"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := scantext.NewSenderClient(conn)

	for {
		var text string
		fmt.Println("Enter text: (quit to exit)")
		fmt.Scan(&text)
		if text == "quit" {
			break
		}

		scanText := scantext.ScanText{
			Content: text,
		}

		_, err := client.SendText(context.Background(), &scanText)
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}

		fmt.Printf("Sent message: %s\n", text)
	}

	fmt.Println("All messages sent successfully")
}
