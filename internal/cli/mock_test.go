package cli_test

import "sorcerer.nz/autoctm/internal/instance"

type mockBroker struct {
	startCalled bool
	startID     string
	startErr    error

	stopCalled bool
	stopErr    error

	pauseCalled bool
	pauseErr    error

	resumeCalled bool
	resumeErr    error

	addLogCalled bool
	addLogErr    error

	removeLogCalled bool
	removeLogErr    error

	listCalled    bool
	listInstances []*instance.Instance
	listErr       error

	getCalled   bool
	getInstance *instance.Instance
	getErr      error
}

func (m *mockBroker) StartInstance() (string, error) {
	m.startCalled = true
	return m.startID, m.startErr
}

func (m *mockBroker) StopInstance(id string) error {
	m.stopCalled = true
	return m.stopErr
}

func (m *mockBroker) Pause(id string) error {
	m.pauseCalled = true
	return m.pauseErr
}

func (m *mockBroker) Resume(id string) error {
	m.resumeCalled = true
	return m.resumeErr
}

func (m *mockBroker) AddLog(id, url string) error {
	m.addLogCalled = true
	return m.addLogErr
}

func (m *mockBroker) RemoveLog(id, url string) error {
	m.removeLogCalled = true
	return m.removeLogErr
}

func (m *mockBroker) ListInstances() ([]*instance.Instance, error) {
	m.listCalled = true
	return m.listInstances, m.listErr
}

func (m *mockBroker) GetInstance(id string) (*instance.Instance, error) {
	m.getCalled = true
	return m.getInstance, m.getErr
}