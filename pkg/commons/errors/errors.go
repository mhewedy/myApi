package errors

type E struct {
	Key     string
	Details []string
	Client  bool
}

func (h *E) Error() string {
	return h.Key
}

// New Client error
func New(key string, extra ...string) error {
	return &E{
		Key:     key,
		Client:  true,
		Details: extra,
	}
}

// New server error
func Err(key string, extra ...string) error {
	return &E{
		Key:     key,
		Client:  false,
		Details: extra,
	}
}
