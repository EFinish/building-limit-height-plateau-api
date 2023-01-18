package config

import (
	"os"
)

const (
	// Env vars
	BuildingdataDbURLEnv = "BUILDINGDATA_ACCESS_BUILDINGDATA_DB_URL"
	GrpcPortEnv          = "BUILDINGDATA_ACCESS_GRPC_PORT"
)

type Config struct {
	BuildingdataDbURL string
	GrpcPort          string
}

func NewConfig() Config {
	c := Config{}

	// Set defaults
	c.BuildingdataDbURL = "mongodb://mongo:27017"
	c.GrpcPort = "9002"

	return c
}

func GetConfig() *Config {
	c := NewConfig()

	BuildingdataDbURL, present := os.LookupEnv(BuildingdataDbURLEnv)
	if present {
		c.BuildingdataDbURL = BuildingdataDbURL
	}

	GrpcPort, present := os.LookupEnv(GrpcPortEnv)
	if present {
		c.GrpcPort = GrpcPort
	}

	return &c
}
