package main

import "ed-henrique.snippetbox/internal/models"

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}
