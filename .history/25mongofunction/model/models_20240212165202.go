package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct{
  ID primitive.ObjectID
  Movie string
  Watched bool
}