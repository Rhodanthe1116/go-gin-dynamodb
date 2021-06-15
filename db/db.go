package db

import (
    "log"

	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
//	"github.com/Rhodanthe1116/go-gin-boilerplate/config"
)

var db *dynamodb.DynamoDB

func Init() {
    /*
	c := config.GetConfig()
	db = dynamodb.New(session.New(&aws.Config{
		Region:      aws.String(c.GetString("db.region")),
		Credentials: credentials.NewEnvCredentials(),
		Endpoint:    aws.String(c.GetString("db.endpoint")),
		DisableSSL:  aws.Bool(c.GetBool("db.disable_ssl")),
	}))
    */
    db = dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Phone"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Phone"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(30),
			WriteCapacityUnits: aws.Int64(10),
		},

		TableName: aws.String("TableUsers"),
	}
	if _, err := db.CreateTable(input); err != nil {
        log.Println("Got error calling CreateTable: %s", err)
	}

	input = &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Phone"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Phone"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(30),
			WriteCapacityUnits: aws.Int64(10),
		},

		TableName: aws.String("TableStores"),
	}
	if _, err := db.CreateTable(input);err != nil {
        log.Println("Got error calling CreateTable: %s", err)
	}
	input = &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UserId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Time"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UserId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Time"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(50),
		},

		TableName: aws.String("TableRecordsUser"),
	}

	if _, err := db.CreateTable(input);err != nil {
        log.Println("Got error calling CreateTable: %s", err)
	}

	input = &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("StoreId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Time"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("StoreId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Time"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(50),
		},

		TableName: aws.String("TableRecordsStore"),
	}

	if _, err := db.CreateTable(input);err != nil {
        log.Println("Got error calling CreateTable: %s", err)
    }
}

func Clear() {
    log.Println("Cleaning tables")
    tdb := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
    input := &dynamodb.DeleteTableInput{
        TableName: aws.String("TableUsers"),
    }
	if _, err := tdb.DeleteTable(input);err != nil {
        log.Println("Got error calling DeleteTable: %s", err)
    }
    input = &dynamodb.DeleteTableInput{
        TableName: aws.String("TableStores"),
    }
	if _, err := tdb.DeleteTable(input);err != nil {
        log.Println("Got error calling DeleteTable: %s", err)
    }
    input = &dynamodb.DeleteTableInput{
        TableName: aws.String("TableRecordsUser"),
    }
	if _, err := tdb.DeleteTable(input);err != nil {
        log.Println("Got error calling DeleteTable: %s", err)
    }
    input = &dynamodb.DeleteTableInput{
        TableName: aws.String("TableRecordsStore"),
    }
	if _, err := tdb.DeleteTable(input);err != nil {
        log.Println("Got error calling DeleteTable: %s", err)
    }
    log.Println("Done")
}

func GetDB() *dynamodb.DynamoDB {
	return db
}
