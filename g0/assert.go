// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type checker struct {
	vu   modules.VU
	fail bool
}

func (c *checker) Errorf(format string, args ...interface{}) {
	state := c.vu.State()
	if state == nil || !state.Logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}

	str := fmt.Sprintf(format, args...)
	idx := strings.Index(str, "\n\tError:")

	state.Logger.Debug("Assertion failed", strings.ReplaceAll(str[idx:], "\t", ""))
}

func (c *checker) Check(name string, succ bool) {
	state := c.vu.State()

	if state == nil {
		return
	}

	if len(name) == 0 {
		name = "unnamed assertion"
	}

	commonTagsAndMeta := state.Tags.GetCurrentValues()

	tags := commonTagsAndMeta.Tags
	if state.Options.SystemTags.Has(metrics.TagCheck) {
		tags = tags.With("check", name)
	}

	check, err := state.Group.Check(name)
	if err != nil {
		return
	}

	ctx := c.vu.Context()
	now := time.Now()

	select {
	case <-ctx.Done():
	default:
		sample := metrics.Sample{
			TimeSeries: metrics.TimeSeries{
				Metric: state.BuiltinMetrics.Checks,
				Tags:   tags,
			},
			Time:     now,
			Metadata: commonTagsAndMeta.Metadata,
			Value:    0,
		}

		if succ {
			atomic.AddInt64(&check.Passes, 1)

			sample.Value = 1
		} else {
			atomic.AddInt64(&check.Fails, 1)
		}

		metrics.PushIfNotDone(ctx, state.Samples, sample)
	}

	if !succ && c.fail {
		common.Throw(c.vu.Runtime(), fmt.Errorf("%w: %s", errAssertionFailed, name))
	}
}

func (c *checker) FailNow() {
	if c.fail {
		common.Throw(c.vu.Runtime(), errAssertionFailed)
	}
}

var (
	_ assert.TestingT  = (*checker)(nil)
	_ require.TestingT = (*checker)(nil)

	errAssertionFailed = errors.New("assertion failed")
)
