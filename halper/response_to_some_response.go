package halper

import "TransKuliner/model/response"

func CustomerResponseToCustomerSomeResonse(customer response.CustomerResponse) response.CustomerSomeResponse {
	customerSomeResonse := response.CustomerSomeResponse{
		ID:   customer.ID,
		Name: customer.Name,
	}
	return customerSomeResonse
}
func ProductResponseToProductSomeResponse(product response.ProductResponse) response.ProductSomeResponse {
	productSomeResponse := response.ProductSomeResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
	return productSomeResponse
}
