package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name            string
	TypeID          uint
	CustomerContact CustomerContact
}

func (Customer) TableName() string {
	return "dbmain.md-customers"
}

// CustomerType : Default(1: Personal, 2: Company, 3: Contractor/ThirdParty)
type CustomerType struct {
	gorm.Model
	Description string
}

func (CustomerType) TableName() string {
	return "dbmain.cd-customer_type"
}

type Customers []Customer
