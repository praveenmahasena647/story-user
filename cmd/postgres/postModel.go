package postgres

type Post struct {
	ID      uint   `db:id`
	Title   string `db:title`
	Post    string `db:post`
	User_id uint   `db:user_id`
}

func NewPost(t, p string, u_id uint) *Post {
	return &Post{
		Title:   t,
		Post:    p,
		User_id: u_id,
	}
}

func (p *Post) Insert() error {
	var _, err = db.Exec(`INSERT INTO posts(title, post, user_id) VALUES ($1,$2,$3);`, p.Title, p.Post, p.User_id)
	return err
}
