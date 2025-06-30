package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMessageHistory(userID int, receiverID int) ([]bson.M, error) {

	messagesCollection := mongoDB.Collection("Messages")
	filter := bson.M{
		"$or": []bson.M{
			{"senderid": userID, "receiverid": receiverID},
			{"senderid": receiverID, "receiverid": userID},
		},
	}

	messages, err := messagesCollection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	defer messages.Close(context.TODO())

	var messagesList []bson.M
	for messages.Next(context.TODO()) {
		var message bson.M
		if err := messages.Decode(&message); err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		messagesList = append(messagesList, message)
	}

	return messagesList, nil

}