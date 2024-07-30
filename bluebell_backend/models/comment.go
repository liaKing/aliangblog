package models

import "time"

type Comment struct {
	PostID     string    `db:"post_id" json:"postId"`
	ParentID   string    `db:"parent_id" json:"parent_id"`
	CommentID  string    `db:"comment_id" json:"comment_id"`
	AuthorID   string    `db:"author_id" json:"author_id"`
	Content    string    `db:"content" json:"content"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}
