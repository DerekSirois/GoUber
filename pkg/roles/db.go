package roles

import "GoUber/pkg/storage"

func GetAll() ([]*Role, error) {
	r := make([]*Role, 0)
	err := storage.Db.Select(&r, "SELECT * FROM users")
	return r, err
}

func GetById(id int) (r *Role, err error) {
	err = storage.Db.Get(r, "SELECT * FROM users WHERE id = $1", id)
	return r, err
}
