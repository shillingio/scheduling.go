package scheduling

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	expiryCushion = time.Duration(10 * time.Minute)
)

type API struct {
	Client *APIProto
	Config Config

	token       *oauth2.Token
	credentials clientcredentials.Config
}

type Config struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	TokenURL     string
}

func NewAPI(config ...Config) (*API, error) {

	api := API{}

	if len(config) > 0 {
		api.Config = config[0]
	}

	if api.Config.BaseURL == "" {
		api.Config.BaseURL = strings.TrimSpace(os.Getenv(EnvBaseURL))
	}

	if api.Config.BaseURL == "" {
		api.Config.TokenURL = strings.TrimSpace(os.Getenv(EnvTokenURL))
	}

	if api.Config.ClientID == "" {
		api.Config.ClientID = strings.TrimSpace(os.Getenv(EnvClientID))
	}

	if api.Config.ClientSecret == "" {
		api.Config.ClientSecret = strings.TrimSpace(os.Getenv(EnvClientSecret))
	}

	if api.Config.BaseURL == "" {
		return nil, ErrMissingBaseURL
	}
	if api.Config.TokenURL == "" {
		return nil, ErrMissingTokenURL
	}
	if api.Config.ClientID == "" {
		return nil, ErrMissingClientID
	}
	if api.Config.ClientSecret == "" {
		return nil, ErrMissingClientSecret
	}

	u, err := url.Parse(api.Config.BaseURL)
	if err != nil {
		return nil, err
	}

	transport := httptransport.New(u.Host, "/scheduling", []string{u.Scheme})
	transport.Transport = SetUserAgent(transport.Transport, "scheduling-client-go/0.1")

	api.Client = New(transport, nil)

	return &api, nil
}

func SetUserAgent(inner http.RoundTripper, userAgent string) http.RoundTripper {
	return &addUGA{
		inner: inner,
		Agent: userAgent,
	}
}

type addUGA struct {
	inner http.RoundTripper
	Agent string
}

func (ug *addUGA) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", ug.Agent)
	return ug.inner.RoundTrip(r)
}
