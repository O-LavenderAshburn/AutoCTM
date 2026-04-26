package cli

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"sorcerer.nz/autoctm/internal/instance"
)

func captureOutput(fn func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func newTestRunner(m *mockBroker) *CLIRunner {
	return NewRunner(New(m))
}

func TestHandle_UnknownCommand(t *testing.T) {
	mock := &mockBroker{}
	r := newTestRunner(mock)

	out := captureOutput(func() { r.handle("/unknown") })

	if !strings.Contains(out, "unknown command") {
		t.Errorf("expected 'unknown command' in output, got: %q", out)
	}
}

func TestHandle_Start_Success(t *testing.T) {
	mock := &mockBroker{startID: "inst-1"}
	r := newTestRunner(mock)

	captureOutput(func() { r.handle("/start") })

	if !mock.startCalled {
		t.Error("expected StartInstance to be called")
	}
}

func TestHandle_Start_BrokerError(t *testing.T) {
	mock := &mockBroker{startErr: fmt.Errorf("broker failed")}
	r := newTestRunner(mock)

	out := captureOutput(func() { r.handle("/start") })

	if !strings.Contains(out, "error:") {
		t.Errorf("expected error output, got: %q", out)
	}
}

func TestHandle_List_Success(t *testing.T) {
	mock := &mockBroker{
		listInstances: []*instance.Instance{
			{ID: "inst-1", Status: "running"},
			{ID: "inst-2", Status: "stopped"},
		},
	}
	r := newTestRunner(mock)

	out := captureOutput(func() { r.handle("/list") })

	if !mock.listCalled {
		t.Error("expected ListInstances to be called")
	}
	for _, want := range []string{"inst-1", "inst-2", "running", "stopped"} {
		if !strings.Contains(out, want) {
			t.Errorf("expected %q in output, got: %q", want, out)
		}
	}
}

func TestHandle_List_BrokerError(t *testing.T) {
	mock := &mockBroker{listErr: fmt.Errorf("db down")}
	r := newTestRunner(mock)

	out := captureOutput(func() { r.handle("/list") })

	if !strings.Contains(out, "error:") {
		t.Errorf("expected error output, got: %q", out)
	}
}

func TestHandle_Quit(t *testing.T) {
	r := newTestRunner(&mockBroker{})

	out := captureOutput(func() { r.handle("quit") })

	if !strings.Contains(out, "bye") {
		t.Errorf("expected 'bye', got: %q", out)
	}
}

func TestHandle_EmptyInput(t *testing.T) {
	r := newTestRunner(&mockBroker{})
	captureOutput(func() { r.handle("") })
}