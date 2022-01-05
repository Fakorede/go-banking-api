package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "John", "US", "111", "2000-01-01", "1"},
		{"1002", "Anthony", "UK", "222", "2000-01-01", "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
