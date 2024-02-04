package halper

import (
	"TransKuliner/model/entity"
	"TransKuliner/model/response"
)

func CategoryToCategorySomeResponse(category entity.Category) response.CategorySomeResponse {
	categorySomeResponse := response.CategorySomeResponse{
		ID:   category.ID,
		Name: category.Name,
	}
	return categorySomeResponse
}
func ProductToProductSomeResponse(product entity.Product) response.ProductSomeResponse {
	productSomeResponse := response.ProductSomeResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
	return productSomeResponse
}
func ProductResponseToProductSomeResponse(product response.ProductResponse) response.ProductSomeResponse {
	productSomeResponse := response.ProductSomeResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
	return productSomeResponse
}

func CustomerToCustomerSomeResonse(customer entity.Customer) response.CustomerSomeResponse {
	customerSomeResonse := response.CustomerSomeResponse{
		ID:   customer.ID,
		Name: customer.Name,
	}
	return customerSomeResonse
}
func CustomerResponseToCustomerSomeResonse(customer response.CustomerResponse) response.CustomerSomeResponse {
	customerSomeResonse := response.CustomerSomeResponse{
		ID:   customer.ID,
		Name: customer.Name,
	}
	return customerSomeResonse
}
