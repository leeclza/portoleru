// Local dev server: serves public/ and the app router on :3000.
package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/leeclza/portofolio-leon/internal/app"
)

func main() {
	router := app.Router()
	static := http.FileServer(http.Dir("public"))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// same precedence as Vercel: existing files in public/ first
		if r.URL.Path != "/" {
			if fi, err := os.Stat(filepath.Join("public", filepath.FromSlash(r.URL.Path))); err == nil && !fi.IsDir() {
				static.ServeHTTP(w, r)
				return
			}
		}
		router.ServeHTTP(w, r)
	})

	log.Println("listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
