package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// The administrator can switch builds / the AI off for the command line alone: the server knows a request is from mujard by this header, and the
// refusal must read as "ask the administrator", not as a fault to debug.
func TestEveryRequestSaysItIsMujardAndASwitchedOffAnswerReadsPlainly(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("X-Mujard-Client")
		w.WriteHeader(403)
		_, _ = w.Write([]byte(`{"error":{"code":"feature_off","message":"Builds & uploads is switched off for the mujard command line right now.","request_id":"R1"}}`))
	}))
	defer srv.Close()

	c := &apiClient{base: srv.URL + "/api/v1", key: "mjd_test_x"}
	status, doc, err := c.call("GET", "/builds", nil, nil)
	if err != nil || status != 403 {
		t.Fatalf("call: status %d, err %v", status, err)
	}
	if got != "mujard-cli/"+version {
		t.Fatalf("X-Mujard-Client = %q, want mujard-cli/%s", got, version)
	}
	text := apiErrText(doc)
	if !strings.Contains(text, "switched off for the mujard command line") || !strings.Contains(text, "Ask the administrator") {
		t.Fatalf("the refusal should say what happened and who to ask: %q", text)
	}
	if strings.Contains(text, "feature_off") || strings.Contains(text, "request id") {
		t.Fatalf("a switch the administrator turned off is not a fault: no error code or request id in %q", text)
	}
}
