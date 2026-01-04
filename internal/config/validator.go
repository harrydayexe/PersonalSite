package config

// Validator is an interface for types that can validate themselves.
type Validator interface {
	// Validate returns an error if the value is invalid, nil otherwise.
	Validate() error
}
