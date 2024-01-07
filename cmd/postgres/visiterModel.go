package postgres

type Model struct {
	Title string
	Post  string
	Name  string
}

func FetchAll() ([]*Model, error) {
	var models = []*Model{}
	var collection, collectionErr = db.Query(`SELECT title,post,name FROM users JOIN posts ON posts.user_id = users.id`)
	if collectionErr != nil {
		return nil, collectionErr
	}
	for collection.Next() {
		var model = &Model{}
		if err := collection.Scan(model.Name, model.Title, model.Post); err != nil {
			continue
		}
		models = append(models, model)
	}
	return models, nil
}
