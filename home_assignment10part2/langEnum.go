package main

type Language string

const (
	French  Language = "fr"
	English Language = "en"
	German  Language = "de"
)

func IsValidLanguage(lang string) bool {
	validLanguages := make(map[Language]bool)
	validLanguages[French] = true
	validLanguages[English] = true
	validLanguages[German] = true

	_, isValid := validLanguages[Language(lang)]
	return isValid
}
