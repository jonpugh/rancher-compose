package project

import "github.com/rancherio/go-rancher/client"

type ServiceConfig struct {
	CapAdd      []string `yaml:"cap_add,omitempty"`
	CapDrop     []string `yaml:"cap_drop,omitempty"`
	CpuShares   int64    `yaml:"cpu_shares,omitempty"`
	Command     string   `yaml:"command,omitempty"`
	Detach      string   `yaml:"detach,omitempty"`
	Dns         []string `yaml:"dns,omitempty"`
	DnsSearch   string   `yaml:"dns_search,omitempty"`
	DomainName  string   `yaml:"domainname,omitempty"`
	Entrypoint  string   `yaml:"entrypoint,omitempty"`
	EnvFile     string   `yaml:"env_file,omitempty"`
	Environment []string `yaml:"environment,omitempty"`
	Hostname    string   `yaml:"hostname,omitempty"`
	Image       string   `yaml:"image,omitempty"`
	Labels      []string `yaml:"labels,omitempty"`
	Links       []string `yaml:"links,omitempty"`
	MemLimit    int64    `yaml:"mem_limit,omitempty"`
	Name        string   `yaml:"name,omitempty"`
	Net         string   `yaml:"net,omitempty"`
	Pid         string   `yaml:"pid,omitempty"`
	Ipc         string   `yaml:"ipc,omitempty"`
	Ports       []string `yaml:"ports,omitempty"`
	Privileged  bool     `yaml:"privileged,omitempty"`
	Restart     string   `yaml:"restart,omitempty"`
	ReadOnly    bool     `yaml:"read_only,omitempty"`
	StdinOpen   bool     `yaml:"stdin_open,omitempty"`
	Tty         bool     `yaml:"tty,omitempty"`
	User        string   `yaml:"user,omitempty"`
	Volumes     []string `yaml:"volumes,omitempty"`
	VolumesFrom []string `yaml:"volumes_from,omitempty"`
	WorkingDir  string   `yaml:"working_dir,omitempty"`
	//`yaml:"build,omitempty"`
	Expose        []string `yaml:"expose,omitempty"`
	ExternalLinks []string `yaml:"external_links,omitempty"`
}

type Project struct {
	Name     string
	configs  map[string]*ServiceConfig
	Services map[string]Service
	file     string
	content  []byte
	client   *client.RancherClient
	factory  ServiceFactory
}

type Service interface {
	Up() error
	Config() *ServiceConfig
}

type ServiceFactory interface {
	Create(project *Project, name string, serviceConfig *ServiceConfig) (Service, error)
}