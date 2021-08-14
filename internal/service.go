package internal

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func createFirestoreClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

type client struct {
	firestore *firestore.Client
}

type Service interface {
	store() *firestore.Client
	add(ctx context.Context, collection string, data interface{}) (*firestore.DocumentRef, *firestore.WriteResult, error)
	get(ctx context.Context, q *Query) (*firestore.DocumentSnapshot, error)
	update(ctx context.Context, q *Query, updates []firestore.Update) (*firestore.WriteResult, error)
	delete(ctx context.Context, q *Query) (*firestore.WriteResult, error)
	createUser(ctx context.Context, data interface{}) (*firestore.DocumentRef, *firestore.WriteResult, error)
	getUserByActivationCode(ctx context.Context, code string) (*firestore.DocumentSnapshot, error)
	getUserDetail(ctx context.Context, email, password string) (*User, error)
	getUserDetailByID(ctx context.Context, id string) (User, error)
	getUserDetailByUsername(ctx context.Context, username string) (*User, error)
	getUserIDWithSession(ctx context.Context, sessionID string) (string, error)
	activateUserByCode(ctx context.Context, code string)
	createSession(ctx context.Context, userID string, userAgent string, ip string) string
	addPost(ctx context.Context, data interface{})
	getPostDetail(ctx context.Context, id string) (Post, error)
	getPost(ctx context.Context) []interface{}
	getPostByTopic(ctx context.Context, topic string) []interface{}
	getPostByUsername(ctx context.Context, username string) []interface{}
}

func createDatabase(ctx context.Context) Service {
	return &client{firestore: createFirestoreClient(ctx)}
}

type Query struct {
	Collection string
	Field      string
	Op         string
	Value      interface{}
	OrderBy    string
}

func (d *client) store() *firestore.Client {
	return d.firestore
}

func (d *client) add(ctx context.Context, collection string, data interface{}) (
	*firestore.DocumentRef, *firestore.WriteResult, error) {
	ref, res, err := d.firestore.Collection(collection).Add(ctx, data)
	return ref, res, err
}

func getDocumentSnapshot(it *firestore.DocumentIterator) (*firestore.DocumentSnapshot, error) {
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

func (d *client) get(ctx context.Context, q *Query) (*firestore.DocumentSnapshot, error) {
	it := d.firestore.CollectionGroup(q.Collection).
		Where(q.Field, q.Op, q.Value).
		OrderBy(q.OrderBy, firestore.Desc).
		Documents(context.Background())

	return getDocumentSnapshot(it)
}

func (d *client) update(ctx context.Context, q *Query, updates []firestore.Update) (
	*firestore.WriteResult, error) {
	item, err := d.get(ctx, q)
	if err != nil {
		return nil, err
	}

	return d.firestore.Collection(q.Collection).Doc(item.Ref.ID).Update(ctx, updates)
}

func (d *client) delete(ctx context.Context, q *Query) (*firestore.WriteResult, error) {
	item, err := d.get(ctx, q)
	if err != nil {
		return nil, err
	}

	return d.firestore.Collection(q.Collection).Doc(item.Ref.ID).Delete(ctx)
}
