package languages_test

import (
	"gofilm/bussinesses/languages"
	languageMock "gofilm/bussinesses/languages/mocks"
	"os"
	"testing"
)

var (
	langRepository languageMock.LangRepository
	langUseCase languages.LangUseCase
	langThird languageMock.LangFromAPI
)

func setup() {
	langUseCase = languages.NewService(&langRepository, &langThird)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}


