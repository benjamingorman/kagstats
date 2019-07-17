package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Harrison-Miller/kagstats/common/models"
	"github.com/gorilla/mux"
)


func getServers(w http.ResponseWriter, r *http.Request) {
	var servers []models.Server
	err := db.Select(&servers, "SELECT * FROM servers")
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, struct {
		Servers []models.Server `json:"servers"`
	}{
		Servers: servers,
	})
}

func getServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serverID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Could not parse id", http.StatusBadRequest)
		return
	}

	var server models.Server
	err = db.Get(&server, "SELECT * FROM servers WHERE ID=?", int64(serverID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Server not found: %v", err), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, &server)
}
