package internal

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (d *database) addPost(ctx context.Context, data interface{}) {
	d.add(ctx, "post", data)
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

func (d *database) getPost(ctx context.Context) []interface{} {
	iter := d.client.Collection("post").
		OrderBy("created", firestore.Desc).
		Documents(ctx)
	defer d.client.Close()
	var data []interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("getPost: %v", err)
		}
		data = append(data, doc.Data())
	}
	return data
}

func (d *database) getPostByTopic(ctx context.Context, topic string) []interface{} {
	iter := d.client.Collection("post").
		OrderBy("created", firestore.Desc).
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
			log.Fatalf("getPostByTopic: %v", err)
		}
		data = append(data, doc.Data())
	}
	return data
}

func (d *database) getPostByUsername(ctx context.Context, username string) []interface{} {
	user, _ := d.getUserDetailByUsername(ctx, username)
	userID := "user"
	if user != nil {
		userID = user.ID
	}

	iter := d.client.Collection("post").
		Where("user", "==", userID).
		Documents(ctx)

	defer d.client.Close()
	var data []interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("getPostByUsername: %v", err)
		}
		data = append(data, doc.Data())
	}
	return data
}
