package models

import (
	"time"
)

type Student struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`  
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`  
	Age       int       `json:"age" gorm:"not null"`  
	Grade     string    `json:"grade" gorm:"type:varchar(10);not null"`  
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`  
}
