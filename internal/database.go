package internal

import (
	"context"
	"errors"
	"log"
	"time"

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

func (d *database) addUser(ctx context.Context, data interface{}) {
	d.addData(ctx, "user", data)
}

func (d *database) getPostDetail(ctx context.Context, id string) (Post, error) {
	iter := d.client.Collection("post").Where("id", "==", id).Documents(ctx)
	var item map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return Post{}, err
		}
		item = doc.Data()
	}

	if item == nil {
		return Post{}, errors.New("not found")
	}

	post := Post{}
	var phone, email, address, file string
	pn := item["phone"]
	em := item["email"]
	ad := item["address"]
	fl := item["file"]
	if pn == nil {
		phone = ""
	} else {
		phone = pn.(string)
	}
	if em == nil {
		email = ""
	} else {
		email = em.(string)
	}
	if ad == nil {
		address = ""
	} else {
		address = ad.(string)
	}
	if fl == nil {
		file = ""
	} else {
		file = fl.(string)
	}

	post.Phone = phone
	post.Email = email
	post.Address = address
	post.ID = item["id"].(string)
	post.Topic = item["topic"].(string)
	post.File = file
	post.Title = item["title"].(string)
	post.Content = item["content"].(string)
	post.Type = item["type"].(string)
	post.User = item["user"].(string)
	post.Created = item["created"].(time.Time)
	return post, nil
}

func (d *database) getUserDetail(ctx context.Context, email, password string) (User, error) {
	iter := d.client.Collection("user").
		Where("email", "==", email).
		Where("password", "==", password).
		Documents(ctx)

	var item map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return User{}, err
		}
		item = doc.Data()
	}

	if item == nil {
		return User{}, errors.New("not found")
	}

	user := User{}
	user.ID = item["id"].(string)
	user.Email = item["email"].(string)
	user.Name = item["name"].(string)
	user.Password = item["password"].(string)

	return user, nil
}

func (d *database) getUserDetailByID(ctx context.Context, id string) (User, error) {
	iter := d.client.Collection("user").
		Where("id", "==", id).Documents(ctx)

	var item map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return User{}, err
		}
		item = doc.Data()
	}

	if item == nil {
		return User{}, errors.New("not found")
	}

	user := User{}
	user.ID = item["id"].(string)
	user.Email = item["email"].(string)
	user.Name = item["name"].(string)

	return user, nil
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
