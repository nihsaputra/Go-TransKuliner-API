package halper

import (
	"TransKuliner/model/entity"
	"TransKuliner/model/response"
)

func SaleToSaleResponse(sale entity.Sale) response.SaleResponse {
	saleResponse := response.SaleResponse{
		ID:        sale.ID,
		Product:   ProductToProductSomeResponse(sale.Product),
		Customer:  CustomerToCustomerSomeResonse(sale.Customer),
		CreatedAt: sale.CreatedAt,
	}
	return saleResponse
}
