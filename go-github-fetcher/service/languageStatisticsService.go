package service

import (
	Entity "go-github-fetcher/entity"
	Helpers "go-github-fetcher/helpers"
	Repo "go-github-fetcher/repository"
	"sync"
)

type LanguageStatisticsService interface {
	TruncateAndInsertManyLanguageStatistics(languageStats []Entity.LanguageStatistics)
	ExtractLanguageFromRepositoriesResponse(repositoriesResponse Entity.RepositoriesResponse, languageStats []Entity.LanguageStatistics, wg *sync.WaitGroup)
}

func TruncateAndInsertManyLanguageStatistics(languageStats []Entity.LanguageStatistics) {
	Repo.RemoveAllLanguageStatistics()
	Repo.InsertManyLanguageStatistics(languageStats)
}

func ExtractLanguageFromRepositoriesResponse(repositoriesResponse Entity.RepositoriesResponse, languageStats *[]Entity.LanguageStatistics, wg *sync.WaitGroup) {
	defer wg.Done()

	var tmpLanguageStats []Entity.LanguageStatistics
	var key int

	for i := 0; i < len(repositoriesResponse.Items); i++ {
		key = Helpers.FindLanguageStatistics(tmpLanguageStats, repositoriesResponse.Items[i].Language)
		if key == -1 {
			tmpLanguageStats = append(tmpLanguageStats, Entity.LanguageStatistics{Language: repositoriesResponse.Items[i].Language, Total: 1})
		} else {
			tmpLanguageStats[key].Total++
		}
	}

	*languageStats = tmpLanguageStats
}
