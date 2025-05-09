package dto

type Comment struct {
	CommentID      string
	PostID         string
	ParentID       string
	UserID         string
	UserName       string
	UserAvatarLink string
	Content        string
	ImageUrl       string
	CreatedAt      string
	Replies        []Comment
}
