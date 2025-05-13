package options

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/srxstack/gin-template/internal/apiserver"
	genericoptions "github.com/srxstack/srxstack/pkg/options"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
)

// 定义支持的服务器模式集合。
var availableServerModes = sets.New(
	"Release",
	"Debug",
	"Test",
)

// ServerOptions 包含服务器配置选项。
type ServerOptions struct {
	// ServerMode 定义 gin 服务器模式：Release、Debug、Test。
	ServerMode  string                      `json:"server-mode" mapstructure:"server-mode"`
	JWTKey      string                      `json:"jwt-key" mapstructure:"jwt-key"`
	Expiration  time.Duration               `json:"expiration" mapstructure:"expiration"`
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`
	TLSOptions  *genericoptions.TLSOptions  `json:"tls" mapstructure:"tls"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		ServerMode:  "Debug",
		JWTKey:      "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		Expiration:  2 * time.Hour,
		HTTPOptions: genericoptions.NewHTTPOptions(),
		TLSOptions:  genericoptions.NewTLSOptions(),
	}
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.ServerMode, "server-mode", o.ServerMode, fmt.Sprintf("Server mode, available options: %v", availableServerModes.UnsortedList()))
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "JWT signing key. Must be at least 6 characters long.")
	fs.DurationVar(&o.Expiration, "expiration", o.Expiration, "The expiration duration of JWT tokens.")
	o.HTTPOptions.AddFlags(fs)
	o.TLSOptions.AddFlags(fs)
}

func (o *ServerOptions) Validate() error {
	errs := []error{}

	if !availableServerModes.Has(o.ServerMode) {
		errs = append(errs, fmt.Errorf("invalid server mode: must be one of %v", availableServerModes.UnsortedList()))
	}

	if len(o.JWTKey) < 6 {
		errs = append(errs, errors.New("JWTKey must be at least 6 characters long"))
	}

	errs = append(errs, o.HTTPOptions.Validate()...)

	return utilerrors.NewAggregate(errs)
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		ServerMode:  o.ServerMode,
		JWTKey:      o.JWTKey,
		Expiration:  o.Expiration,
		HTTPOptions: o.HTTPOptions,
		TLSOptions:  o.TLSOptions,
	}, nil
}
