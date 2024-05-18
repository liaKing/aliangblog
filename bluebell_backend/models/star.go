package models

import "time"

type Star struct {
	StartId   string    `json:"starId" db:"star_id"`
	ParentId  string    `json:"parentId" db:"parent_id"`
	UserId    string    `json:"userId" db:"user_id"`
	TargetId  string    `json:"targetId" db:"target_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
