package main

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/transform"
)

var log = NewWithPath("service.log")

type ToCRLF struct{}

func (ToCRLF) Reset() {}

func (ToCRLF) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nDst < len(dst) && nSrc < len(src) {
		if c := src[nSrc]; c == '\n' {
			if nDst+1 == len(dst) {
				break
			}
			dst[nDst] = '\r'
			dst[nDst+1] = '\n'
			nSrc++
			nDst += 2
		} else {
			dst[nDst] = c
			nSrc++
			nDst++
		}
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

type RotateLogWriter struct {
	lock sync.Mutex
	path string
	fp   *os.File
}

func (logger *RotateLogWriter) Write(p []byte) (n int, err error) {
	logger.lock.Lock()
	defer logger.lock.Unlock()

	if logger.fp == nil {
		logger.fp, err = os.OpenFile(logger.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			log.Println("Failed to open log at '" + logger.path + "'")
			return
		}
	}

	return logger.fp.Write(p)
}

func NewWithPath(logPath string) (log *logrus.Logger) {
	log = logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
			funcname := frame.Function[strings.LastIndex(frame.Function, ".")+1:]
			line := strconv.Itoa(frame.Line)
			_, filename := path.Split(frame.File)
			filename = filename + ":" + line

			return funcname, filename
		},
	})

	log.SetOutput(transform.NewWriter(&RotateLogWriter{path: logPath}, ToCRLF{}))

	return
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	log.Infof("service is running ...")
	for {
		log.Infof("-------------------")
		time.Sleep(time.Second)
	}
}

func (p *program) Stop(s service.Service) error {
	log.Info("service stop ...")
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
		log.Error(err)
	}

	s.Run()
}
