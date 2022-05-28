package helpers

import (
	Entity "go-github-fetcher/entity"
)

/*
Function to find if a language statistics already exist in the slice
*/
func FindLanguageStatistics(slice []Entity.LanguageStatistics, val string) int {
	for i, item := range slice {
		if item.Language == val {
			return i
		}
	}
	return -1
}

/*
Function to find if a license statistics already exist in the slice
*/
func FindLicenseStatistics(slice []Entity.LicenseStatistics, val string) int {
	for i, item := range slice {
		if item.License == val {
			return i
		}
	}
	return -1
}

/*
Function to convert a slice of LanguageStatistics to a slice of interface
*/
func LanguageStatisticsSliceToInterfaceSlice(lss []Entity.LanguageStatistics) []interface{} {
	var iS []interface{}

	for _, b := range lss {
		iS = append(iS, b)
	}

	return iS
}

/*
Function to convert a slice of LicenseStatistics to a slice of interface
*/
func LicenseStatisticsSliceToInterfaceSlice(lss []Entity.LicenseStatistics) []interface{} {
	var iS []interface{}

	for _, b := range lss {
		iS = append(iS, b)
	}

	return iS
}
