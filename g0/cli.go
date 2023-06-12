// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package g0

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func isRunCommand(args []string) (bool, int) {
	argn := len(args)

	scriptIndex := argn - 1
	if scriptIndex < 0 {
		return false, scriptIndex
	}

	var runIndex int

	for idx := 0; idx < argn; idx++ {
		arg := args[idx]
		if arg == "run" && runIndex == 0 {
			runIndex = idx

			break
		}
	}

	if runIndex == 0 {
		return false, -1
	}

	return true, scriptIndex
}

func redirectStdin() {
	isRun, scriptIndex := isRunCommand(os.Args)
	if !isRun {
		return
	}

	script := os.Args[scriptIndex]
	if script == "-" || !strings.HasSuffix(script, ".go") {
		return
	}

	os.Setenv(envScript, script)

	os.Args[scriptIndex] = "-"

	reader, writer, err := os.Pipe()
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	defer writer.Close()

	origStdin := os.Stdin

	os.Stdin = reader

	_, err = writer.Write([]byte(jsScript))
	if err != nil {
		writer.Close()

		os.Stdin = origStdin

		logrus.WithError(err).Fatal()
	}
}

const jsScript = `//js
export * from 'k6/x/g0'
export { default } from 'k6/x/g0'
import execution from 'k6/execution'
global.execution = execution
`
