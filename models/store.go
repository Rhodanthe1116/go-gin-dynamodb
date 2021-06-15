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

type Store struct {
    Phone string
    Password string
    UUID string
    Name string
    Address string
}

func (store Store) Signup() (*Store, error) {
	db := db.GetDB()
	item, err := dynamodbattribute.MarshalMap(store)
	if err != nil {
		errors.New("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("TableStores"),
        ConditionExpression: aws.String("attribute_not_exists(Phone)"),
	}
	if _, err := db.PutItem(params); err != nil {
		return nil, err
	}
	return &store, nil
}

func GetStoreByPhone(phone string) (*Store, error) {
	db := db.GetDB()
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Phone": {
				S: aws.String(phone),
			},
		},
		TableName:      aws.String("TableStores"),
		ConsistentRead: aws.Bool(true),
	}
	resp, err := db.GetItem(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var store *Store
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &store); err != nil {
		log.Println(err)
		return nil, err
	}
    if store.Phone != phone {
        return nil, fmt.Errorf("Phone not exists")
    }
	return store, nil
}
