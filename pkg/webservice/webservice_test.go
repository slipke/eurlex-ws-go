package webservice

import "testing"

func TestCreateWebservice(t *testing.T) {
	u := "testuser"
	p := "testpass"

	cfg := NewConfig(u, p)

	ws := NewWebservice(cfg)

	if ws.cfg != cfg {
		t.Errorf("Config not set properly, got: %+v, want: %+v", ws.cfg, cfg)
	}
}
