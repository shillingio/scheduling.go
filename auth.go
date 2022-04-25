package scheduling

import (
	"context"
	"fmt"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (a *API) Authenticate() error {

	var refresh bool

	switch {
	case a == nil:
		// fmt.Println("API is nil")
		return fmt.Errorf("API is nil")
	case a.token == nil:
		refresh = true
	case a.token != nil && time.Now().After(a.token.Expiry.Add(-expiryCushion)):
		refresh = true
	}

	if refresh {
		if a.Config.Debug {
			fmt.Println("Authenticate")
			fmt.Println("  OAuth Token URL :", a.Config.TokenURL)
			fmt.Println("  Client ID       :", a.Config.ClientID)
			fmt.Println("  Client Secret   :", a.Config.ClientSecret)
		}

		a.credentials = clientcredentials.Config{
			ClientID:     a.Config.ClientID,
			ClientSecret: a.Config.ClientSecret,
			Scopes:       []string{"offline"},
			TokenURL:     a.Config.TokenURL,
		}

		var err error
		a.token, err = a.credentials.Token(context.WithValue(
			context.Background(),
			oauth2.HTTPClient,
			a.HTTPClient,
		))
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *API) AuthInfoWriter() runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+a.token.AccessToken)
		return nil
	})
}

func (a *API) WithAuth(op *runtime.ClientOperation) {

	if err := a.Authenticate(); err != nil {
		fmt.Println("Error authenticating:", err.Error())
		return
	}

	op.AuthInfo = a.AuthInfoWriter()
}
