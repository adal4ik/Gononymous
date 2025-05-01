package dao

type Session struct {
	UsersId   string `json:"users_id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
}
