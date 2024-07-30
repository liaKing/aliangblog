package models

import (
	"encoding/json"
	"errors"
	"time"
)

type Post struct {
	Id            int       `json:"id" db:"id"`
	PostID        string    `json:"post_id" db:"post_id"`
	Title         string    `json:"title" db:"title"`
	Content       string    `json:"content" db:"content"`
	AuthorId      string    `json:"author_id" db:"author_id"`
	CommunityID   string    `json:"community_id" db:"community_id"`
	Status        int32     `json:"status" db:"status"`
	CreateTime    time.Time `json:"-" db:"created_time"`
	UpdateTime    time.Time `json:"updateTime" db:"updated_time"`
	Published     bool      `json:"published" db:"published"`
	StarNumber    int64     `json:"starNumber" db:"star_number"`
	CommentNumber int64     `json:"commentNumber" db:"comment_number"`
	CollectNumber int64     `json:"collectNumber" db:"collect_number"`
}

func (p *Post) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Title       string `json:"title" db:"title"`
		Content     string `json:"content" db:"content"`
		CommunityID string `json:"community_id" db:"community_id"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Title) == 0 {
		err = errors.New("帖子标题不能为空")
	} else if len(required.Content) == 0 {
		err = errors.New("帖子内容不能为空")
	} else if required.CommunityID == "" {
		err = errors.New("未指定版块")
	} else {
		p.Title = required.Title
		p.Content = required.Content
		p.CommunityID = required.CommunityID
	}
	return
}

type ApiPostDetail struct {
	*Post
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
}
