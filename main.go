package main

import (
	"log"
	"os"
	"time"

	"github.com/kardianos/service"
)

const (
	serviceName        = "reposync"
	serviceDisplayName = "Repository sync service"
	serviceDescription = "Watches a directory for bundle artifacts and applies them to remote repositories."

	pollInterval = 10 * time.Second
)

type program struct {
	exit chan struct{}
}

func main() {
	svcConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceDisplayName,
		Description: serviceDescription,
	}

	args := os.Args[1:]
	p := &program{}

	s, err := service.New(p, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(args) > 0 {
		// Puller runs once and returns bundles to process
		if args[0] == "pull" {
			if err := pull(); err != nil {
				log.Fatal(err)
			}
			return
		}

		// Support for installing/starting/stopping the service via command line arguments
		err = service.Control(s, args[0])
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Run as a standalone process if no arguments are provided
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Performs a single pull-run and returns.
func pull() error {
	return nil
}

func (p *program) Start(s service.Service) error {
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

// Runs the main processing function continuously
func (p *program) run() {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-p.exit:
			return
		case <-ticker.C:
			processBundles()
		}
	}
}

func processBundles() {
}

func (p *program) Stop(s service.Service) error {
	close(p.exit)
	return nil
}
