// Code generated by "make api"; DO NOT EDIT.
package scopes

import (
	"context"
	"fmt"

	"github.com/hashicorp/watchtower/api"
)

// DeleteProject returns true iff the project existed when the delete attempt was made.
func (s Organization) DeleteProject(ctx context.Context, project *Project) (bool, *api.Error, error) {
	if s.Client == nil {
		return false, nil, fmt.Errorf("nil client in DeleteProject request")
	}
	if s.Id == "" {

		// Assume the client has been configured with organization already and
		// move on

	} else {
		// If it's explicitly set here, override anything that might be in the
		// client

		ctx = context.WithValue(ctx, "org", s.Id)

	}
	if project.Id == "" {
		return false, nil, fmt.Errorf("empty project ID field in DeleteProject request")
	}

	req, err := s.Client.NewRequest(ctx, "DELETE", fmt.Sprintf("projects/%s", project.Id), nil)
	if err != nil {
		return false, nil, fmt.Errorf("error creating DeleteProject request: %w", err)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return false, nil, fmt.Errorf("error performing client request during DeleteProject call: %w", err)
	}

	type deleteResponse struct {
		Existed bool
	}
	target := &deleteResponse{}

	apiErr, err := resp.Decode(target)
	if err != nil {
		return false, nil, fmt.Errorf("error decoding DeleteProject repsonse: %w", err)
	}

	return target.Existed, apiErr, nil
}
