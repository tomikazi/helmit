// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
)

type DaemonSetsClient interface {
	DaemonSets() DaemonSetsReader
}

func NewDaemonSetsClient(resources resource.Client, filter resource.Filter) DaemonSetsClient {
	return &daemonSetsClient{
		Client: resources,
		filter: filter,
	}
}

type daemonSetsClient struct {
	resource.Client
	filter resource.Filter
}

func (c *daemonSetsClient) DaemonSets() DaemonSetsReader {
	return NewDaemonSetsReader(c.Client, c.filter)
}
