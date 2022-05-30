package helper_test

import (
	Entity "go-github-fetcher/entity"
	Helper "go-github-fetcher/helpers"
	"reflect"

	"testing"
)

func initLanguageStatisticsSliceForTesting() []Entity.LanguageStatistics {
	var languageStats []Entity.LanguageStatistics

	languageStats = append(languageStats, Entity.LanguageStatistics{Language: "Typescript", Total: 2})
	languageStats = append(languageStats, Entity.LanguageStatistics{Language: "Java", Total: 1})
	languageStats = append(languageStats, Entity.LanguageStatistics{Language: "Golang", Total: 10})

	return languageStats
}

func TestInitLanguageStatisticsForTesting(t *testing.T) {
	languageStats := initLanguageStatisticsSliceForTesting()

	if len(languageStats) != 3 {
		t.Error("The language slice generated for testing should contains 3 languages")
	}
}

func TestFindLanguageStatisticsExist(t *testing.T) {
	languageStats := initLanguageStatisticsSliceForTesting()
	languageNameToFound := "Golang"

	key := Helper.FindLanguageStatistics(languageStats, languageNameToFound)
	if key == -1 {
		t.Error("The language 'Golang' should be present but was not found")
	}
}

func TestFindLanguageStatisticsNotExist(t *testing.T) {
	languageStats := initLanguageStatisticsSliceForTesting()
	languageNameToFound := "C#"

	key := Helper.FindLanguageStatistics(languageStats, languageNameToFound)
	if key != -1 {
		t.Error("The language 'C#' should not be present but was found")
	}
}

func TestLanguageStatisticsSliceToInterfaceSlice(t *testing.T) {
	languageStats := initLanguageStatisticsSliceForTesting()

	languageStatsInterface := Helper.LanguageStatisticsSliceToInterfaceSlice(languageStats)

	if reflect.TypeOf(languageStatsInterface).String() != "[]interface {}" {
		t.Error("The language slice should be converted into an interface slice")
	}

	if len(languageStatsInterface) != 3 {
		t.Error("The interface slice should contains same number of element than the language slice, expected 3 but get ", len(languageStatsInterface))
	}
}
