package broker

import (
	"fmt"

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


func New() Broker {
	return &simpleBroker{}
}

//TODO Implement stubs later when cli is finished
func (b *simpleBroker) StartInstance() (string, error) {
	fmt.Println("[broker] StartInstance called")
	return "inst-1", nil
}

func (b *simpleBroker) StopInstance(id string) error {
	return nil
}

func (b *simpleBroker) Pause(id string) error {
	return nil
}

func (b *simpleBroker) Resume(id string) error {
	return nil
}

func (b *simpleBroker) AddLog(instanceID, url string) error {
	return nil
}

func (b *simpleBroker) RemoveLog(instanceID, url string) error {
	return nil
}

func (b *simpleBroker) ListInstances() ([]*instance.Instance, error) {
	return nil, nil
}

func (b *simpleBroker) GetInstance(id string) (*instance.Instance, error) {
	return nil, nil
}