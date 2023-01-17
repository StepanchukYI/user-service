//go:build unit
// +build unit

package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_StatusCheckReadiness(t *testing.T) {
	var srv Server

	srv.ready = false
	assert.Equal(t, "http service: not started yet", srv.HealthCheck().Error())
}
