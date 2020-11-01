// +build integration

package fordpass

func TestClient_Status(t *testing.T) {
	car := NewClient(
		os.Getenv("FORDPASS_USERNAME"),
		os.Getenv("FORDPASS_PASSWORD"),
		os.Getenv("FORDPASS_VIN"))

	status, err := car.Status()
	if err != nil {
		t.Errorf("TestClient_Status() Error %v", err.Error())
	}

	if status.Vin != c.Vin {
		t.Errorf("TestClient_Status() got: %s want: %s", status.Vin, c.Vin)
	}
}
