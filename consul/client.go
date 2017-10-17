package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Client struct {
	consul *api.Client
}

func NewClient(addr string) (*Client, error) {
	var config = api.DefaultConfig()
	config.Address = addr

	var client, err = api.NewClient(config)

	if err != nil {
		fmt.Errorf("unable to init Consul client: %v", err)
		return nil, err
	}

	return &Client{consul: client}, nil
}

func (c *Client) RegisterService(port int, id string, name string, tags []string) error {
	return c.consul.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:   id,
		Name: name,
		Port: port,
		Tags: tags,
	})
}

func (c *Client) UnRegisterService(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

func (c *Client) GetAgents() {
	if events, _, err := c.consul.Event().List("test", nil); err == nil {
		for _, event := range events {
			fmt.Println(event)
		}
	}

	if services, err := c.consul.Agent().Services(); err == nil {
		for id, service := range services {
			fmt.Println(fmt.Sprintf("%s: %v", id, service))
		}
	}
}
