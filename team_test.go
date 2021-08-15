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
