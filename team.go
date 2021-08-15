package hellosign

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
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
	subURLTeam        = "/team"
	subURLTeamCreate  = "/team/create"
	subURLTeamDestroy = "/team/destroy"

	// teamFieldName is a field for team name
	teamFieldName = "name"
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

func (t *TeamAPI) Create(ctx context.Context, teamName string) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	teamNameField, err := writer.CreateFormField(teamFieldName)
	if err != nil {
		return Team{}, err
	}
	teamNameField.Write([]byte(teamName))
	writer.Close()

	resp, err := t.client.callAPI(
		ctx,
		requestParam{
			path:   t.client.BaseURL + subURLTeamCreate,
			method: http.MethodPost,
			writer: writer,
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

func (t *TeamAPI) Update(ctx context.Context, teamName string) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	teamNameField, err := writer.CreateFormField(teamFieldName)
	if err != nil {
		return Team{}, err
	}
	teamNameField.Write([]byte(teamName))
	writer.Close()

	resp, err := t.client.callAPI(
		ctx,
		requestParam{
			path:   t.client.BaseURL + subURLTeam,
			method: http.MethodPost,
			writer: writer,
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
