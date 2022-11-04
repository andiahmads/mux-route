package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "raisa", City: "jakarta", Zipcode: "0100", DateOfBirth: "2020-21-12", Status: "1"},
		{Id: "1002", Name: "raisa1", City: "jakarta1", Zipcode: "01001", DateOfBirth: "2020-21-12", Status: "1"},
		{Id: "1003", Name: "raisa2", City: "jakarta2", Zipcode: "01002", DateOfBirth: "2020-21-12", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
