package privacy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	"google.golang.org/protobuf/proto"
)

func TestFilterAllowedFields(t *testing.T) {
	cases := []struct {
		User           string
		Roles          []string
		Input          proto.Message
		ExpectedOutput proto.Message
	}{
		{
			User: "test",
			Input: &idmv1.User{
				Id:          "user",
				DisplayName: "display-name",
				FirstName:   "first-name",
			},
			ExpectedOutput: &idmv1.User{
				Id:          "user",
				DisplayName: "display-name",
			},
		},
		{
			User: "user",
			Input: &idmv1.User{
				Id:          "user",
				DisplayName: "display-name",
				FirstName:   "first-name",
			},
			ExpectedOutput: &idmv1.User{
				Id:          "user",
				DisplayName: "display-name",
				FirstName:   "first-name",
			},
		},
		{
			User: "user",
			Input: &idmv1.Profile{
				User: &idmv1.User{
					Id:          "other",
					DisplayName: "display-name",
					FirstName:   "first-name",
				},
				EmailAddresses: []*idmv1.EMail{
					{
						Address:  "test@example.com",
						Verified: true,
					},
				},
			},
			ExpectedOutput: &idmv1.Profile{
				User: &idmv1.User{
					Id:          "other",
					DisplayName: "display-name",
				},
			},
		},
		{
			User: "user",
			Input: &idmv1.Profile{
				User: &idmv1.User{
					Id:          "user",
					DisplayName: "display-name",
					FirstName:   "first-name",
				},
			},
			ExpectedOutput: &idmv1.Profile{
				User: &idmv1.User{
					Id:          "user",
					DisplayName: "display-name",
					FirstName:   "first-name",
				},
			},
		},
	}

	for idx, testCase := range cases {
		t.Run(fmt.Sprintf("test-case-%d", idx), func(t *testing.T) {
			assert.NoError(t, FilterAllowedFields(testCase.Input, testCase.User, testCase.Roles))
			assert.True(t, proto.Equal(testCase.Input, testCase.ExpectedOutput))
		})
	}
}
