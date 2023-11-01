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

func GetByIdDB(id int) (*User, error) {
	u := &User{}
	err := storage.Db.Get(u, "SELECT * FROM users WHERE id = $1", id)
	return u, err
}

func GetAllDriversDB() ([]*User, error) {
	u := make([]*User, 0)
	err := storage.Db.Select(&u, "SELECT u.* FROM users u JOIN roles r ON u.role_id = r.id  WHERE r.name = 'Driver' and u.active = true")
	return u, err
}
