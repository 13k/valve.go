package steamweb_test

import (
	"os"
	"time"

	"github.com/13k/valve.go/steamweb"
)

func ExampleNew_inline() {
	// Inline options
	_, _ = steamweb.New(
		steamweb.WithUserAgent("mylib 1.2.3"),
		steamweb.WithKey("<api_key>"),
		steamweb.WithTimeout(3*time.Second),
		steamweb.WithRetryCount(5),
	)
}

func ExampleNew_conditional() {
	// Conditional options
	options := []steamweb.ClientOption{
		steamweb.WithKey("<api_key>"),
	}

	if os.Getenv("DEBUG") != "" {
		options = append(options, steamweb.WithDebug())
	}

	_, _ = steamweb.New(options...)
}
