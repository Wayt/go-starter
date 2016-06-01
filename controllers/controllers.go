package controllers

import (
	"github.com/wayt/go-starter/models"
)

type controller struct {
	Model models.Model
}

var controllers []controller

func register(m models.Model) controller {
	ctl := controller{
		Model: m,
	}

	controllers = append(controllers, ctl)

	return ctl
}
