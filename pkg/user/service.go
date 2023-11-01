package user

import (
	"fmt"
	geo "github.com/kellydunn/golang-geo"
	"math"
	"net/http"
)

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

func GetClosestDriverService(r *http.Request) (*User, error) {
	id, err := GetAuthenticatedUserId(r)
	if err != nil {
		return nil, err
	}

	u, err := GetByIdDB(id)
	if err != nil {
		return nil, err
	}

	d, err := GetAllDriversDB()
	if err != nil {
		return nil, err
	}

	if len(d) <= 0 {
		return nil, fmt.Errorf("no driver")
	}

	p := geo.NewPoint(u.Latitude, u.Longitude)

	closestKey, closestDist := math.MaxInt32, math.MaxFloat64
	for key, driver := range d {
		p2 := geo.NewPoint(driver.Latitude, driver.Longitude)
		dist := p.GreatCircleDistance(p2)
		if dist < closestDist {
			closestKey = key
			closestDist = dist
		}
	}

	return d[closestKey], nil
}
