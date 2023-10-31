package user

func CreateService(u User) error {
	hash, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	u.Active = true

	err = CreateDB(u)
	return err
}
