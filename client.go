package main

import (
	ctx "context"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/msemjan/go-grpc/types"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while connecting to the server: %s\n", err)
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	_, err = client.SendNotification(context.Background(), &pb.Notification{Title: "Client notification", Body: "Hello from the client", From: "Client", For: "User"})
	if err != nil {
		log.Fatalf("Error when calling SendNotification: %s\n", err)
	}

	log.Println("Notification successfully sent\n")

	notifications, err := client.GetNotifications(ctx.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error when calling GetNotifications: %s\n", err)
	}

	log.Printf("Notifications received:")
	for _, notification := range notifications.Notifications {
		log.Printf("Title: '%s', Body: '%s', From: '%s', For: '%s'\n", notification.Title, notification.Body, notification.From, notification.For)
	}
}
