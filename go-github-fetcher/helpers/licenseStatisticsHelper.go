package helpers

import (
	Entity "go-github-fetcher/entity"
)

//FindLicenseStatistics find if a license statistics already exist in the slice
func FindLicenseStatistics(slice []Entity.LicenseStatistics, val string) int {
	for i, item := range slice {
		if item.License == val {
			return i
		}
	}
	return -1
}

//LicenseStatisticsSliceToInterfaceSlice convert a slice of LicenseStatistics to a slice of interface
func LicenseStatisticsSliceToInterfaceSlice(lss []Entity.LicenseStatistics) []interface{} {
	var iS []interface{}

	for _, b := range lss {
		iS = append(iS, b)
	}

	return iS
}
