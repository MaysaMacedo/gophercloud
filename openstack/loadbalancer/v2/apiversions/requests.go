package apiversions

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// lists all the load balancer API versions available to end-users.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, apiVersionsURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}
