package helpers

import (
	Entity "go-github-fetcher/entity"
)

//FindLanguageStatistics find if a language statistics already exist in the slice
func FindLanguageStatistics(slice []Entity.LanguageStatistics, val string) int {
	for i, item := range slice {
		if item.Language == val {
			return i
		}
	}
	return -1
}

//LanguageStatisticsSliceToInterfaceSlice convert a slice of LanguageStatistics to a slice of interface
func LanguageStatisticsSliceToInterfaceSlice(lss []Entity.LanguageStatistics) []interface{} {
	var iS []interface{}

	for _, b := range lss {
		iS = append(iS, b)
	}

	return iS
}
