package config

import "errors"

// Validasi error yang spesifik
var (
	ErrMissingField = errors.New("missing required field")
)

func (c *Config) Validate() error {
	if c.DBURL == "" {
		return ErrMissingField
	}
	if c.Port == "" {
		return ErrMissingField
	}
	if c.JWTAuthSecret == "" {
		return ErrMissingField
	}
	if c.JWTAuthExpInHour.String() == "" {
		return ErrMissingField
	}
	if c.JWTRefreshSecret == "" {
		return ErrMissingField
	}
	if c.JWTRefreshExpInHour.String() == "" {
		return ErrMissingField
	}
	return nil
}
