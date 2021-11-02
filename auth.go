package scheduling

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (a *API) AuthInfoWriter() runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+a.token.AccessToken)
		return nil
	})
}

func (a *API) WithAuth(op *runtime.ClientOperation) {

	if a.token == nil || time.Now().After(a.token.Expiry.Add(-expiryCushion)) {
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
			http.DefaultClient,
		))
		if err != nil {
			fmt.Println("token error", err.Error())
			return
		}
	}

	op.AuthInfo = a.AuthInfoWriter()
}
