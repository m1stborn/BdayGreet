package model

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	Birth     time.Time `json:"birth"`
	Email     string    `json:"email"`
}

var (
	DBRecords = []User{
		{
			FirstName: "Robert",
			LastName:  "Yen",
			Gender:    "Male",
			Birth:     time.Date(1975, 8, 8, 0, 0, 0, 0, time.UTC),
			Email:     "robert.yen@linecorp.com",
		},
		{
			FirstName: "Cid",
			LastName:  "Change",
			Gender:    "Male",
			Birth:     time.Date(1990, 10, 10, 0, 0, 0, 0, time.UTC),
			Email:     "cid.change@linecorp.com",
		},
		{
			FirstName: "Miki",
			LastName:  "Lai",
			Gender:    "Female",
			Birth:     time.Date(1993, 4, 5, 0, 0, 0, 0, time.UTC),
			Email:     "miki.lai@linecorp.com",
		},
		{
			FirstName: "Sherry",
			LastName:  "Chen",
			Gender:    "Female",
			Birth:     time.Date(1993, 8, 8, 0, 0, 0, 0, time.UTC),
			Email:     "sherry.chen@linecorp.com",
		},
		{
			FirstName: "Peter",
			LastName:  "Wang",
			Gender:    "Male",
			Birth:     time.Date(1950, 12, 22, 0, 0, 0, 0, time.UTC),
			Email:     "peter.wang@linecorp.com",
		},
	}
)

func (d *Database) CreateUser(user *User) {
	if err := d.db.Create(user).Error; err != nil {
		//logger.WithFields(log.Fields{
		//	"pkg":  "model",
		//	"func": "CreateUser",
		//}).Error(err)
		fmt.Println("CreateUser Error", err)
	}
}

func (d *Database) GetUserByDate(month int, day int) ([]User, error) {
	var users []User
	d.db.Where("DATE_PART('month', birth) = ? and DATE_PART('day', birth) = ?", month, day).Find(&users)

	return users, nil
}
