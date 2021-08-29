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

// CheckWarning check if there are warning messages
func (t Team) CheckWarnings() bool {
	return len(t.Warnings) > 0
}

const (
	// subURLTeam is a sub url path for team
	subURLTeam = "/team"

	// teamFieldName is a field for team name
	teamFieldName = "name"

	// teamFieldMemberEmailAddress is a field for user member email address in a team
	teamFieldMemberEmailAddress = "email_address"

	// teamFieldMemberAccountID is a field for user member account id in a team
	teamFieldMemberAccountID = "account_id"

	//teamFieldNewOwnerEmailAddress is a field for the email address of an Account on this Team to
	// receive all documents, templates, and API apps (if applicable) from the removed Account
	teamFieldNewOwnerEmailAddress = "new_owner_email_address"
)

var (
	// subURLTeamCreate is a sub url path for create a new team
	subURLTeamCreate = subURLTeam + "/create"

	// subURLTeamDelete is a sub url path for delete a team
	subURLTeamDelete = subURLTeam + "/destroy"

	// subURLTeamAddMember is a sub url path for add a new member to a team
	subURLTeamAddMember = subURLTeam + "/add_member"

	// subURLTeamRemoveMember is a sub url path for remove a member in a team
	subURLTeamRemoveMember = subURLTeam + "/remove_member"
)

// Get returns information about your Team as well as a list of its members.
// Ref: https://app.hellosign.com/api/reference#get_team
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

// Create makes a new team and makes your API account a member
// Ref: https://app.hellosign.com/api/reference#create_team
func (t *TeamAPI) Create(ctx context.Context, teamName string) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	err := writer.WriteField(teamFieldName, teamName)
	if err != nil {
		return Team{}, err
	}

	err = writer.Close()
	if err != nil {
		return Team{}, err
	}

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

// Update the name of your team.
// Ref: https://app.hellosign.com/api/reference#update_team
func (t *TeamAPI) Update(ctx context.Context, teamName string) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	err := writer.WriteField(teamFieldName, teamName)
	if err != nil {
		return Team{}, err
	}

	err = writer.Close()
	if err != nil {
		return Team{}, err
	}

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

// Delete will deletes your Team.
// Can only be invoked when you have a Team with only one member (yourself).
// Ref: https://app.hellosign.com/api/reference#delete_team
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

// TeamAddMemberParam is request param for add a member to a team.
// You can send AccountID or EmailAddress
// If both AccountID and EmailAddress are provided, HelloSign will use AccountID.
type TeamAddMemberParam struct {
	AccountID    string
	EmailAddress string
}

// AddMember will invite a user to your team.
// If the user does not currently have a HelloSign Account, a new one will be created for them.
// If a user is already a part of another Team, a "team_invite_failed" error will be returned.
// Ref: https://app.hellosign.com/api/reference#add_user_to_team
func (t *TeamAPI) AddMember(ctx context.Context, param TeamAddMemberParam) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	err := writer.WriteField(teamFieldMemberAccountID, param.AccountID)
	if err != nil {
		return Team{}, err
	}

	err = writer.WriteField(teamFieldMemberEmailAddress, param.EmailAddress)
	if err != nil {
		return Team{}, err
	}

	err = writer.Close()
	if err != nil {
		return Team{}, err
	}

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

// TeamRemoveMemberParam is request param for remove a member in a team.
// You can send AccountID or EmailAddress
// If both EmailAddress and AccountId are provided, HelloSign will use AccountID.
// NewOwnerEmailAddress is an optinal parameter (available only for Enterprise plans),
// use to receive all documents, templates, and API from the removed account.
type TeamRemoveMemberParam struct {
	AccountID            string
	EmailAddress         string
	NewOwnerEmailAddress string
}

// RemovesMember will remove the provided user Account from your Team
// Ref: https://app.hellosign.com/api/reference#remove_user_from_team
func (t *TeamAPI) RemoveMember(ctx context.Context, param TeamRemoveMemberParam) (Team, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	err := writer.WriteField(teamFieldMemberAccountID, param.AccountID)
	if err != nil {
		return Team{}, err
	}

	err = writer.WriteField(teamFieldMemberEmailAddress, param.EmailAddress)
	if err != nil {
		return Team{}, err
	}

	err = writer.WriteField(teamFieldNewOwnerEmailAddress, param.NewOwnerEmailAddress)
	if err != nil {
		return Team{}, err
	}

	err = writer.Close()
	if err != nil {
		return Team{}, err
	}

	resp, err := t.client.callAPI(
		ctx,
		requestParam{
			path:   t.client.BaseURL + subURLTeamRemoveMember,
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
