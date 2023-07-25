package main

import "snippetbox.achmadalfanahsani.com/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}