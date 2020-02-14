package app

import (
	"github.com/gorilla/mux"
	"github.com/iAmPlus/skills-text-extraction-go/internal/tflite"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"strconv"
)


type Application struct {
	router *mux.Router
	predictor tflite.Predictor
}

func CreateApplication() Application {
	a := Application{ }
	a.load()
	return a
}

func (a *Application) setLogFormat() {
	log.SetOutput(logrus.New().Writer())
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func (a *Application) load() {
	a.setLogFormat()
	a.loadRoutes()

	err := a.loadTFLiteModel()
	if err != nil {
		log.Println("error creating cache")
	}
}

func (a *Application) loadTFLiteModel() error {
	model :=  tflite.NewModel()
	model.CreateInterpreterPool()
	a.predictor = model
	return nil
}

func (a *Application) loadRoutes() {
	a.router = mux.NewRouter().StrictSlash(true)
	a.addRoutes()
}

func (a *Application) addRoutes() {
	a.addHealthChecks()
	a.addStackTraces()
	a.addPredictionRoute()
}

func (a *Application) StartServer(port int) error {
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(a.router)
	return http.ListenAndServe(":"+strconv.Itoa(port), n)
}