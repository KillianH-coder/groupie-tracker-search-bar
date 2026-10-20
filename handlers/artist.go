package handlers

import (
	"net/http"
	"text/template"
)

var errorMap = map[string]interface{}{
	"Code":    0,
	"Message": "",
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artistId := r.URL.Path[len("/artist/"):]
	if artistId == "" {
		tmpl := template.Must(template.ParseFiles("./templates/error.html"))
		e := errorMap
		e["Code"] = 400
		e["Message"] = "Artist ID is missing"

		tmpl.Execute(w, e)
		return
	}
	artist, status := GetArtistByID(artistId)
	switch status {
	case http.StatusOK:
		tmpl := template.Must(template.ParseFiles("./templates/artist.html"))
		tmpl.Execute(w, artist)
	case http.StatusNotFound:
		tmpl := template.Must(template.ParseFiles("./templates/error.html"))
		e := errorMap
		e["Code"] = 404
		e["Message"] = "Artist not found!"
		tmpl.Execute(w, e)
	default:
		e := errorMap
		e["Code"] = 500
		e["Message"] = "Unknown Error"
	}
}
