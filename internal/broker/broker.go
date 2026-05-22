package broker

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/google/uuid"
	"sorcerer.nz/autoctm/internal/instance"
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

type managedInstance struct {
	id      string
	process *os.Process
	conn    net.Conn
	status  string
}

func New() Broker {
	return &instanceBroker{}
}

func (b *instanceBroker) StartInstance() (string, error) {
	// Generate a UUID for the instance (will become the DB primary key later)
	id := uuid.New().String()
	fmt.Printf("Recv command")

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
