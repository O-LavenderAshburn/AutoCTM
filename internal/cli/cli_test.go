package cli

import (
	"fmt"
	"testing"

	"sorcerer.nz/autoctm/internal/instance"
)

func TestStart_Success(t *testing.T) {
	mock := &mockBroker{startID: "inst-1"}
	c := New(mock)

	err := c.Start()

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if !mock.startCalled {
		t.Error("expected StartInstance to be called")
	}
}

func TestStart_BrokerError(t *testing.T) {
	mock := &mockBroker{startErr: fmt.Errorf("broker failed")}
	c := New(mock)

	err := c.Start()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestListInstances_Success(t *testing.T) {
	mock := &mockBroker{
		listInstances: []*instance.Instance{
			{ID: "inst-1"},
		},
	}
	c := New(mock)

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
	c := New(mock)

	err := c.ListInstances()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}