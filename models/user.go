package models

import (
    "fmt"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/Rhodanthe1116/go-gin-boilerplate/db"
)

type User struct {
    Phone string
    Password string
    UUID string
}

func (user User) Signup() (*User, error) {
	db := db.GetDB()
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		errors.New("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("TableUsers"),
        ConditionExpression: aws.String("attribute_not_exists(Phone)"),
	}
	if _, err := db.PutItem(params); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByPhone(phone string) (*User, error) {
	db := db.GetDB()
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Phone": {
				S: aws.String(phone),
			},
		},
		TableName:      aws.String("TableUsers"),
		ConsistentRead: aws.Bool(true),
	}
	resp, err := db.GetItem(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var user *User
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
		log.Println(err)
		return nil, err
	}
    if user.Phone != phone {
        return nil, fmt.Errorf("Phone not exists")
    }
	return user, nil
}
