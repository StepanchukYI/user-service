//go:build unit
// +build unit

package websocket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_StatusCheckReadiness(t *testing.T) {
	var srv Server

	srv.ready = false
	assert.Equal(t, "wss service: not started yet", srv.HealthCheck().Error())
}
