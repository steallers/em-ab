package repository

import (
	. "github.com/steallers/em-ab/servers/model"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func CreateCustomerRepository(dbConnectors *gorm.DB) CustomerInterface {
	// Somewhat about customer Initializing
	return CustomerRepository{
		db: dbConnectors,
	}
}

func (repo CustomerRepository) InsertCustomer(customer *Customer) error {
	return repo.db.Create(&customer).Error
}

func (repo CustomerRepository) UpdateCustomer(customer *Customer) error {
	return repo.db.Model(&Customers{}).Updates(&customer).Error
}

func (repo CustomerRepository) DeleteCustomer(customerID uint64) error {
	return repo.db.Delete(&Customers{}, customerID).Error
}

func (repo CustomerRepository) FindCustomerByID(customerID uint64) (customer Customer, err error) {
	return customer, repo.db.Where("id = ?", customerID).First(&customer).Error
}


func (repo CustomerRepository) GetAllCustomerWithCondition(condition map[string]string) {

}
