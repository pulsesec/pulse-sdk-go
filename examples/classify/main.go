package main

import (
	"errors"
	"log"
	"os"

	"github.com/pulsesec/pulse-sdk-go"
)

var (
	PULSE_SITE_KEY   = os.Getenv("PULSE_SITE_KEY")
	PULSE_SECRET_KEY = os.Getenv("PULSE_SECRET_KEY")

	client = pulse.New(PULSE_SITE_KEY, PULSE_SECRET_KEY)
)

func classify(token string) {
	isBot, err := client.Classify(token)
	if err != nil {
		if errors.Is(err, pulse.ErrTokenNotFound) {
			log.Printf("Token %s not found", token)
			return
		}

		if errors.Is(err, pulse.ErrTokenUsed) {
			log.Printf("Token %s already used", token)
			return
		}

		log.Panicf("Failed to classify token %s: %v", token, err)
	}

	log.Printf("Token %s is a bot: %t", token, isBot)
}

var _ = classify
