package route

import (
	"config"
	"data"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const (
	layoutFN   = "layout"
	errorMsg   = "msg"
	errorUrl   = "/err?msg="
	indexUrl   = "/"
	loginURL   = "/login"
	cookieName = "_cookie"
)

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("%s\\%s.html", config.AppConfig.ServerConfig.TemplateBase, file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, layoutFN, data)
}

func errorMessage(w http.ResponseWriter, r *http.Request, message string) {
	url := []string{}
	http.Redirect(w, r, strings.Join(url, ""), http.StatusFound)
}

func verifySession(w http.ResponseWriter, r *http.Request) (data.Session, error) {
	cookie, err := r.Cookie(cookieName)
	session := data.Session{}
	if err != nil {
		return session, err
	}
	session.Uuid = cookie.Value
	if ok, _ := session.Check(); !ok {
		return session, errors.New("invalid session")
	}
	return session, nil
}
