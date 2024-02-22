
//💔❣💕💞💓💗
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID   primitive.ObjectId `json:"_id,omitempty" bson:"_id,omitempty`
	Task  string  	`json:"task,omitempty"`
	Status  bool  1json:"status,omitempty" 
}