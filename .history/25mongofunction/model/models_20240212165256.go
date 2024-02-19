package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct{
  ID primitive.ObjectID   `json:"_id,omitempty" bson`
  Movie string
  Watched bool
}