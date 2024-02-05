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

func SaleDetailToSaleDetailResponse(saleDetails []entity.SaleDetail) []response.SaleDetailResponse {
	var saleDetailResponses []response.SaleDetailResponse

	for _, saleDetail := range saleDetails {
		saleDetailResponse := response.SaleDetailResponse{
			Product:    saleDetail.Product.Name,
			Category:   saleDetail.Product.Category.Name,
			Price:      saleDetail.Price,
			Quantity:   saleDetail.Quantity,
			TotalPrice: saleDetail.Price * saleDetail.Quantity,
		}
		saleDetailResponses = append(saleDetailResponses, saleDetailResponse)
	}
	return saleDetailResponses
}
