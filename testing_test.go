package golden

import (
	"fmt"
	"regexp"
	"strings"
)

var _ TestingTB = new(bufferTB)

type bufferTB struct {
	logs []string
	name string
}

func (m *bufferTB) Bytes() []byte {
	return []byte(strings.Join(m.logs, "\n"))
}

func (m *bufferTB) Errorf(format string, args ...interface{}) {
	m.Logf(format, args...)
	m.Fail()
}

func (m *bufferTB) Fail() {
	m.Logf("golden_test: method called %T.Fail()", m)
}

func (m *bufferTB) FailNow() {
	msg := fmt.Sprintf("golden_test: method called %T.FailNow()", m)
	m.Logf(msg)
	panic(msg)
}

func (m *bufferTB) Fatalf(format string, args ...interface{}) {
	m.Logf(format, args...)
	m.FailNow()
}

func (m *bufferTB) Helper() {
	m.Logf("golden_test: method called %T.Helper()", m)
}

func (m *bufferTB) Logf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	re := regexp.MustCompile(`(?im)Error Trace:([\w\s:.]+)Error:`)
	msg = re.ReplaceAllString(msg, "Error Trace:\n\tError:")
	m.logs = append(m.logs, msg)
}

func (m bufferTB) Name() string {
	return m.name
}
