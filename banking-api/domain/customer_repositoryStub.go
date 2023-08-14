package domain

type CustomerRepoStub struct {
	customers []Customer
}

func (s *CustomerRepoStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepoStub() *CustomerRepoStub {
	cus := []Customer{
		{Id: "1001", Name: "Filip", City: "Sofia", Zipcode: "1606", DateOfBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Rob", City: "Sofia", Zipcode: "1609", DateOfBirth: "2003-01-28", Status: "1"},
	}

	return &CustomerRepoStub{
		customers: cus,
	}
}
