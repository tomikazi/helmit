// Code generated by helmit-generate. DO NOT EDIT.

package v2alpha1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
)

type CronJobsClient interface {
	CronJobs() CronJobsReader
}

func NewCronJobsClient(resources resource.Client, filter resource.Filter) CronJobsClient {
	return &cronJobsClient{
		Client: resources,
		filter: filter,
	}
}

type cronJobsClient struct {
	resource.Client
	filter resource.Filter
}

func (c *cronJobsClient) CronJobs() CronJobsReader {
	return NewCronJobsReader(c.Client, c.filter)
}
