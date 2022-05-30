package helper_test

import (
	Entity "go-github-fetcher/entity"
	Helper "go-github-fetcher/helpers"
	"reflect"

	"testing"
)

func initLicenseStatisticsSliceForTesting() []Entity.LicenseStatistics {
	var licenseStats []Entity.LicenseStatistics

	licenseStats = append(licenseStats, Entity.LicenseStatistics{License: "MIT", Total: 2})
	licenseStats = append(licenseStats, Entity.LicenseStatistics{License: "Apache2", Total: 1})
	licenseStats = append(licenseStats, Entity.LicenseStatistics{License: "BSD3", Total: 10})

	return licenseStats
}

func TestInitLicenseStatisticsForTesting(t *testing.T) {
	licenseStats := initLicenseStatisticsSliceForTesting()

	if len(licenseStats) != 3 {
		t.Error("The license slice generated for testing should contains 3 licenses")
	}
}

func TestFindLicenseStatisticsExist(t *testing.T) {
	licenseStats := initLicenseStatisticsSliceForTesting()
	licenseNameToFound := "MIT"

	key := Helper.FindLicenseStatistics(licenseStats, licenseNameToFound)
	if key == -1 {
		t.Error("The license 'MIT' should be present but was not found")
	}
}

func TestFindLicenseStatisticsNotExist(t *testing.T) {
	licenseStats := initLicenseStatisticsSliceForTesting()
	licenseNameToFound := "DWTFYW"

	key := Helper.FindLicenseStatistics(licenseStats, licenseNameToFound)
	if key != -1 {
		t.Error("The license 'DWTFYW' should not be present but was found")
	}
}

func TestLicenseStatisticsSliceToInterfaceSlice(t *testing.T) {
	licenseStats := initLicenseStatisticsSliceForTesting()

	licenseStatsInterface := Helper.LicenseStatisticsSliceToInterfaceSlice(licenseStats)

	if reflect.TypeOf(licenseStatsInterface).String() != "[]interface {}" {
		t.Error("The license slice should be converted into an interface slice")
	}

	if len(licenseStatsInterface) != 3 {
		t.Error("The interface slice should contains same number of element than the license slice, expected 3 but get ", len(licenseStatsInterface))
	}
}
