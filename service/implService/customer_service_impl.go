package implService

import (
	"TransKuliner/halper"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
	"TransKuliner/service"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
}

func (c *CustomerServiceImpl) GetAll() []response.CustomerResponse {
	var customerResponses []response.CustomerResponse
	findAll, err := c.CustomerRepository.FindAll()
	halper.PanicIfError(err)

	for _, customer := range findAll {
		customerResponse := response.CustomerResponse{
			ID:          customer.ID,
			Name:        customer.Name,
			Email:       customer.Email,
			PhoneNumber: customer.PhoneNumber,
			CreatedAt:   customer.CreatedAt,
			UpdatedAt:   customer.UpdatedAt,
		}
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses
}

func (c *CustomerServiceImpl) GetById(id uint) (response.CustomerResponse, error) {
	findById, err := c.CustomerRepository.FindById(id)

	if err != nil {
		return response.CustomerResponse{}, err
	}

	customerResponse := response.CustomerResponse{
		ID:          findById.ID,
		Name:        findById.Name,
		Email:       findById.Email,
		PhoneNumber: findById.PhoneNumber,
		CreatedAt:   findById.CreatedAt,
		UpdatedAt:   findById.UpdatedAt,
	}

	return customerResponse, nil
}

func (c *CustomerServiceImpl) Create(request request.CustomerRequest) (response.CustomerResponse, error) {
	customer := entity.Customer{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
	save, err := c.CustomerRepository.Save(customer)
	if err != nil {
		return response.CustomerResponse{}, err
	}

	customerResponse := response.CustomerResponse{
		ID:          save.ID,
		Name:        save.Name,
		Email:       save.Email,
		PhoneNumber: save.PhoneNumber,
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
	}

	return customerResponse, nil
}

func (c *CustomerServiceImpl) Update(request request.CustomerUpdateRequest) (response.CustomerResponse, error) {
	findById, err := c.CustomerRepository.FindById(request.ID)
	if err != nil {
		return response.CustomerResponse{}, err
	}

	findById.Name = request.Name
	findById.Email = request.Email
	findById.PhoneNumber = request.PhoneNumber

	save, err := c.CustomerRepository.Save(findById)
	if err != nil {
		return response.CustomerResponse{}, err
	}

	customerResponse := response.CustomerResponse{
		ID:          save.ID,
		Name:        save.Name,
		Email:       save.Email,
		PhoneNumber: save.PhoneNumber,
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
	}

	return customerResponse, nil
}

func (c *CustomerServiceImpl) Delete(id uint) error {
	findById, err := c.CustomerRepository.FindById(id)
	if err != nil {
		return err
	}
	err = c.CustomerRepository.Delete(findById)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerService(repository repository.CustomerRepository) service.CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: repository,
	}
}
