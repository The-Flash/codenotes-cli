package notes

type InitConfig struct {
	Directory string
}

type noteConfig struct {
	message  string
	path     string
	isSticky bool
	lineNo   int
}

type noteConfigOption func(*noteConfig)

func NewNoteConfig(message string, path string, options ...noteConfigOption) (*noteConfig, error) {
	config := noteConfig{
		message: message,
		path:    path,
	}

	for _, setter := range options {
		setter(&config)
	}
	err := validateNoteConfig(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func WithIsSticky(isSticky bool) noteConfigOption {
	return func(nc *noteConfig) {
		nc.isSticky = isSticky
	}
}

func WithLineNo(lineNo int) noteConfigOption {
	return func(nc *noteConfig) {
		nc.lineNo = lineNo
	}
}

func validateNoteConfig(config *noteConfig) error {
	_ = config
	// message and path are required
	// if a path is a folder, then the following must be true

	return nil
}
