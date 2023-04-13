package models

import (

	"time"

)

type Medidor struct {

	Id  				int				`gorm:"<-:create; not null; type:varchar(100); primarykey; unique_index" json:"id" swaggertype:"integer"`
	Brand 				string			`gorm:"<-:create; not null; type:varchar(100)" json:"brand"`
	Address 			string 			`gorm:"not null; type:varchar(1000)" json:"address"`
	Serial 				string 			`gorm:"<-:create; not null; type:varchar(100)" json:"serial"`
	Lines 				uint8 			`gorm:"not null; default=0; size:10" json:"lines" swaggertype:"integer"`
	IsActive 			bool 			`gorm:"not null; default=false" json:"is_active"`
	InstallationDate 	time.Time 		`gorm:"<-:create; not null" json:"installation_date" example:"0000-12-31T19:00:00-05:00"`
	RetirementDate	 	time.Time		`json:"retirement_date" example:"0000-12-31T19:00:00-05:00" `
	CreatedAt 			time.Time 		`gorm:"not null" json:"created_at" swaggerignore:"true"`

}

type Medidores []*Medidor 