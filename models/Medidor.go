package models

import (

	"time"

)

type Medidor struct {

	//gorm.Model

	Id  				int				`gorm:"<-:create; not null; type:varchar(100); primarykey; unique_index" json:"id"`
	Brand 				string			`gorm:"<-:create; not null; type:varchar(100)" json:"brand"`
	Address 			string 			`gorm:"not null; type:varchar(1000)" json:"address"`
	Serial 				string 			`gorm:"<-:create; not null; type:varchar(100)" json:"serial"`
	Lines 				uint8 			`gorm:"not null; default=0; size:10" json:"lines"`
	IsActive 			bool 			`gorm:"not null; default=false" json:"is_active"`
	InstallationDate 	time.Time 		`gorm:"<-:create; not null" json:"installation_date"`
	RetirementDate	 	time.Time		`json:"retirement_date"`
	CreatedAt 			time.Time 		`gorm:"not null" json:"created_at"`

}

type Medidores []*Medidor