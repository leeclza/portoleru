package app

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/leeclza/portofolio-leon/pkg/data"
	"github.com/leeclza/portofolio-leon/pkg/views"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, views.Home())
	})

	mux.HandleFunc("GET /pengalaman", func(w http.ResponseWriter, r *http.Request) {
		tab, it := pengalamanQuery(r)
		render(w, r, views.PengalamanPage(tab, it))
	})

	mux.HandleFunc("GET /pengalaman/body", func(w http.ResponseWriter, r *http.Request) {
		tab, it := pengalamanQuery(r)
		render(w, r, views.PengalamanBody(tab, it))
	})

	mux.HandleFunc("GET /pengalaman/item", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		item, ok := data.FindItem(r.URL.Query().Get("tab"), id)
		if !ok {
			http.NotFound(w, r)
			return
		}
		render(w, r, views.Modal(item))
	})

	mux.HandleFunc("GET /pengalaman/close", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// anything no other route matches
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		views.NotFound(r.URL.Path).Render(r.Context(), w)
	})

	return mux
}

func pengalamanQuery(r *http.Request) (string, bool) {
	tab := r.URL.Query().Get("tab")
	if !data.ValidTab(tab) {
		tab = "kepanitiaan"
	}
	// IT filter exists on the committees tab and on All
	it := (tab == "kepanitiaan" || tab == data.TabAll) && r.URL.Query().Get("it") == "1"
	return tab, it
}

func render(w http.ResponseWriter, r *http.Request, c templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
