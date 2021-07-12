package internal

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func createClient(ctx context.Context) *firestore.Client {
	projectID := "fastro-319406"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

type database struct {
	client *firestore.Client
}

func (d *database) addData(ctx context.Context, collection string, data interface{}) (
	*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, res, err := d.client.Collection(collection).Add(ctx, data)
	return ref, res, err
}

func (d *database) addPost(ctx context.Context, data interface{}) {
	d.addData(ctx, "post", data)
}

func (d *database) getPost(ctx context.Context) []interface{} {
	iter := d.client.Collection("post").Documents(ctx)
	defer d.client.Close()
	var data []interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		data = append(data, doc.Data())
	}
	return data
}
