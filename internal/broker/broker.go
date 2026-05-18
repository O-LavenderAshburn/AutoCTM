package broker

import (
	"fmt"

	"sorcerer.nz/autoctm/internal/instance"
	"github.com/google/uuid"
)

type Broker interface {
	StartInstance() (string, error)
	StopInstance(id string) error
	Pause(id string) error
	Resume(id string) error
	AddLog(instanceID, url string) error
	RemoveLog(instanceID, url string) error
	ListInstances() ([]*instance.Instance, error)
	GetInstance(id string) (*instance.Instance, error)
}

type instanceBroker struct {
    mu        sync.Mutex
    instances map[string]*managedInstance
}

type  managedInstance struct {
	id string
	process *os.process
	conn net.Conn
	status string
}



func New() Broker {
	return &instanceBroker{}
}

func (b *instanceBroker) StartInstance() (string, error) {
    // Generate a UUID for the instance (will become the DB primary key later)
    id := uuid.New().String()

    // Spawn autoctm-instance as a child process, passing only the ID per spec 3.2.3
    cmd := exec.Command("autoctm-instance", id)
    if err := cmd.Start(); err != nil {
        return "", fmt.Errorf("failed to spawn instance: %w", err)
    }

    b.mu.Lock()
    b.instances[id] = &managedInstance{
        id:      id,
        process: cmd.Process,
        status:  "running",
    }
    b.mu.Unlock()

    // Wait for the process to exit in the background.

	// TODO: Implement in IPC
	// Dispatch the Grim Reaper on process to avoid Zombies.
	go func() { cmd.Wait() }()

    fmt.Printf("[broker] started instance %s (pid %d)\n", id, cmd.Process.Pid)
    return id, nil
}

func (b *instanceBroker) StopInstance(id string) error {
	return nil
}

func (b *instanceBroker) Pause(id string) error {
	return nil
}

func (b *instanceBroker) Resume(id string) error {
	return nil
}

func (b *instanceBroker) AddLog(instanceID, url string) error {
	return nil
}

func (b *instanceBroker) RemoveLog(instanceID, url string) error {
	return nil
}

func (b *instanceBroker) ListInstances() ([]*instance.Instance, error) {
	return nil, nil
}

func (b *instanceBroker) GetInstance(id string) (*instance.Instance, error) {
	return nil, nil
}