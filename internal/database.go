package internal

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func createClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

type database struct {
	client *firestore.Client
}

func (d *database) add(ctx context.Context, collection string, data interface{}) (
	*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, res, err := d.client.Collection(collection).Add(ctx, data)
	return ref, res, err
}

type Query struct {
	Collection string
	Field      string
	Op         string
	Value      interface{}
	OrderBy    string
}

func (d *database) get(ctx context.Context, q *Query) (*firestore.DocumentSnapshot, error) {
	it := d.client.CollectionGroup(q.Collection).
		Where(q.Field, q.Op, q.Value).
		OrderBy(q.OrderBy, firestore.Desc).
		Documents(context.Background())

	var item *firestore.DocumentSnapshot
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		item = doc
	}

	return item, nil
}

func (d *database) update(ctx context.Context, q *Query, updates []firestore.Update) (
	*firestore.WriteResult, error) {
	item, err := d.get(ctx, q)
	if err != nil {
		return nil, err
	}

	return d.client.Collection(q.Collection).Doc(item.Ref.ID).Update(ctx, updates)
}

func (d *database) delete(ctx context.Context, q *Query) (*firestore.WriteResult, error) {
	item, err := d.get(ctx, q)
	if err != nil {
		return nil, err
	}

	return d.client.Collection(q.Collection).Doc(item.Ref.ID).Delete(ctx)
}
