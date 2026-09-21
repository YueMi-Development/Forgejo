// Copyright 2024 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package test

import (
	"testing"

	"forgejo.org/modules/testlogger"
)

// DeclareExpectedErrors marks the listed log message substrings as
// expected for the duration of the current test. The messages will
// still be logged at their normal level, but they will not cause the
// test to fail. Registration is removed automatically when the test
// finishes.
//
// Use this for negative tests that intentionally exercise invalid
// input or invalid state — for example, OAuth callbacks that
// exercise bad credentials, project moves that submit non-existent
// issues, or pull-request merges that fail approvals. Production
// logging is unchanged; only the test logger's failure escalation
// is suppressed for the registered substrings.
func DeclareExpectedErrors(t testing.TB, patterns ...string) {
	t.Helper()
	testlogger.ExpectedErrors.Add(t.Name(), patterns...)
	t.Cleanup(func() { testlogger.ExpectedErrors.Remove(t.Name()) })
}
