package cli

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	"github.com/tierklinik-dobersberg/apis/pkg/data"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (root *Root) ResolveUserIds(ctx context.Context, idOrNames []string) ([]string, error) {
	allUsers, err := root.Users().
		ListUsers(ctx, connect.NewRequest(&idmv1.ListUsersRequest{
			FieldMask: &fieldmaskpb.FieldMask{
				Paths: []string{"users.user.id", "users.user.username"},
			},
		}))

	if err != nil {
		return nil, err
	}

	idByName := data.IndexSlice(allUsers.Msg.Users, func(u *idmv1.Profile) string {
		return u.User.Username
	})

	userIds := data.IndexSlice(allUsers.Msg.Users, func(u *idmv1.Profile) string {
		return u.User.Id
	})

	resultMap := make(map[string]struct{}, len(idOrNames))
	for _, idOrName := range idOrNames {
		if _, isId := userIds[idOrName]; isId {
			resultMap[idOrName] = struct{}{}

			continue
		}

		profile, isName := idByName[idOrName]
		if isName {
			resultMap[profile.User.Id] = struct{}{}

			continue
		}

		return nil, fmt.Errorf("user id or name %q not found", idOrName)
	}

	return maps.Keys(resultMap), nil
}
