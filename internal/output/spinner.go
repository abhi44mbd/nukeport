package output

import (
	"time"

	"github.com/theckman/yacspin"
)

func StartSpinner(message string) (*yacspin.Spinner, error) {
	cfg := yacspin.Config{
		Frequency: 100 * time.Millisecond,
		CharSet:   yacspin.CharSets[14],
		Suffix:    " " + message,
	}

	spinner, err := yacspin.New(cfg)

	if err != nil {
		return nil, err
	}

	err = spinner.Start()

	if err != nil {
		return nil, err
	}

	return spinner, nil
}
