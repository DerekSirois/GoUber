package user

import "GoUber/pkg/storage"

func CreateDB(u User) error {
	_, err := storage.Db.Exec("INSERT INTO users(firstName, lastName, email, password, latitude, longitude, active, role_id) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8)", u.FirstName, u.LastName, u.Email, u.Password, u.Latitude, u.Longitude, u.Active, u.RoleId)
	return err
}

func GetByEmailDB(email string) (*User, error) {
	u := &User{}
	err := storage.Db.Get(u, "SELECT * FROM users WHERE email = $1", email)
	return u, err
}
