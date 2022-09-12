package handlers

import (
	"net/http"

	"github.com/kuklaph/bookings/pkg/config"
	"github.com/kuklaph/bookings/pkg/models"
	"github.com/kuklaph/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (r *Repository) Home(w http.ResponseWriter, re *http.Request) {
	remoteIP := re.RemoteAddr
	r.App.Session.Put(re.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (r *Repository) About(w http.ResponseWriter, re *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := r.App.Session.GetString(re.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
