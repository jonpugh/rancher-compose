package client

const (
	DNS_SERVICE_TYPE = "dnsService"
)

type DnsService struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	EnvironmentId string `json:"environmentId,omitempty" yaml:"environment_id,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type DnsServiceCollection struct {
	Collection
	Data []DnsService `json:"data,omitempty"`
}

type DnsServiceClient struct {
	rancherClient *RancherClient
}

type DnsServiceOperations interface {
	List(opts *ListOpts) (*DnsServiceCollection, error)
	Create(opts *DnsService) (*DnsService, error)
	Update(existing *DnsService, updates interface{}) (*DnsService, error)
	ById(id string) (*DnsService, error)
	Delete(container *DnsService) error

	ActionActivate(*DnsService) (*Service, error)

	ActionAddservicelink(*DnsService, *AddRemoveServiceLinkInput) (*Service, error)

	ActionCreate(*DnsService) (*Service, error)

	ActionDeactivate(*DnsService) (*Service, error)

	ActionRemove(*DnsService) (*Service, error)

	ActionRemoveservicelink(*DnsService, *AddRemoveServiceLinkInput) (*Service, error)

	ActionSetservicelinks(*DnsService, *SetServiceLinksInput) (*Service, error)

	ActionUpdate(*DnsService) (*Service, error)
}

func newDnsServiceClient(rancherClient *RancherClient) *DnsServiceClient {
	return &DnsServiceClient{
		rancherClient: rancherClient,
	}
}

func (c *DnsServiceClient) Create(container *DnsService) (*DnsService, error) {
	resp := &DnsService{}
	err := c.rancherClient.doCreate(DNS_SERVICE_TYPE, container, resp)
	return resp, err
}

func (c *DnsServiceClient) Update(existing *DnsService, updates interface{}) (*DnsService, error) {
	resp := &DnsService{}
	err := c.rancherClient.doUpdate(DNS_SERVICE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DnsServiceClient) List(opts *ListOpts) (*DnsServiceCollection, error) {
	resp := &DnsServiceCollection{}
	err := c.rancherClient.doList(DNS_SERVICE_TYPE, opts, resp)
	return resp, err
}

func (c *DnsServiceClient) ById(id string) (*DnsService, error) {
	resp := &DnsService{}
	err := c.rancherClient.doById(DNS_SERVICE_TYPE, id, resp)
	return resp, err
}

func (c *DnsServiceClient) Delete(container *DnsService) error {
	return c.rancherClient.doResourceDelete(DNS_SERVICE_TYPE, &container.Resource)
}

func (c *DnsServiceClient) ActionActivate(resource *DnsService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *DnsServiceClient) ActionAddservicelink(resource *DnsService, input *AddRemoveServiceLinkInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "addservicelink", &resource.Resource, input, resp)

	return resp, err
}

func (c *DnsServiceClient) ActionCreate(resource *DnsService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *DnsServiceClient) ActionDeactivate(resource *DnsService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *DnsServiceClient) ActionRemove(resource *DnsService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *DnsServiceClient) ActionRemoveservicelink(resource *DnsService, input *AddRemoveServiceLinkInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "removeservicelink", &resource.Resource, input, resp)

	return resp, err
}

func (c *DnsServiceClient) ActionSetservicelinks(resource *DnsService, input *SetServiceLinksInput) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "setservicelinks", &resource.Resource, input, resp)

	return resp, err
}

func (c *DnsServiceClient) ActionUpdate(resource *DnsService) (*Service, error) {

	resp := &Service{}

	err := c.rancherClient.doAction(DNS_SERVICE_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
