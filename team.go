package hellosign

import (
	"context"
	"encoding/json"
	"net/http"
)

// TeamAPI is a service to team API
type TeamAPI service

// Team represent information about your team and its members
type Team struct {
	Name            string    `json:"name"`
	Accounts        []Account `json:"accounts"`
	InvitedAccounts []Account `json:"invited_accounts"`
}

const (
	// subURLTeam is a sub url path for team
	subURLTeam = "/team"
)

func (t *TeamAPI) Get(ctx context.Context) (Team, error) {
	resp, err := t.client.callAPI(
		ctx,
		requestParam{
			path:   t.client.BaseURL + subURLTeam,
			method: http.MethodGet,
		},
	)

	if err != nil {
		return Team{}, err
	}

	team := Team{}
	err = json.NewDecoder(resp.Body).Decode(&team)
	if err != nil {
		return Team{}, err
	}

	return team, nil
}
