package main

import (
	"fmt"
	"net/http"

	"github.com/TopoVista/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
   ____    ____        _    ____ ___ 
  / ___|  / ___|      / \  |  _ \_ _|
 | |  _  | |   _____  / _ \ | |_) | | 
 | |_| | | |__|_____/ ___ \|  __/| | 
  \____|  \____|    /_/   \_\_|  |___|
`)
    err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}


}	