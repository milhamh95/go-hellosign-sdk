package hellosign_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/matryer/is"

	"github.com/milhamhidayat/go-hellosign-sdk"
	"github.com/milhamhidayat/go-hellosign-sdk/testdata"
)

func TestTeam_Get(t *testing.T) {
	is := is.New(t)

	teamJSON := testdata.GetGolden(t, "team")

	team := hellosign.Team{}
	err := json.Unmarshal(teamJSON, &team)
	is.NoErr(err)

	teamNotFoundJSON := testdata.GetGolden(t, "team-not-found")

	tests := map[string]struct {
		teamHTTPClient *http.Client
		expectedTeam   hellosign.Team
		expectedError  error
	}{
		"success": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusOK, teamJSON, make(http.Header)),
			expectedTeam:   team,
			expectedError:  nil,
		},
		"not found": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusNotFound, teamNotFoundJSON, make(http.Header)),
			expectedTeam:   hellosign.Team{},
			expectedError:  errors.New("not_found: Team does not exist"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)

			apiClient := hellosign.NewClient("123")
			apiClient.HTTPClient = test.teamHTTPClient
			resp, err := apiClient.TeamAPI.Get(context.TODO())
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedTeam, resp)
		})
	}
}

func TestTeam_Create(t *testing.T) {
	is := is.New(t)

	teamJSON := testdata.GetGolden(t, "team")

	team := hellosign.Team{}
	err := json.Unmarshal(teamJSON, &team)
	is.NoErr(err)

	tests := map[string]struct {
		teamHTTPClient *http.Client
		teamName       string
		expectedTeam   hellosign.Team
		expectedError  error
	}{
		"success": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusOK, teamJSON, make(http.Header)),
			teamName:       "Team HelloSign",
			expectedTeam:   team,
			expectedError:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)

			apiClient := hellosign.NewClient("123")
			apiClient.HTTPClient = test.teamHTTPClient
			resp, err := apiClient.TeamAPI.Create(context.TODO(), test.teamName)
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedTeam, resp)
		})
	}
}

func TestTeam_Update(t *testing.T) {
	is := is.New(t)

	teamJSON := testdata.GetGolden(t, "team")

	team := hellosign.Team{}
	err := json.Unmarshal(teamJSON, &team)
	is.NoErr(err)

	tests := map[string]struct {
		teamHTTPClient *http.Client
		teamName       string
		expectedTeam   hellosign.Team
		expectedError  error
	}{
		"success": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusOK, teamJSON, make(http.Header)),
			teamName:       "Team HelloSign",
			expectedTeam:   team,
			expectedError:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)

			apiClient := hellosign.NewClient("123")
			apiClient.HTTPClient = test.teamHTTPClient
			resp, err := apiClient.TeamAPI.Create(context.TODO(), test.teamName)
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedTeam, resp)
		})
	}
}

func TestTeam_Delete(t *testing.T) {
	tests := map[string]struct {
		teamHTTPClient *http.Client
		expectedError  error
	}{
		"success": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusOK, []byte{}, make(http.Header)),
			expectedError:  nil,
		},
	}

	for testname, test := range tests {
		t.Run(testname, func(t *testing.T) {
			is := is.New(t)

			apiClient := hellosign.NewClient("123")
			apiClient.HTTPClient = test.teamHTTPClient
			err := apiClient.TeamAPI.Delete(context.TODO())
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
		})
	}
}

func TestTeam_AddMember(t *testing.T) {
	is := is.New(t)

	teamJSON := testdata.GetGolden(t, "team-add-member")

	team := hellosign.Team{}
	err := json.Unmarshal(teamJSON, &team)
	is.NoErr(err)

	teamNotFoundJSON := testdata.GetGolden(t, "team-not-found")

	tests := map[string]struct {
		teamHTTPClient *http.Client
		expectedTeam   hellosign.Team
		expectedError  error
	}{
		"success": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusOK, teamJSON, make(http.Header)),
			expectedTeam:   team,
			expectedError:  nil,
		},
		"team not found": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusNotFound, teamNotFoundJSON, make(http.Header)),
			expectedTeam:   hellosign.Team{},
			expectedError:  errors.New("not_found: Team does not exist"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)

			apiClient := hellosign.NewClient("123")
			apiClient.HTTPClient = test.teamHTTPClient
			resp, err := apiClient.TeamAPI.AddMember(context.TODO(), hellosign.TeamAddMemberParam{
				EmailAddress: "team_member@hellosign.com",
			})
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedTeam, resp)
		})
	}
}

func TestTeam_RemoveMember(t *testing.T) {
	is := is.New(t)

	teamJSON := testdata.GetGolden(t, "team")

	team := hellosign.Team{}
	err := json.Unmarshal(teamJSON, &team)
	is.NoErr(err)

	tests := map[string]struct {
		teamHTTPClient *http.Client
		expectedTeam   hellosign.Team
		expectedError  error
	}{
		"success": {
			teamHTTPClient: testdata.MockHTTPClient(t, http.StatusOK, teamJSON, make(http.Header)),
			expectedTeam:   team,
			expectedError:  nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)

			apiClient := hellosign.NewClient("123")
			apiClient.HTTPClient = test.teamHTTPClient
			resp, err := apiClient.TeamAPI.RemoveMember(context.TODO(), hellosign.TeamRemoveMemberParam{
				EmailAddress: "team_member1@hellosign.com",
			})
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedTeam, resp)
		})
	}
}
