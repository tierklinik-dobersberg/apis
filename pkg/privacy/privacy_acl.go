package privacy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/mennanov/fmutils"
	commonv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func FilterAllowedFields(msg proto.Message, currentUser string, currentRoles []string) error {
	return filterAllowedFields(msg, currentUser, currentRoles, false)
}

func filterAllowedFields(msg proto.Message, currentUser string, currentRoles []string, isAlreadyAllowed bool) error {
	hasPermission, msgOpts := hasPermission(currentUser, currentRoles, msg)

	// if there is no constraint on the allowed user/roles
	// we fall back to whether or not access has been allowed
	// by a parent message.
	if !hasPermission && msgOpts.OwnerFieldName == "" && len(msgOpts.AllowedRoles) == 0 {
		hasPermission = isAlreadyAllowed
	}

	if !hasPermission {
		mask := fmutils.NestedMaskFromPaths(msgOpts.FieldMask.Paths)
		mask.Filter(msg)
	}

	rtf := msg.ProtoReflect()
	var errs = new(multierror.Error)

	// recursively check sub-messages for PrivacyACL options and apply them as well.
	rtf.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if !v.IsValid() {
			return true
		}

		if fd.Kind() == protoreflect.MessageKind {
			if fd.IsList() {
				list := v.List()
				for i := 0; i < list.Len(); i++ {
					if err := filterAllowedFields(list.Get(i).Message().Interface(), currentUser, currentRoles, hasPermission); err != nil {
						errs.Errors = append(errs.Errors, err)
					}
				}

				return true
			}

			if err := filterAllowedFields(v.Message().Interface(), currentUser, currentRoles, hasPermission); err != nil {
				errs.Errors = append(errs.Errors, err)
			}
		}

		return true
	})

	return errs.ErrorOrNil()
}

func hasPermission(currentUser string, currentRoles []string, msg proto.Message) (bool, *commonv1.PrivacyACL) {
	opts := msg.ProtoReflect().Descriptor().Options()
	ext := proto.GetExtension(opts, commonv1.E_Readable)

	msgOpts, ok := ext.(*commonv1.PrivacyACL)
	if !ok || msgOpts == nil {
		return true, nil
	}

	if msgOpts.OwnerFieldName != "" {
		ownerValue, err := getValueFromPath(msg, strings.Split(msgOpts.OwnerFieldName, "."))
		if err != nil {
			panic(err)
		}

		if ownerValue.IsValid() && ownerValue.String() == currentUser {
			return true, msgOpts
		}
	}

	if len(msgOpts.AllowedRoles) > 0 {
		for _, role := range msgOpts.AllowedRoles {
			if slices.Contains(currentRoles, role) {
				return true, msgOpts
			}
		}
	}

	return false, msgOpts
}

func getValueFromPath(msg proto.Message, path []string) (protoreflect.Value, error) {
	rtm := msg.ProtoReflect()
	fieldDesc := rtm.Descriptor().Fields().ByName(protoreflect.Name(path[0]))

	val := rtm.Get(fieldDesc)

	if len(path) == 1 {
		return val, nil
	}

	if fieldDesc.Kind() == protoreflect.MessageKind {
		return getValueFromPath(val.Message().Interface(), path[1:])
	}

	return val, fmt.Errorf("unexpected protobuf field type %s, expected protoreflect.MessageKind", fieldDesc.Kind())
}
