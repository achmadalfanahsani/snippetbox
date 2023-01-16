package main

import "github.com/achmadalfanahsani/sippetbox/internal/models"

// struct templateData digunakan sebagai struct 
// induk untuk data dinamis yang ingin diberikan
// ke template HTML.
type templateData struct {
	Snippet *models.Snippet
}