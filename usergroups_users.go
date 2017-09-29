package slack

// Auto-generated by internal/cmd/genmethods/genmethods.go. DO NOT EDIT!

import (
	"context"
	"net/url"
	"strconv"

	"github.com/lestrrat/go-slack/objects"
	"github.com/pkg/errors"
)

var _ = strconv.Itoa
var _ = objects.EpochTime(0)

// UsergroupsUsersListCall is created by UsergroupsUsersService.List method call
type UsergroupsUsersListCall struct {
	service         *UsergroupsUsersService
	includeDisabled bool
	usergroup       string
}

// UsergroupsUsersUpdateCall is created by UsergroupsUsersService.Update method call
type UsergroupsUsersUpdateCall struct {
	service      *UsergroupsUsersService
	includeCount bool
	usergroup    string
	users        string // Comma-separated list of user IDs
}

// List creates a UsergroupsUsersListCall object in preparation for accessing the usergroups.users.list endpoint
func (s *UsergroupsUsersService) List(usergroup string) *UsergroupsUsersListCall {
	var call UsergroupsUsersListCall
	call.service = s
	call.usergroup = usergroup
	return &call
}

// IncludeDisabled sets the value for optional includeDisabled parameter
func (c *UsergroupsUsersListCall) IncludeDisabled(includeDisabled bool) *UsergroupsUsersListCall {
	c.includeDisabled = includeDisabled
	return c
}

// Values returns the UsergroupsUsersListCall object as url.Values
func (c *UsergroupsUsersListCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeDisabled {
		v.Set("include_disabled", "true")
	}

	if len(c.usergroup) <= 0 {
		return nil, errors.New(`missing required parameter usergroup`)
	}
	v.Set("usergroup", c.usergroup)
	return v, nil
}

// Do executes the call to access usergroups.users.list endpoint
func (c *UsergroupsUsersListCall) Do(ctx context.Context) (objects.UsergroupUsersList, error) {
	const endpoint = "usergroups.users.list"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		objects.UsergroupUsersList `json:"users"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.users.list`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.UsergroupUsersList, nil
}

// Update creates a UsergroupsUsersUpdateCall object in preparation for accessing the usergroups.users.update endpoint
func (s *UsergroupsUsersService) Update(usergroup string, users string) *UsergroupsUsersUpdateCall {
	var call UsergroupsUsersUpdateCall
	call.service = s
	call.usergroup = usergroup
	call.users = users
	return &call
}

// IncludeCount sets the value for optional includeCount parameter
func (c *UsergroupsUsersUpdateCall) IncludeCount(includeCount bool) *UsergroupsUsersUpdateCall {
	c.includeCount = includeCount
	return c
}

// Values returns the UsergroupsUsersUpdateCall object as url.Values
func (c *UsergroupsUsersUpdateCall) Values() (url.Values, error) {
	v := url.Values{}
	v.Set(`token`, c.service.token)

	if c.includeCount {
		v.Set("include_count", "true")
	}

	if len(c.usergroup) <= 0 {
		return nil, errors.New(`missing required parameter usergroup`)
	}
	v.Set("usergroup", c.usergroup)

	if len(c.users) <= 0 {
		return nil, errors.New(`missing required parameter users`)
	}
	v.Set("users", c.users)
	return v, nil
}

// Do executes the call to access usergroups.users.update endpoint
func (c *UsergroupsUsersUpdateCall) Do(ctx context.Context) (*objects.Usergroup, error) {
	const endpoint = "usergroups.users.update"
	v, err := c.Values()
	if err != nil {
		return nil, err
	}
	var res struct {
		SlackResponse
		*objects.Usergroup `json:"usergroup"`
	}
	if err := c.service.client.postForm(ctx, endpoint, v, &res); err != nil {
		return nil, errors.Wrap(err, `failed to post to usergroups.users.update`)
	}
	if !res.OK {
		return nil, errors.New(res.Error.String())
	}

	return res.Usergroup, nil
}
