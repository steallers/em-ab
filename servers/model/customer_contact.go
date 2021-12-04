package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CustomerContact struct {
	ID            string //combination datetime and others
	CustomerID    uint
	PICName       string
	ContactTypeID uint
	ContactType   ContactType
	Value         string
}

func (CustomerContact) TableName() string {
	return "dbmain.hd-customer_contact"
}

type ContactType struct {
	gorm.Model
	Name        string
	Description string
	LogoURL     string
	HTMLRef     string
}

func (ContactType) TableName() string {
	return "dbmain.cd-contact_type"
}

func (cc *CustomerContact) GenerateID() {
	cc.ID = fmt.Sprintf("%d_%d", time.Now().UnixNano(), cc.CustomerID)
}
