package models

import "time"

type Collect struct {
	CollectId string    `json:"collectId" db:"collect_id"`
	ParentId  string    `json:"parentId" db:"parent_id"`
	UserId    string    `json:"userId" db:"user_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
