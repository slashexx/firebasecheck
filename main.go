package main

import (
	"context"
	// "encoding/json"
	"fmt"
	"log"

	"gofr.dev/pkg/gofr"

	"github.com/slashexx/firebasecheck/firebaseSetup"
)

type sampleData struct {
	ID      string `json:'id'`
	Title   string `json:'title'`
	Content string `jsonL'content'`
}

func main() {
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	app.POST("/posts", addPostHandler)

	app.Run()
	fmt.Println("Server is running.")
}

func addPostHandler(ctx *gofr.Context) (interface{}, error) {
	firebasectx := context.Background()

	fmt.Println("Reached the posting data function")

	client, err := firebaseSetup.InitializeFirebase().Firestore(firebasectx)

	if err != nil {
		log.Fatalln("Error initialising Firestore")
	}

	defer client.Close()

	var post sampleData

	ctx.Bind(&post)

	if err != nil {
		log.Fatalln(`Error decoding the post body`)
	}

	_, err2 := client.Collection("posts").NewDoc().Create(firebasectx, post)

	if err2 != nil {
		log.Fatalln(`Error creating doc in firestore` + err2.Error())
	}

	return "Post added", err2

}
