package main

import (
	"bytes"
	"html/template"
	"strings"
	"testing"
)

func loadTemplates(t *testing.T) *template.Template {
	t.Helper()
	funcMap := template.FuncMap{
		"add":   func(a, b int) int { return a + b },
		"split": strings.Split,
	}
	tmpl, err := template.New("").Funcs(funcMap).ParseGlob("templates/*.html")
	if err != nil {
		t.Fatalf("failed to parse templates: %v", err)
	}
	return tmpl
}

func TestHomeTemplate(t *testing.T) {
	tmpl := loadTemplates(t)

	data := struct {
		Recipes []RecipeSimple
	}{
		Recipes: []RecipeSimple{
			{ID: 1, Title: "Spaghetti Carbonara", TimeMinutes: 25, Price: "12.50", Tags: []Tag{{ID: 1, Name: "Italian"}}},
		},
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "home.html", data); err != nil {
		t.Fatalf("failed to render home.html: %v", err)
	}

	body := buf.String()

	for _, want := range []string{"Spaghetti Carbonara", "Italian", "25", "12.50", "VIEW RECIPE"} {
		if !strings.Contains(body, want) {
			t.Errorf("expected rendered home.html to contain %q", want)
		}
	}
}

func TestHomeTemplateEmpty(t *testing.T) {
	tmpl := loadTemplates(t)

	data := struct{ Recipes []RecipeSimple }{Recipes: nil}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "home.html", data); err != nil {
		t.Fatalf("failed to render home.html with no recipes: %v", err)
	}

	if !strings.Contains(buf.String(), "No recipes found") {
		t.Error("expected empty state message when no recipes")
	}
}

func TestRecipeDetailTemplate(t *testing.T) {
	tmpl := loadTemplates(t)

	recipe := &Recipe{
		ID:          1,
		Title:       "Garlic Butter Salmon",
		TimeMinutes: 20,
		Price:       "22.00",
		Link:        "http://example.com/salmon",
		Description: "Step 1: Season the salmon.\n\nStep 2: Cook in butter.",
		Ingredients: []Ingredient{
			{ID: 1, Name: "Salmon Fillet", Amount: "4", Unit: "fillets"},
			{ID: 2, Name: "Butter", Amount: "3", Unit: "tbsp"},
		},
		Tags: []Tag{
			{ID: 1, Name: "Seafood"},
			{ID: 2, Name: "Healthy"},
		},
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "recipe_detail.html", recipe); err != nil {
		t.Fatalf("failed to render recipe_detail.html: %v", err)
	}

	body := buf.String()

	for _, want := range []string{
		"Garlic Butter Salmon",
		"20",
		"22.00",
		"http://example.com/salmon",
		"Salmon Fillet",
		"Butter",
		"Seafood",
		"Healthy",
		"Step 1: Season the salmon.",
		"Step 2: Cook in butter.",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("expected rendered recipe_detail.html to contain %q", want)
		}
	}
}
