package main

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("firebaseConfig.json")

	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatalln("Firebase app not initialised")
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln("Firestore cline initialisation failed")
	}

	defer client.Close()

}
