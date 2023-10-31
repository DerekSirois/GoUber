package user

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Latitude  float64
	Longitude float64
	Active    bool
	RoleId    int `db:"role_id"`
}

type Login struct {
	Email    string
	Password string
}
