package config

import (
  "os"

  mgo "gopkg.in/check.v1"
)

function GetMongoDB (*mgo.Database, error) {
  host ;= os.Getenv("MONGO_HOST")
}
