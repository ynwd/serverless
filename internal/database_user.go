package internal

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (d *database) createUser(ctx context.Context, data interface{}) (
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

	return d.add(ctx, "user", data)
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
		return User{}, errors.New("getUserDetailByID:not found")
	}

	user := User{}
	user.ID = item["id"].(string)
	user.Email = item["email"].(string)
	user.Name = item["name"].(string)
	user.Username = item["username"].(string)

	return user, nil
}

func (d *database) getUserIDWithSession(ctx context.Context, sessionID, userAgent string) (string, error) {
	if sessionID == "" {
		err := errors.New("getUserIDWithSession: sessionID empty")
		log.Println(err.Error())
		return "", err
	}
	iter := d.client.Collection("session").
		Where("sessionID", "==", sessionID).
		Where("userAgent", "==", userAgent).
		Documents(ctx)

	var item map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("getSession error: %v", err.Error())
			return "", err
		}
		item = doc.Data()
	}
	return item["userID"].(string), nil
}
