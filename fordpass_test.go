// +build integration

package fordpass

import (
	"os"
	"testing"
)

func TestClient_Authenticate(t *testing.T) {
	c := NewClient(
		os.Getenv("FORDPASS_USERNAME"),
		os.Getenv("FORDPASS_PASSWORD"),
		os.Getenv("VIN"))

	err := c.Authenticate()
	if err != nil {
		t.Error(err)
	}
}
