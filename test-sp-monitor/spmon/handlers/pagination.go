package handlers

import (
	"sp-monitoring/models"
	"strconv"
)

func FingCurrentPage(pageStr string, totalItems, itemsPerPage int, data []models.DynamicData) ([]models.PageInfo, []models.DynamicData) {

	// Get the current page number from the query parameters

	currentPage, _ := strconv.Atoi(pageStr)
	if currentPage == 0 {
		currentPage = 1
	}

	// Calculate the total number of pages
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage

	// Determine the start and end indices for the current page
	startIdx := (currentPage - 1) * itemsPerPage
	endIdx := currentPage * itemsPerPage
	if endIdx > len(data) {
		endIdx = len(data)
	}
	currentPageData := data[startIdx:endIdx]

	// Generate PageInfo for pagination links
	var pagination []models.PageInfo
	for i := 1; i <= totalPages; i++ {
		info := models.PageInfo{
			PageNumber: i,
			Current:    i == currentPage,
		}
		pagination = append(pagination, info)
	}

	return pagination, currentPageData

}
