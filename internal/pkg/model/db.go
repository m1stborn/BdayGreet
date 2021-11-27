package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var DB = Database{}

func (d *Database) Init(uri string) {

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		fmt.Println("DB connection Error")
	}

	d.db = db

	//drop old table
	if dropErr := d.db.Migrator().DropTable(&User{}); dropErr != nil {
		//logger.WithFields(log.Fields{
		//	"pkg":  "model",
		//	"func": "Init",
		//}).Error(dropErr)
		fmt.Println("DropTable Error", dropErr)
	}

	//create all the table
	if !d.db.Migrator().HasTable(&User{}) {
		err = d.db.Migrator().CreateTable(&User{})
	} else {
		err = d.db.Migrator().AutoMigrate(&User{})
	}

	// insert all the data
	for _, user := range DBRecords {
		d.CreateUser(&user)
	}
}
