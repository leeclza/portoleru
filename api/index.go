// Vercel Go Function entry point. vercel.json rewrites every non-static
// request here; static files are served by Vercel from public/.
package handler

import (
	"net/http"

	"github.com/leeclza/portofolio-leon/pkg/app"
)

var router = app.Router()

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
