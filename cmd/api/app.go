package api

import (
	"errors"
	"fmt"
	"github.com/alma-amirseitov/TgBot/cmd/bot"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Bot *bot.Bot
}

func (app *Application) Serve(addr string) error {
	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("starting server at addr: %s", addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/message", app.sendMessage)

	return router
}

func (app *Application) sendMessage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Printf("ParseForm() err: %v", err)
		return
	}
	message := r.FormValue("message")
	err := app.Bot.MessageHandler(message)
	if err != nil {
		log.Printf("%v", err)
	}
}
