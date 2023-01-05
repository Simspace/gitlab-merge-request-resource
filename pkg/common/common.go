package common

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

// AuthedTransport provides a mechanism for GraphQL requests to provide authorization
type AuthedTransport struct {
	Key     string
	Wrapped http.RoundTripper
}

// RoundTrip sets the Authorization header and wraps http.RoundTrip
func (a *AuthedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+a.Key)
	return a.Wrapped.RoundTrip(req)
}

func Fatal(doing string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", doing, err)
	os.Exit(1)
}

func GetDefaultClient(insecure bool) *http.Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	return http.DefaultClient
}
