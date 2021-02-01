// +build unit

package handler

import (
	"testing"

	"github.com/codershangfeng/vote-service/app/internal/api/restapi/operations"
)

func TestGetHealthHandler(t *testing.T) {
	got := getHealthHandler(operations.NewGetHealthParams())

	if got.(*operations.GetHealthOK) == operations.NewGetHealthOK() {
		t.Errorf("Expected NewGetHealthOK response, but got: %v1", got)
	}
}
