package scheduling

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/shillingio/scheduling.go/scheduler"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	userAgent = "scheduling-client-go/0.1"
)

var (
	expiryCushion = time.Duration(10 * time.Minute)
)

type API struct {
	Client *APIProto
	Config Config

	HTTPClient *http.Client

	token       *oauth2.Token
	credentials clientcredentials.Config
}

type Config struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
	TokenURL     string
	Debug        bool

	Options []scheduler.ClientOption
}

func NewAPI(config ...Config) (*API, error) {

	api := API{
		HTTPClient: http.DefaultClient,
	}

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
	transport.Transport = SetUserAgent(transport.Transport, userAgent)
	if api.Config.Debug {
		transport.Transport = DebugRoundTripper(transport.Transport)
	}

	api.Client = New(transport, nil)

	// fmt.Println("Authenticate")
	// if err := api.Authenticate(); err != nil {
	// 	return nil, err
	// }
	// fmt.Println("Authenticated", api.token.AccessToken)

	return &api, nil
}

type debugRoundTripper struct {
	base http.RoundTripper
}

func (rt debugRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("---------------------------------")
	fmt.Println("------------ REQUEST ------------")
	fmt.Println("URL:", req.URL.String())
	fmt.Println("Method:", req.Method)
	fmt.Println("Headers:")

	req.Header.Set("User-Agent", userAgent)

	for i := range req.Header {
		fmt.Println(" ", i, ":", req.Header.Get(i))
	}
	if req.Body != nil {
		fmt.Println("Body:")
		reqBody := req.Body
		body, err := ioutil.ReadAll(reqBody)
		if err != nil {
			return nil, err
		}
		if body != nil {
			var out interface{}
			if err = json.Unmarshal(body, &out); err == nil {
				d, err := json.MarshalIndent(out, "", "  ")
				if err != nil {
					return nil, err
				}
				fmt.Println(string(d))
			}
			if err != nil {
				fmt.Println("Error unmarshalling request body:", err.Error())
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}

	resp, err := rt.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("--------------------------------")
	fmt.Println("---                          ---")
	fmt.Println("----------- RESPONSE -----------")
	fmt.Println("Status:", resp.Status)
	fmt.Println("Status Code:", resp.StatusCode)
	if resp != nil {
		fmt.Println("Headers:")
		for i := range resp.Header {
			fmt.Println(" ", i, ":", resp.Header.Get(i))
		}
	}
	if resp.Body != nil {
		fmt.Println("Body:")
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading body", err.Error())
		}
		if body != nil {
			var out interface{}
			if err = json.Unmarshal(body, &out); err == nil {
				// if err != nil {
				// 	return nil, err
				// }

				d, err := json.MarshalIndent(out, "", "  ")
				if err != nil {
					return nil, err
				}
				fmt.Println(string(d))
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}

	// fmt.Println("resp", *resp)
	fmt.Println("--------- END RESPONSE ---------")
	fmt.Println("--------------------------------")
	return resp, nil
}

func DebugRoundTripper(base http.RoundTripper) http.RoundTripper {
	return debugRoundTripper{base: base}
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
