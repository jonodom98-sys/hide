package mocks

import (
	"github.com/hide-org/hide/pkg/model"
	"github.com/stretchr/testify/mock"
)

type MockLanguageDetector struct {
	mock.Mock
}

func (m *MockLanguageDetector) DetectLanguage(file *model.File) string {
	args := m.Called(file)
	return args.String(0)
}

func (m *MockLanguageDetector) DetectMainLanguage(files []*model.File) string {
	args := m.Called(files)
	return args.String(0)
}

func (m *MockLanguageDetector) DetectLanguages(files []*model.File) map[string]int {
	args := m.Called(files)
	var result map[string]int
	if r := args.Get(0); r != nil {
		result = r.(map[string]int)
	}
	return result
}
