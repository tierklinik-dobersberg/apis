package cli

import (
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/calendar/v1/calendarv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/comment/v1/commentv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1/rosterv1connect"
)

func (root *Root) OffTime() rosterv1connect.OffTimeServiceClient {
	return rosterv1connect.NewOffTimeServiceClient(root.HttpClient, root.Config().BaseURLS.Roster)
}

func (root *Root) WorkTime() rosterv1connect.WorkTimeServiceClient {
	return rosterv1connect.NewWorkTimeServiceClient(root.HttpClient, root.Config().BaseURLS.Roster)
}

func (root *Root) WorkShift() rosterv1connect.WorkShiftServiceClient {
	return rosterv1connect.NewWorkShiftServiceClient(root.HttpClient, root.Config().BaseURLS.Roster)
}

func (root *Root) Roster() rosterv1connect.RosterServiceClient {
	return rosterv1connect.NewRosterServiceClient(root.HttpClient, root.Config().BaseURLS.Roster)
}

func (root *Root) Constraints() rosterv1connect.ConstraintServiceClient {
	return rosterv1connect.NewConstraintServiceClient(root.HttpClient, root.Config().BaseURLS.Roster)
}

func (root *Root) Users() idmv1connect.UserServiceClient {
	return idmv1connect.NewUserServiceClient(root.HttpClient, root.Config().BaseURLS.Idm)
}

func (root *Root) Roles() idmv1connect.RoleServiceClient {
	return idmv1connect.NewRoleServiceClient(root.HttpClient, root.Config().BaseURLS.Idm)
}

func (root *Root) SelfService() idmv1connect.SelfServiceServiceClient {
	return idmv1connect.NewSelfServiceServiceClient(root.HttpClient, root.Config().BaseURLS.Idm)
}

func (root *Root) Auth() idmv1connect.AuthServiceClient {
	return idmv1connect.NewAuthServiceClient(root.HttpClient, root.Config().BaseURLS.Idm)
}

func (root *Root) Calendar() calendarv1connect.CalendarServiceClient {
	return calendarv1connect.NewCalendarServiceClient(root.HttpClient, root.Config().BaseURLS.Calendar)
}

func (root *Root) CallService() pbx3cxv1connect.CallServiceClient {
	return pbx3cxv1connect.NewCallServiceClient(root.HttpClient, root.Config().BaseURLS.CallService)
}

func (root *Root) Comments() commentv1connect.CommentServiceClient {
	return commentv1connect.NewCommentServiceClient(root.HttpClient, root.Config().BaseURLS.CommentService)
}
