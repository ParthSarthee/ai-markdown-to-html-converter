package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/russross/blackfriday/v2"
)

func main(){
	fmt.Println("Main function started")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", HandleHome)
	mux.HandleFunc("GET /preview", HandlePreview)
	mux.HandleFunc("POST /api/convert", HandleConvert)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server started on :8080")
	defer fmt.Println("Server stopped")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
	
	// Serve the editor.html template
	tmpl, err := template.ParseFiles("templates/editor.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		fmt.Println("Template error:", err)
		return
	}
	
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Println("Template execution error:", err)
	}
}

// PreviewData represents the data passed to the preview template
type PreviewData struct {
	HTML     template.HTML
	Markdown string
}

func HandlePreview(w http.ResponseWriter, r *http.Request) {
	// Get HTML content from query parameters or session
	htmlContent := r.URL.Query().Get("html")
	markdownContent := r.URL.Query().Get("markdown")
	
	// If no content provided, show empty preview
	data := PreviewData{
		HTML:     template.HTML(htmlContent),
		Markdown: markdownContent,
	}
	
	// Serve the preview.html template
	tmpl, err := template.ParseFiles("templates/preview.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		fmt.Println("Template error:", err)
		return
	}
	
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		fmt.Println("Template execution error:", err)
	}
}

// ConvertRequest represents the JSON request structure
type ConvertRequest struct {
	Markdown string `json:"markdown"`
}

// ConvertResponse represents the JSON response structure
type ConvertResponse struct {
	HTML string `json:"html"`
}

// API Route
func HandleConvert(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Parse JSON request
	var req ConvertRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Check if markdown is empty
	if req.Markdown == "" {
		response := ConvertResponse{
			HTML: "<p>No markdown content provided.</p>",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	// Convert markdown to HTML using blackfriday
	// Configure blackfriday with common extensions
	extensions := blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags | blackfriday.HrefTargetBlank,
	})
	
	htmlBytes := blackfriday.Run([]byte(req.Markdown), blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(renderer))
	
	// Create response
	response := ConvertResponse{
		HTML: string(htmlBytes),
	}
	
	// Send JSON response
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}