// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package requirements

import (
  "fmt"
	assert "github.com/stretchr/testify/assert"
	http "net/http"
	url "net/url"
	time "time"
)

func (a *Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Condition(a.t, comp, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Conditionf(comp assert.Comparison, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Conditionf(a.t, comp, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Contains(a.t, s, contains, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Containsf(a.t, s, contains, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) DirExists(path string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.DirExists(a.t, path, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) DirExistsf(path string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.DirExistsf(a.t, path, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ElementsMatch(a.t, listA, listB, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ElementsMatchf(a.t, listA, listB, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Empty(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Emptyf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Emptyf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Equal(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.EqualError(a.t, theError, errString, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.EqualErrorf(a.t, theError, errString, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.EqualValues(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.EqualValuesf(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Equalf(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Error(err error, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Error(a.t, err, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ErrorAs(a.t, err, target, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ErrorAsf(err error, target interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ErrorAsf(a.t, err, target, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ErrorContains(theError error, contains string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ErrorContains(a.t, theError, contains, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ErrorContainsf(theError error, contains string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ErrorContainsf(a.t, theError, contains, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ErrorIs(a.t, err, target, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.ErrorIsf(a.t, err, target, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Errorf(err error, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Errorf(a.t, err, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Eventually(a.t, condition, waitFor, tick, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Eventuallyf(a.t, condition, waitFor, tick, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Exactly(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Exactlyf(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Fail(a.t, failureMessage, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.FailNow(a.t, failureMessage, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) FailNowf(failureMessage string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.FailNowf(a.t, failureMessage, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Failf(failureMessage string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Failf(a.t, failureMessage, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) False(value bool, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.False(a.t, value, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Falsef(value bool, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Falsef(a.t, value, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) FileExists(path string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.FileExists(a.t, path, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) FileExistsf(path string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.FileExistsf(a.t, path, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Greater(a.t, e1, e2, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.GreaterOrEqual(a.t, e1, e2, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.GreaterOrEqualf(a.t, e1, e2, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Greaterf(a.t, e1, e2, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPBodyContains(a.t, handler, method, url, values, str, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPBodyContainsf(a.t, handler, method, url, values, str, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPBodyNotContains(a.t, handler, method, url, values, str, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPBodyNotContainsf(a.t, handler, method, url, values, str, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPError(a.t, handler, method, url, values, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPErrorf(a.t, handler, method, url, values, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPRedirect(a.t, handler, method, url, values, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPRedirectf(a.t, handler, method, url, values, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPStatusCode(a.t, handler, method, url, values, statuscode, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPStatusCodef(a.t, handler, method, url, values, statuscode, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPSuccess(a.t, handler, method, url, values, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.HTTPSuccessf(a.t, handler, method, url, values, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Implements(a.t, interfaceObject, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Implementsf(a.t, interfaceObject, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InDelta(a.t, expected, actual, delta, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InDeltaMapValues(a.t, expected, actual, delta, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InDeltaMapValuesf(a.t, expected, actual, delta, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InDeltaSlice(a.t, expected, actual, delta, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InDeltaSlicef(a.t, expected, actual, delta, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InDeltaf(a.t, expected, actual, delta, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InEpsilon(a.t, expected, actual, epsilon, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InEpsilonSlice(a.t, expected, actual, epsilon, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InEpsilonSlicef(a.t, expected, actual, epsilon, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.InEpsilonf(a.t, expected, actual, epsilon, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsDecreasing(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsDecreasing(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsDecreasingf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsDecreasingf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsIncreasing(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsIncreasing(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsIncreasingf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsIncreasingf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsNonDecreasing(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsNonDecreasingf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsNonDecreasingf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsNonIncreasing(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsNonIncreasingf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsNonIncreasingf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsType(a.t, expectedType, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.IsTypef(a.t, expectedType, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.JSONEq(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.JSONEqf(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Len(a.t, object, length, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Lenf(object interface{}, length int, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Lenf(a.t, object, length, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Less(a.t, e1, e2, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.LessOrEqual(a.t, e1, e2, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.LessOrEqualf(a.t, e1, e2, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Lessf(a.t, e1, e2, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Negative(e interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Negative(a.t, e, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Negativef(e interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Negativef(a.t, e, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Never(a.t, condition, waitFor, tick, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Neverf(a.t, condition, waitFor, tick, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Nil(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Nilf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Nilf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NoDirExists(path string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NoDirExists(a.t, path, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NoDirExistsf(path string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NoDirExistsf(a.t, path, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NoError(a.t, err, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NoErrorf(err error, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NoErrorf(a.t, err, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NoFileExists(path string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NoFileExists(a.t, path, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NoFileExistsf(path string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NoFileExistsf(a.t, path, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotContains(a.t, s, contains, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotContainsf(a.t, s, contains, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotEmpty(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotEmptyf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotEmptyf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotEqual(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotEqualValues(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotEqualValuesf(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotEqualf(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotErrorIs(a.t, err, target, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotErrorIsf(a.t, err, target, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotNil(a.t, object, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotNilf(object interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotNilf(a.t, object, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotPanics(a.t, f, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotPanicsf(f assert.PanicTestFunc, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotPanicsf(a.t, f, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotRegexp(a.t, rx, str, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotRegexpf(a.t, rx, str, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotSame(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotSamef(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotSubset(a.t, list, subset, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotSubsetf(a.t, list, subset, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotZero(a.t, i, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) NotZerof(i interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.NotZerof(a.t, i, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Panics(a.t, f, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) PanicsWithError(errString string, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.PanicsWithError(a.t, errString, f, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) PanicsWithErrorf(errString string, f assert.PanicTestFunc, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.PanicsWithErrorf(a.t, errString, f, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.PanicsWithValue(a.t, expected, f, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) PanicsWithValuef(expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.PanicsWithValuef(a.t, expected, f, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Panicsf(f assert.PanicTestFunc, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Panicsf(a.t, f, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Positive(e interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Positive(a.t, e, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Positivef(e interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Positivef(a.t, e, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Regexp(a.t, rx, str, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Regexpf(a.t, rx, str, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Same(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Samef(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Subset(a.t, list, subset, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Subsetf(a.t, list, subset, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) True(value bool, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.True(a.t, value, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Truef(value bool, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Truef(a.t, value, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.WithinDuration(a.t, expected, actual, delta, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.WithinDurationf(a.t, expected, actual, delta, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.WithinRange(a.t, actual, start, end, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.WithinRangef(a.t, actual, start, end, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.YAMLEq(a.t, expected, actual, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.YAMLEqf(a.t, expected, actual, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Zero(a.t, i, msgAndArgs...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprint(msgAndArgs...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}

func (a *Assertions) Zerof(i interface{}, msg string, args ...interface{}) {
	if h, ok := a.t.(tHelper); ok {
		h.Helper()
	}
	succ := assert.Zerof(a.t, i, msg, args...)
	if c, ok := a.t.(tChecker); ok {
		name := fmt.Sprintf(msg, args...)
		c.Check(name, succ)
	}
	if !succ {
		a.t.FailNow()
	}
}
