package cli

import "sorcerer.nz/internal/broker"

type CLI struct {
	store   InstanceStore
	broker  Broker
	context *InstanceContext
}

// Active instance snapshot used by CLI routing
type InstanceContext struct {
	ID        string
	Status    string
	StartedAt time.Time
	Active    bool
}

type InstanceStore interface {
	GetByID(id string) (*instance.Instance, error)
	List() ([]*instance.Instance, error)
}

type Broker interface {
	StartInstance() (string, error)
	AddLog(instanceID, url string) error
	RemoveLog(instanceID, url string) error
	Pause(instanceID string) error
	Resume(instanceID string) error
	Status(instanceID string) (*instance.Instance, error)
	Shutdown(instanceID string) error
}


func New(store InstanceStore, broker Broker) *CLI {
	return &CLI{
		store:  store,
		broker: broker,
	}
}


