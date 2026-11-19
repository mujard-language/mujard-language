package main

// The shared client for the developer API (docs/DEVELOPER_API_REFERENCE.md).
// One place owns: where the API is, how the key is found, the Bearer header,
// idempotency keys on POSTs, and turning the {"error":{...}} envelope into one
// readable line. Commands stay thin wrappers over call/callRaw.
import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type apiClient struct {
	base string // ends in /api/v1
	key  string // mjd_live_... / mjd_test_... - never printed
}

// apiV1 turns the API folder (where health.php lives) into the developer API root.
// Passing a root that already ends in /api/v1 is fine.
func apiV1(apiRoot string) string {
	r := strings.TrimRight(apiRoot, "/")
	if strings.HasSuffix(r, "/api/v1") {
		return r
	}
	return r + "/api/v1"
}

// addAPIFlags registers --api and --key and resolves them flag > env > saved config > default.
func addAPIFlags(fs *flag.FlagSet) *apiClient {
	c := &apiClient{}
	saved := loadConfig()
	fs.StringVar(&c.base, "api", firstNonEmpty(envGet("MUJARD_API"), saved.API, "https://mujard.com/mujard-api"), "the server's API folder (where health.php is); /api/v1 is added")
	fs.StringVar(&c.key, "key", firstNonEmpty(envGet("MUJARD_KEY"), saved.Key), "a developer key (mjd_live_... / mjd_test_...); saved by \"mujard login\"")
	return c
}

// finish resolves the base to the v1 root after flag parsing.
func (c *apiClient) finish() { c.base = apiV1(c.base) }

// needKey is the polite refusal every keyed command shares.
func (c *apiClient) needKey(w io.Writer) bool {
	if c.key != "" {
		return false
	}
	fmt.Fprintln(w, "This command needs a developer key. Run \"mujard login --key mjd_live_...\" once, or pass --key / set MUJARD_KEY.")
	return true
}

// idemKey: a fresh Idempotency-Key (8-80 of A-Za-z0-9_-) so a retried POST never charges twice.
func idemKey() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "mujardcli-" + fmt.Sprint(time.Now().UnixNano())
	}
	return "mujardcli-" + hex.EncodeToString(b[:])
}

// call sends a JSON request and decodes the JSON answer.
func (c *apiClient) call(method, path string, body any, headers map[string]string) (int, map[string]any, error) {
	status, raw, _, err := c.callRaw(method, path, body, headers)
	if err != nil {
		return 0, nil, err
	}
	var doc map[string]any
	_ = json.Unmarshal(raw, &doc)
	return status, doc, nil
}

// callRaw is call without the JSON decode - for the artifact download, whose body is bytes.
func (c *apiClient) callRaw(method, path string, body any, headers map[string]string) (int, []byte, http.Header, error) {
	var rd io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return 0, nil, nil, err
		}
		rd = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, c.base+path, rd)
	if err != nil {
		return 0, nil, nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.key != "" {
		req.Header.Set("Authorization", "Bearer "+c.key)
	}
	// Says who is asking, so the administrator can switch builds or the AI off for the command line alone (App Control -> What works where).
	req.Header.Set(clientHeader, clientName())
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := (&http.Client{Timeout: 60 * time.Second}).Do(req)
	if err != nil {
		return 0, nil, nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 64<<20))
	return resp.StatusCode, raw, resp.Header, nil
}

// clientHeader / clientName: the X-Mujard-Client header every developer-API request carries ("mujard-cli/<version>").
const clientHeader = "X-Mujard-Client"

func clientName() string { return "mujard-cli/" + version }

// apiErrText flattens {"error":{"code","message","request_id"}} into one line for printing.
func apiErrText(doc map[string]any) string {
	e, _ := doc["error"].(map[string]any)
	if e == nil {
		return "unexpected answer from the server"
	}
	if str(e, "code") == "feature_off" {
		// the administrator switched it off for the command line: nothing to fix on this side
		return str(e, "message") + " Ask the administrator, or use the app instead."
	}
	s := str(e, "code") + ": " + str(e, "message")
	if d, _ := e["details"].(map[string]any); d != nil && str(d, "required_scope") != "" {
		s += " (this key needs the \"" + str(d, "required_scope") + "\" scope - make a new key with it)"
	}
	if id := str(e, "request_id"); id != "" {
		s += " (request id " + id + " - quote it if you ask for help)"
	}
	return s
}

// failAPI prints an API error and returns the exit code for it (1; 2 stays for usage errors).
func failAPI(w io.Writer, status int, doc map[string]any) int {
	fmt.Fprintf(w, "Failed (HTTP %d): %s\n", status, apiErrText(doc))
	return 1
}
