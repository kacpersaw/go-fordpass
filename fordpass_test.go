// +build integration

package fordpass

import (
	"os"
	"testing"
)

func TestClient_Status(t *testing.T) {
	c := NewClient(
		os.Getenv("FORDPASS_USERNAME"),
		os.Getenv("FORDPASS_PASSWORD"),
		os.Getenv("FORDPASS_VIN"))

	status, err := c.Status()
	if err != nil {
		t.Errorf("TestClient_Status() Error %v", err.Error())
	}

	if status.Vin != c.Vin {
		t.Errorf("TestClient_Status() got: %s want: %s", status.Vin, c.Vin)
	}
}
