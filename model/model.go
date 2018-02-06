package model

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	UserKey string
	Posts   []Post
}

type Post struct {
	PostID    int
	Content   string
	StartDate time.Time
	EndDate   time.Time
}

func GetUser(db *sql.DB, userKey string) (User, error) {
	rows, err := db.Query("SELECT post_id FROM users_posts WHERE user_key = $1", userKey)
	if err != nil {
		return User{}, err
	}
	var posts []Post
	for rows.Next() {
		var postID int
		err := rows.Scan(&postID)
		if err != nil {
			return User{}, err
		}
		post, err := GetPost(db, postID)
		if err != nil {
			return User{}, err
		}
		posts = append(posts, post)
	}
	return User{
		UserKey: userKey,
		Posts:   posts,
	}
}

func GetPost(db *sql.DB, postID int) (Post, error) {
	var (
		content   string
		startDate time.Time
		endDate   time.Time
	)
	err := db.QueryRow("SELECT content, start_date, end_date, FROM posts WHERE post_id = $1", postID).Scan(&content, &startDate, &endDate)
	if err != nil {
		return Post{}, err
	}
	return Post{
		PostID:    postID,
		Content:   content,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
