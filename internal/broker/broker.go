package broker

import "sorcerer.nz/internal/instance"

type Broker interface {
	StartInstance() (string, error)
	StopInstance(id string) error
	ListInstances() ([]*instance.Instance, error)
	GetInstance(id string) (*instance.Instance, error)
}