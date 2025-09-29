package config

import "os"

var DbURI = os.Getenv("DATABASE_URL")
