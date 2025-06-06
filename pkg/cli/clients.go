package cli

import (
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/calendar/v1/calendarv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/comment/v1/commentv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/customer/v1/customerv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1/longrunningv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/office_hours/v1/office_hoursv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/orthanc_bridge/v1/orthanc_bridgev1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/printing/v1/printingv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1/rosterv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/tasks/v1/tasksv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/events"
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

func (root *Root) Customer() customerv1connect.CustomerServiceClient {
	return customerv1connect.NewCustomerServiceClient(root.HttpClient, root.Config().BaseURLS.CustomerService)
}

func (root *Root) Patient() customerv1connect.PatientServiceClient {
	return customerv1connect.NewPatientServiceClient(root.HttpClient, root.Config().BaseURLS.CustomerService)
}

func (root *Root) CustomerImport() customerv1connect.CustomerImportServiceClient {
	return customerv1connect.NewCustomerImportServiceClient(root.HttpClient, root.Config().BaseURLS.CustomerService)
}

func (root *Root) Boards() tasksv1connect.BoardServiceClient {
	return tasksv1connect.NewBoardServiceClient(root.HttpClient, root.Config().BaseURLS.TaskService)
}

func (root *Root) Tasks() tasksv1connect.TaskServiceClient {
	return tasksv1connect.NewTaskServiceClient(root.HttpClient, root.Config().BaseURLS.TaskService)
}

func (root *Root) OfficeHourService() office_hoursv1connect.OfficeHourServiceClient {
	return office_hoursv1connect.NewOfficeHourServiceClient(root.HttpClient, root.Config().OfficeHourService)
}

func (root *Root) EventsService() eventsv1connect.EventServiceClient {
	return eventsv1connect.NewEventServiceClient(root.HttpClient, root.Config().EventsService)
}

func (root *Root) OrthancBridge() orthanc_bridgev1connect.OrthancBridgeClient {
	return orthanc_bridgev1connect.NewOrthancBridgeClient(root.HttpClient, root.Config().OrthancBridge)
}

func (root *Root) EventStreamClient(factory events.ClientFactory) (*events.Client, error) {
	cli := events.NewClient(factory)

	if err := cli.Start(root.Context()); err != nil {
		return nil, err
	}

	return cli, nil
}

func (root *Root) LongRunning() longrunningv1connect.LongRunningServiceClient {
	return longrunningv1connect.NewLongRunningServiceClient(root.HttpClient, root.Config().LongRunning)
}

func (root *Root) PrintService() printingv1connect.PrintServiceClient {
	return printingv1connect.NewPrintServiceClient(root.HttpClient, root.Config().PrintService)
}
