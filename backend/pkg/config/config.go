package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Port uint16
	Mail MailSettings
}

type MailSettings struct {
	FromMail  string
	FromEmail string
	ApiKey    string
	SecretKey string
}

type Provider interface {
	Get(ctx context.Context) (Config, error)
}

type MongodbProvider struct {
	c *mongo.Collection
}

const connectMongodbURLEnvVar = "CONFIG_MONGODB_CONNECT_URL"
const databaseName = "battleshipConfig"
const collectionName = "config"

func NewMongodbProvider() (*MongodbProvider, error) {
	connectURL := os.Getenv(connectMongodbURLEnvVar)
	opt := options.ClientOptions{}
	opt.ApplyURI(connectURL)
	if err := opt.Validate(); err != nil {
		return nil, fmt.Errorf(`failed to validate client options: [connect_url: %s, err: %w]`, connectURL, err)
	}
	cl, err := mongo.NewClient(&opt)
	if err != nil {
		return nil, fmt.Errorf(`failed to create client: [connect_url: %s, err: %w]`, connectURL, err)
	}
	return &MongodbProvider{c: cl.Database(databaseName).Collection(collectionName)}, nil
}

func (p *MongodbProvider) Get(ctx context.Context) (Config, error) {
	return Config{}, nil
}
