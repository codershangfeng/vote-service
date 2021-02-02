// +build unit

package handler

import (
	"testing"

	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations/probe"
)

func TestGetHealthHandler(t *testing.T) {
	got := GetHealthHandler(probe.NewGetHealthParams())

	if got.(*probe.GetHealthOK) == probe.NewGetHealthOK() {
		t.Errorf("Expected NewGetHealthOK response, but got: %v1", got)
	}
}
