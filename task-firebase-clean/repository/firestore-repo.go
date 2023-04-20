package repository

import (
	"context"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	"github.com/sushanpth/learn-go/task-firebase-clean/entity"
	"google.golang.org/api/iterator"
)

type repo struct{}

// NewFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	collectionName = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("FIREBASE_PROJECT"))

	if err != nil {
		log.Fatalf("Failed to create firestore client")
		return nil, err
	}

	defer client.Close()
	_, _, err1 := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err1 != nil {
		log.Fatalf("Failed to add new post: %v", err1)
		return nil, err1
	}

	return post, nil

}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, os.Getenv("FIREBASE_PROJECT"))

	if err != nil {
		log.Fatalf("Failed to create firestore client")
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	list := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := list.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of post: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
