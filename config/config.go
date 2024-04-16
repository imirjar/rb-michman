package config

import (
	"github.com/imirjar/Michman/config/diver"
	"github.com/imirjar/Michman/config/michman"
)

func NewDiverConfig() *diver.DiverConfig {
	return diver.NewDiverConfig()
}

func NewMichmanConfig() *michman.MichmanConfig {
	return michman.NewMichmanConfig()
}
