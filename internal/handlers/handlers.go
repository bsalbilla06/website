package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// PageData holds data to be passed to templates
type PageData struct {
	Title      string
	ActivePage string
}

// RenderTemplate parses and executes the given template file
func RenderTemplate(w http.ResponseWriter, page string, data PageData) {
	// Paths to templates
	basePath := filepath.Join("web", "templates", "layouts", "base.html")
	pagePath := filepath.Join("web", "templates", "pages", page+".html")

	// Parse templates
	tmpl, err := template.ParseFiles(basePath, pagePath)
	if err != nil {
		log.Printf("Error parsing template %s: %v", page, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute template
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", page, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HomeHandler handles the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := PageData{
		Title:      "Home",
		ActivePage: "home",
	}
	RenderTemplate(w, "home", data)
}

// AboutHandler handles the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:      "About Me",
		ActivePage: "about",
	}
	RenderTemplate(w, "about", data)
}

// ExperienceHandler handles the experience page
func ExperienceHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:      "Experience",
		ActivePage: "experience",
	}
	RenderTemplate(w, "experience", data)
}

// ContactHandler handles the contact page
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:      "Contact",
		ActivePage: "contact",
	}
	RenderTemplate(w, "contact", data)
}
