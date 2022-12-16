// run as administrator is not enough
// this script should run as root (e.g. with service)

package main

import (
	"simple/logger"
	"simple/utils"
	"time"

	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	logger.Log.Infof("service is running ...")
	utils.GetCredential()
	for {
		time.Sleep(time.Second)
	}
}

func (p *program) Stop(s service.Service) error {
	logger.Log.Info("service stop ...")
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceExampleSimple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Log.Error(err)
	}

	s.Run()
}
