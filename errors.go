package scheduling

import "fmt"

var (
	ErrMissingBaseURL      = fmt.Errorf("new scheduling api: missing base URL, pass in a complete config or set the environment variable %s", EnvBaseURL)
	ErrMissingTokenURL     = fmt.Errorf("new scheduling api: missing token URL, pass in a complete config or set the environment variable %s", EnvTokenURL)
	ErrMissingClientID     = fmt.Errorf("new scheduling api: missing client id, pass in a complete config or set the environment variable %s", EnvClientID)
	ErrMissingClientSecret = fmt.Errorf("new scheduling api: missing client secret, pass in a complete config or set the environment variable %s", EnvClientSecret)
)
