package app

import (
	"github.com/sirupsen/logrus"
	
	"log"
)


type Application struct {
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