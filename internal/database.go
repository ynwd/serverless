package internal

import (
	"context"
	"errors"
	"fmt"
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

func (d *database) addUser(ctx context.Context, data interface{}) (
	*firestore.DocumentRef, *firestore.WriteResult, error) {
	msg := ""
	user := data.(map[string]interface{})
	username := user["username"]
	email := user["email"]

	u := d.client.Collection("user").Where("username", "==", username).Documents(ctx)
	resU, err := u.GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(resU) > 0 {
		msg = fmt.Sprintf("Username '%v' telah terdaftar. Gunakan yang lain", username)
		return nil, nil, errors.New(msg)
	}

	e := d.client.Collection("user").Where("email", "==", email).Documents(ctx)
	resE, err := e.GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(resE) > 0 {
		msg = fmt.Sprintf("Email '%v' telah terdaftar. Gunakan yang lain", email)
		return nil, nil, errors.New(msg)
	}

	return d.addData(ctx, "user", data)
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
	var phone, email, address, file, video string
	pn := item["phone"]
	em := item["email"]
	ad := item["address"]
	fl := item["file"]
	vd := item["video"]
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
	if vd == nil {
		video = ""
	} else {
		video = vd.(string)
	}

	post.Phone = phone
	post.Email = email
	post.Address = address
	post.ID = item["id"].(string)
	post.Topic = item["topic"].(string)
	post.File = file
	post.Video = video
	post.Title = item["title"].(string)
	post.Content = item["content"].(string)
	post.Type = item["type"].(string)
	post.User = item["user"].(string)
	post.Created = item["created"].(time.Time)
	post.Price = item["price"].(int64)
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
	user.Username = item["username"].(string)

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

func (d *database) getPostByTopic(ctx context.Context, topic string) []interface{} {
	iter := d.client.Collection("post").
		Where("topic", "==", topic).
		Documents(ctx)
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
