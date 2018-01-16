package main

import (
	"testing"

	"github.com/devimc/api"
)

func TestExercise(t *testing.T) {
	if exercise() != false {
		t.Fatal("Should be false")
	}

	api.ValidUserFuncImpl = api.MockValidUser
	if exercise() != false {
		t.Fatal("Should be false")
	}

	api.WriteLogFuncImpl = api.MockWriteLog
	if exercise() != true {
		t.Fatal("Should be true")
	}
}
