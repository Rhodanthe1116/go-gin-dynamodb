package db

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/vsouza/go-gin-boilerplate/config"
)

var db *dynamodb.DynamoDB

func Init() {
	c := config.GetConfig()
	db = dynamodb.New(session.New(&aws.Config{
		Region:      aws.String(c.GetString("db.region")),
		Credentials: credentials.NewEnvCredentials(),
		Endpoint:    aws.String(c.GetString("db.endpoint")),
		DisableSSL:  aws.Bool(c.GetBool("db.disable_ssl")),
	}))

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("BirthDay"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Gender"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("PhotoURL"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Name"),
				KeyType:       aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},

		TableName: aws.String("TableUsers"),
	}
	_, err := db.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

}

func GetDB() *dynamodb.DynamoDB {
	return db
}
