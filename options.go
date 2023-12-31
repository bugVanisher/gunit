package gunit

type option func(*configuration)

var Options singleton

type singleton struct{}

// SkipAll is an option meant to be passed to gunit.Run(...)
// and causes each and every "Test" method in the corresponding
// fixture to be skipped (as if each had been prefixed with
// "Skip"). Even "Test" methods marked with the "Focus" prefix
// will be skipped.
func (singleton) SkipAll() option {
	return func(this *configuration) {
		this.SkippedTestCases = true
	}
}

// LongRunning is an option meant to be passed to
// gunit.Run(...) and, in the case that the -short testing
// flag has been passed at the command line, it causes each
// and every "Test" method in the corresponding fixture to
// be skipped (as if each had been prefixed with "Skip").
func (singleton) LongRunning() option {
	return func(this *configuration) {
		this.LongRunningTestCases = true
	}
}

// SequentialTestCases is an option meant to be passed to
// gunit.Run(...) and prevents gunit from calling t.Parallel()
// on the inner instances of *testing.T passed to the 'subtests'
// corresponding to "Test" methods which are created during
// the natural course of the corresponding invocation of
// gunit.Run(...).
func (singleton) SequentialTestCases() option {
	return func(this *configuration) {
		this.SequentialTestCases = true
	}
}

// composite allows graceful chaining of options.
func (singleton) composite(options ...option) option {
	return func(this *configuration) {
		for _, option := range options {
			option(this)
		}
	}
}
