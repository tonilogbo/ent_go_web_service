package artists

import (
	"encoding/json"
	"entrds/cors"
	"entrds/ent"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const artistPath = "artists"

func errorMessage(message string) []byte {
	return []byte(fmt.Sprintf(`{"message":"%s"}`,message))
}

func handleArtists(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		artistList, err := getArtistList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(artistList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var artist ent.Artist
		err := json.NewDecoder(r.Body).Decode(&artist)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newArtist, err := insertArtist(artist)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		j, err := json.Marshal(newArtist)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
		}	
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleArtist(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/",artistPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idString := urlPathSegments[len(urlPathSegments)-1]
	if idString == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMessage("No id given!"))
		return
	}
	artistID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		artist, err := getArtist(artistID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if artist == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(artist)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPut:
		var artist ent.Artist
		err := json.NewDecoder(r.Body).Decode(&artist)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		updatedArtist, err := updateArtist(artistID, artist)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if updatedArtist == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(updatedArtist)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if _, err = w.Write(j); err != nil {
			log.Fatal(err)
		}
	case http.MethodDelete:
		if err := deleteArtist(artistID); err != nil {
			if err.Error() == ArtistNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(errorMessage("Artist deleted."))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupRoutes(apiBasePath string) {
	artistsHandler := http.HandlerFunc(handleArtists)
	artistHandler := http.HandlerFunc(handleArtist)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, artistPath), cors.Middleware(artistsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, artistPath), cors.Middleware(artistHandler))
}