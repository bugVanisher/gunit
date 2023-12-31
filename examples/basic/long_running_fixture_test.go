package basic

import (
	"testing"

	"github.com/bugVanisher/gunit"
)

func TestHowToMarkAnEntireFixtureAsLongRunning(t *testing.T) {
	gunit.Run(new(HowToMarkAnEntireFixtureAsLongRunning), t,
		gunit.Options.LongRunning(), // <-- Just pass the LongRunning option to gunit.Run(...)!
	)
}

type HowToMarkAnEntireFixtureAsLongRunning struct {
	*gunit.Fixture
}

func (this *HowToMarkAnEntireFixtureAsLongRunning) TestTheseTestsWillBeSkippedIfTheShortFlagWasPassedAtTheCommandLine() {
}
