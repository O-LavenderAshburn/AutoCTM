package cli_test

import (
	"fmt"
	"testing"

	"sorcerer.nz/autoctm/internal/cli"
	"sorcerer.nz/autoctm/internal/instance"
)

// test start
func TestStart_Success(t *testing.T) {
	mock := &mockBroker{startID: "inst-1"}
	c := cli.New(mock)

	err := c.Start()

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if !mock.startCalled {
		t.Error("expected StartInstance to be called")
	}
}

// BrokerError cases verify that CLI surfaces errors rather than swallowing them
func TestStart_BrokerError(t *testing.T) {
	mock := &mockBroker{startErr: fmt.Errorf("broker failed")}
	c := cli.New(mock)

	err := c.Start()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}


// ListInstances tests verify that CLI.ListInstances correctly delegates
// to the broker and surfaces any errors returned.
func TestListInstances_Success(t *testing.T) {
	mock := &mockBroker{
		listInstances: []*instance.Instance{
			{ID: "inst-1"},
		},
	}
	c := cli.New(mock)

	err := c.ListInstances()

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if !mock.listCalled {
		t.Error("expected ListInstances to be called")
	}
}

func TestListInstances_BrokerError(t *testing.T) {
	mock := &mockBroker{listErr: fmt.Errorf("broker failed")}
	c := cli.New(mock)

	err := c.ListInstances()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}