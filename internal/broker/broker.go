package broker

import "sorcerer.nz/autoctm/internal/instance"

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