// Copyright 2023 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {

	var (
		l ILogger
	)

	BeforeEach(func() {
		l, _ = NewLogger(DebugLevel, os.Stdout)
	})

	Context("test simple log", func() {
		It("should print log correctly", func() {
			//{"level":"info","ts":1677735613.411057,"msg":"info test","statusCode":"200","url":"http://www.baidu.com","msgCode":100010}
			l.Info("info test", String("statusCode", "200"), String("url", "http://www.sphereex.com"), Int("msgCode", 100010))
			l.Debug("debug test", String("statusCode", "200"), String("url", "http://www.sphereex.com"), Int("msgCode", 100010))
			l.Debug("debug test", Error(fmt.Errorf("get database err")))
		})
	})

	Context("test muilty log", func() {
		It("should print two log file correctly", func() {
			file1, err := os.OpenFile("./access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			file2, err := os.OpenFile("./error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}

			options := []TeeOption{
				{
					W:        file1,
					logLevel: InfoLevel,
				},
				{
					W:        file2,
					logLevel: ErrorLevel,
				},
			}

			l, _ := NewLoggerWithTee(options)

			l.Info("info test", String("statusCode", "200"), String("url", "http://www.sphereex.com"), Int("msgCode", 100010))
			l.Error("error test", String("app", "crash"), Error(fmt.Errorf("get database err")))
		})

		It("should print stdout and log file correctly", func() {
			file, err := os.OpenFile("./tee-error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}

			options := []TeeOption{
				{
					W:        os.Stdout,
					logLevel: InfoLevel,
				},
				{
					W:        file,
					logLevel: ErrorLevel,
				},
			}

			l, _ := NewLoggerWithTee(options)

			l.Info("info test", String("app", "start ok"), Int("major version", 3))
			l.Error("debug test", String("app", "crash"), Int("reason", -1))
		})
	})
})
