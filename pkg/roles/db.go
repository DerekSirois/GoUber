package roles

import "GoUber/pkg/storage"

func GetAll() ([]*Role, error) {
	r := make([]*Role, 0)
	err := storage.Db.Select(&r, "SELECT * FROM roles")
	return r, err
}

func GetById(id int) (*Role, error) {
	r := &Role{}
	err := storage.Db.Get(r, "SELECT * FROM roles WHERE id = $1", id)
	return r, err
}
