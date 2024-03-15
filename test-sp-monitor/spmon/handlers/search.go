package handlers

import (
	"sp-monitoring/models"
	"strings"
)

func SearchPage(searchQuery string, data []models.DynamicData) []models.DynamicData {

	// Filter data based on the search query
	var filteredData []models.DynamicData

	for _, item := range data {
		// Customize this part based on your actual search logic
		if strings.Contains(item.Secret, searchQuery) ||
			strings.Contains(item.Metadata, searchQuery) ||
			strings.Contains(item.Keyvault, searchQuery) {
			filteredData = append(filteredData, item)
		}
	}

	return filteredData
}
