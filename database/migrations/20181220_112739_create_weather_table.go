package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateWeatherTable_20181220_112739 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateWeatherTable_20181220_112739{}
	m.Created = "20181220_112739"

	migration.Register("CreateWeatherTable_20181220_112739", m)
}

// Run the migrations
func (m *CreateWeatherTable_20181220_112739) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE IF NOT EXISTS weather (
		id integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
		location_name varchar(255) NOT NULL DEFAULT '' ,
		temperature varchar(255) NOT NULL DEFAULT '' ,
		wind varchar(255) NOT NULL DEFAULT '' ,
		pressure varchar(255) NOT NULL DEFAULT '' ,
		humidity varchar(255) NOT NULL DEFAULT '' ,
		sunrise datetime NOT NULL,
		sunset datetime NOT NULL,` +
		"`lon` double NOT NULL DEFAULT 0," +
		`lat double NOT NULL DEFAULT 0,
		requested_time datetime NOT NULL,
		timestamp datetime NOT NULL
	) ENGINE=InnoDB;`)

}

// Reverse the migrations
func (m *CreateWeatherTable_20181220_112739) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE 'weather'")
}
