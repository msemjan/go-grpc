package types

import (
	"context"
	"log"
)

type MyGRPCServer struct {
	UnimplementedNotificationServiceServer
}

var notifications = &Notifications{
	Notifications: []*Notification{
		&Notification{Title: "New message from 'Jerry'", Body: "Hi, how are you doing?", From: "Messenger", For: "User"},
		&Notification{Title: "Low batter", Body: "Plug in charger as soon as possible", From: "System", For: "User"},
		&Notification{Title: "Alarm", Body: "Wake up!", From: "Alarm", For: "User"},
	},
}

func (s *MyGRPCServer) SendNotification(ctx context.Context, notification *Notification) (*ServerResponse, error) {
	log.Printf("New notification: %s", notification.String())
	notifications.Notifications = append(notifications.Notifications, notification)
	return &ServerResponse{Status: 0}, nil
}

func (s *MyGRPCServer) GetNotifications(ctx context.Context, _ *Empty) (*Notifications, error) {
	log.Println("User is requesting notifications")
	return notifications, nil
}
