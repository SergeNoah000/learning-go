package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// Servir les fichiers statiques du frontend
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/", apis.Static(os.DirFS("./public"), true))
		return se.Next()
	})

	// Hook avant création d'article - associer l'auteur
	app.OnRecordCreate("articles").BindFunc(func(e *core.RecordEvent) error {
		if e.Record.GetString("author") == "" {
			authRecord := e.RequestInfo().Auth
			if authRecord != nil {
				e.Record.Set("author", authRecord.Id)
			}
		}
		return e.Next()
	})

	// Hook avant création de commentaire - associer l'utilisateur
	app.OnRecordCreate("comments").BindFunc(func(e *core.RecordEvent) error {
		if e.Record.GetString("user") == "" {
			authRecord := e.RequestInfo().Auth
			if authRecord != nil {
				e.Record.Set("user", authRecord.Id)
			}
		}
		return e.Next()
	})

	// Hook avant création de like - associer l'utilisateur
	app.OnRecordCreate("likes").BindFunc(func(e *core.RecordEvent) error {
		if e.Record.GetString("user") == "" {
			authRecord := e.RequestInfo().Auth
			if authRecord != nil {
				e.Record.Set("user", authRecord.Id)
			}
		}
		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
