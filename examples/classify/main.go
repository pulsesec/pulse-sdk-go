package main

import (
	"context"
	"errors"
	"os"

	pulse "github.com/pulsesec/pulse-sdk-go"
)

var (
	client = pulse.New(os.Getenv("PULSE_SITE_KEY"), os.Getenv("PULSE_SECRET_KEY"))
)

func classify(token string) bool {
	isBot, err := client.Classify(context.Background(), token)
	if err != nil {
		if errors.Is(err, pulse.ErrTokenNotFound) {
			panic("Token not found")
		}

		if errors.Is(err, pulse.ErrTokenUsed) {
			panic("Token already used")
		}

		if errors.Is(err, pulse.ErrTokenExpired) {
			panic("Token expired")
		}

		panic(err)
	}

	return isBot
}

var _ = classify
