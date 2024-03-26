package cli

import (
	"github.com/briandowns/spinner"
	"time"
)

type SpinnerModel struct {
	SP *spinner.Spinner
}

func Spinner(mode int, Suffix string) SpinnerModel {
	s := spinner.New(spinner.CharSets[mode], 100*time.Millisecond)
	s.Suffix = " " + Suffix
	s.Start()
	return SpinnerModel{
		SP: s,
	}

}

func (s SpinnerModel) Stop() {
	s.SP.Stop()
}
