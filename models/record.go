package models

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/Rhodanthe1116/go-gin-boilerplate/db"
)

type Record struct {
    UserId string
    StoreId string
    Time int64
}

func (record Record) Record() (*Record, error) {
	db := db.GetDB()
	item, err := dynamodbattribute.MarshalMap(record)
	if err != nil {
		errors.New("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("TableRecordsUser"),
	}
	if _, err := db.PutItem(params); err != nil {
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}
	params = &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("TableRecordsStore"),
	}
	if _, err := db.PutItem(params); err != nil {
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}
	return &record, nil
}
