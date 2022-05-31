package service

import (
	Entity "go-github-fetcher/entity"
	Helpers "go-github-fetcher/helpers"
	Repo "go-github-fetcher/repository"

	"sync"
)

type LicenseStatisticsService interface {
	TruncateAndInsertManyLicenseStatistics(licenseStats []Entity.LicenseStatistics)
	ExtractLicenseFromRepositoriesRepsonse(repositoriesResponse Entity.RepositoriesResponse, licenseStats []Entity.LicenseStatistics, wg *sync.WaitGroup)
}

func TruncateAndInsertManyLicenseStatistics(licenseStats []Entity.LicenseStatistics) {
	Repo.RemoveAllLicenseStatistics()
	Repo.InsertManyLicenseStatistics(licenseStats)
}

func ExtractLicenseFromRepositoriesRepsonse(repositoriesResponse Entity.RepositoriesResponse, licenseStats *[]Entity.LicenseStatistics, wg *sync.WaitGroup) {
	defer wg.Done()

	var tmpLicenseStats []Entity.LicenseStatistics
	var key int

	for i := 0; i < len(repositoriesResponse.Items); i++ {
		key = Helpers.FindLicenseStatistics(tmpLicenseStats, repositoriesResponse.Items[i].License.Name)
		if key == -1 {
			tmpLicenseStats = append(tmpLicenseStats, Entity.LicenseStatistics{License: repositoriesResponse.Items[i].License.Name, Total: 1})
		} else {
			tmpLicenseStats[key].Total++
		}
	}

	*licenseStats = tmpLicenseStats
}
