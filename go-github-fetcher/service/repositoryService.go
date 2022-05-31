package service

import (
	Entity "go-github-fetcher/entity"
	Helpers "go-github-fetcher/helpers"
	Repo "go-github-fetcher/repository"

	"log"
	"sync"
)

type RepositoriesResponseService interface {
	RetrieveRepositoriesAndSaveIt(url string)
}

func RetrieveRepositoriesAndSaveIt(url string) {
	resp, err := HttpCall(url)
	if err != nil {
		log.Fatal(err)
	}

	repositoriesResponse := Helpers.ConvertBytesToRepositoriesResponse(resp)

	//Insert Repositories
	var idToDelete string

	for i := 0; i < len(repositoriesResponse.Items); i++ {
		if i == 0 {
			idToDelete = Repo.InsertRepository(repositoriesResponse.Items[i])
		}
		Repo.InsertRepository(repositoriesResponse.Items[i])
	}
	Repo.RemoveAllRepositoryBefore(idToDelete)

	//Extract concurently LicenseStatistics and LanguageStatistics
	var wg sync.WaitGroup

	var languageStatistics []Entity.LanguageStatistics
	var licenseStatistics []Entity.LicenseStatistics

	go ExtractLanguageFromRepositoriesResponse(repositoriesResponse, &languageStatistics, &wg)
	wg.Add(1)

	go ExtractLicenseFromRepositoriesRepsonse(repositoriesResponse, &licenseStatistics, &wg)
	wg.Add(1)

	wg.Wait()

	//Insert them
	TruncateAndInsertManyLanguageStatistics(languageStatistics)
	TruncateAndInsertManyLicenseStatistics(licenseStatistics)
}
