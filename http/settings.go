package http

import (
	"encoding/json"
	"net/http"

	"github.com/filebrowser/filebrowser/v2/rules"
	"github.com/filebrowser/filebrowser/v2/settings"
)

type settingsData struct {
	Signup           bool                  `json:"signup"`
	CreateUserDir    bool                  `json:"createUserDir"`
	ShowLogout	     bool                  `json:"showLogout"`
	UserHomeBasePath string                `json:"userHomeBasePath"`
	LogoutPath		 string          	   `json:"logoutPath"`
	Defaults         settings.UserDefaults `json:"defaults"`
	Rules            []rules.Rule          `json:"rules"`
	Branding         settings.Branding     `json:"branding"`
	Tus              settings.Tus          `json:"tus"`
	Shell            []string              `json:"shell"`
	Commands         map[string][]string   `json:"commands"`
}

var settingsGetHandler = withAdmin(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	data := &settingsData{
		Signup:           d.settings.Signup,
		CreateUserDir:    d.settings.CreateUserDir,
		ShowLogout:	      d.settings.ShowLogout,
		UserHomeBasePath: d.settings.UserHomeBasePath,
		LogoutPath: 	  d.settings.LogoutPath,
		Defaults:         d.settings.Defaults,
		Rules:            d.settings.Rules,
		Branding:         d.settings.Branding,
		Tus:              d.settings.Tus,
		Shell:            d.settings.Shell,
		Commands:         d.settings.Commands,
	}

	return renderJSON(w, r, data)
})

var settingsPutHandler = withAdmin(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	req := &settingsData{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, err
	}

	d.settings.Signup = req.Signup
	d.settings.CreateUserDir = req.CreateUserDir
	d.settings.ShowLogout = req.ShowLogout
	d.settings.UserHomeBasePath = req.UserHomeBasePath
	d.settings.LogoutPath = req.LogoutPath
	d.settings.Defaults = req.Defaults
	d.settings.Rules = req.Rules
	d.settings.Branding = req.Branding
	d.settings.Tus = req.Tus
	d.settings.Shell = req.Shell
	d.settings.Commands = req.Commands

	err = d.store.Settings.Save(d.settings)
	return errToStatus(err), err
})
