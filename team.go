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

// Team represent information about your team
type Team struct {
	Team     TeamDetail `json:"team"`
	Warnings []Warnings `json:"warnings,omitempty"`
}

// TeamDetail represent detail information about your team and its members
type TeamDetail struct {
	Name            string    `json:"name"`
	Accounts        []Account `json:"accounts"`
	InvitedAccounts []Account `json:"invited_accounts"`
}

const (
	// subURLTeam is a sub url path for team
	subURLTeam             = "/team"
	subURLTeamCreate       = "/team/create"
	subURLTeamDelete       = "/team/destroy"
	subURLTeamAddMember    = "/team/add_member"
	subURLTeamRemoveMember = "/team/remove_member"

	// teamFieldName is a field for team name
	teamFieldName = "name"

	// teamFieldMemberEmailAddress is a field for manage team member
	teamFieldMemberEmailAddress = "email_address"

	// teamFieldMemberAccountID is a field for user member accountID in a team
	teamFieldMemberAccountID = "account_id"
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

	err := writer.WriteField(teamFieldName, teamName)
	if err != nil {
		return Team{}, err
	}
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

	err := writer.WriteField(teamFieldName, teamName)
	if err != nil {
		return Team{}, err
	}
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

func (t *TeamAPI) Delete(ctx context.Context) error {
	_, err := t.client.callAPI(
		ctx,
		requestParam{
			path:   t.client.BaseURL + subURLTeamDelete,
			method: http.MethodPost,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (t *TeamAPI) AddMember(ctx context.Context, emailAddress, accountID string) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	err := writer.WriteField(teamFieldMemberEmailAddress, emailAddress)
	if err != nil {
		return Team{}, err
	}

	err = writer.WriteField(teamFieldMemberAccountID, accountID)
	if err != nil {
		return Team{}, err
	}
	writer.Close()

	resp, err := t.client.callAPI(
		ctx,
		requestParam{
			path:   t.client.BaseURL + subURLTeamAddMember,
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

func (t *TeamAPI) RemoveMember(ctx context.Context, emailAddress, accountID string) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	emailAddressField, err := writer.CreateFormField(teamFieldMemberEmailAddress)
	if err != nil {
		return Team{}, err
	}
	_, err = emailAddressField.Write([]byte(emailAddress))
	if err != nil {
		return Team{}, err
	}

	accountIDField, err := writer.CreateFormField(teamFieldMemberAccountID)
	if err != nil {
		return Team{}, err
	}
	_, err = accountIDField.Write([]byte(accountID))
	if err != nil {
		return Team{}, err
	}

	writer.Close()

	resp, err := t.client.callAPI(
		ctx,
		requestParam{
			path:   t.client.BaseURL + subURLTeamAddMember,
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
