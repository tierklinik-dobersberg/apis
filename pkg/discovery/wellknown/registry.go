package wellknown

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/bufbuild/connect-go"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/calendar/v1/calendarv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/comment/v1/commentv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/customer/v1/customerv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/office_hours/v1/office_hoursv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1/rosterv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/tasks/v1/tasksv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/typeserver/v1/typeserverv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/cli"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
)

type Service[T any] struct {
	Name    string
	Factory func(cli connect.HTTPClient, ep string, opts ...connect.ClientOption) T
}

func (s *Service[T]) Create(ctx context.Context, discoverer discovery.Discoverer, opts ...connect.ClientOption) (T, error) {
	return NewClient(ctx, discoverer, *s, opts...)
}

func Create[T any](name string, factory func(connect.HTTPClient, string, ...connect.ClientOption) T) Service[T] {
	return Service[T]{
		Name:    name,
		Factory: factory,
	}
}

var (
	CalendarV1ServiceScope   = "tkd.calendar.v1"
	CommentV1ServiceScope    = "tkd.comment.v1"
	CustomerV1ServiceScope   = "tkd.customer.v1"
	EventV1ServiceScropt     = "tkd.events.v1"
	IdmV1ServiceScope        = "tkd.idm.v1"
	OfficeHourV1ServiceScope = "tkd.office_hours.v1"
	Pbx3cxV1ServiceScope     = "tkd.pbx3cx.v1"
	RosterV1ServiceScope     = "tkd.roster.v1"
	TaskV1ServiceScope       = "tkd.tasks.v1"
	TypeV1ServiceScope       = "tkd.typeserver.v1"
)

var (
	// tkd/calendar/v1
	CalendarService = Create(CalendarV1ServiceScope, calendarv1connect.NewCalendarServiceClient)
	HolidayService  = Create(CalendarV1ServiceScope, calendarv1connect.NewHolidayServiceClient)

	// tkd/comment/v1
	CommentService = Create(CommentV1ServiceScope, commentv1connect.NewCommentServiceClient)

	// tkd/customer/v1
	CustomerService       = Create(CustomerV1ServiceScope, customerv1connect.NewCustomerServiceClient)
	CustomerImportService = Create(CustomerV1ServiceScope, customerv1connect.NewCustomerImportServiceClient)

	// tkd/events/v1
	EventService = Create(EventV1ServiceScropt, eventsv1connect.NewEventServiceClient)

	// tkd/idm/v1
	UserService   = Create(IdmV1ServiceScope, idmv1connect.NewUserServiceClient)
	RoleService   = Create(IdmV1ServiceScope, idmv1connect.NewRoleServiceClient)
	NotifyService = Create(IdmV1ServiceScope, idmv1connect.NewNotifyServiceClient)
	AuthService   = Create(IdmV1ServiceScope, idmv1connect.NewAuthServiceClient)

	// tkd/office_hours/v1
	OfficeHoursService = Create(OfficeHourV1ServiceScope, office_hoursv1connect.NewOfficeHourServiceClient)

	// tkd/pbx3cx/v1
	CallService      = Create(Pbx3cxV1ServiceScope, pbx3cxv1connect.NewCallServiceClient)
	VoiceMailService = Create(Pbx3cxV1ServiceScope, pbx3cxv1connect.NewVoiceMailServiceClient)

	// tkd/roster/v1
	RosterService    = Create(RosterV1ServiceScope, rosterv1connect.NewRosterServiceClient)
	WorkShiftService = Create(RosterV1ServiceScope, rosterv1connect.NewWorkShiftServiceClient)
	WorkTimeService  = Create(RosterV1ServiceScope, rosterv1connect.NewWorkTimeServiceClient)
	OffTimeService   = Create(RosterV1ServiceScope, rosterv1connect.NewOffTimeServiceClient)

	// tkd/tasks/v1
	BoardService = Create(TaskV1ServiceScope, tasksv1connect.NewBoardServiceClient)
	TaskService  = Create(TaskV1ServiceScope, tasksv1connect.NewTaskServiceClient)

	// tkd/typeserver/v1
	TypeService = Create(TypeV1ServiceScope, typeserverv1connect.NewTypeResolverServiceClient)
)

func NewClient[T any](ctx context.Context, d discovery.Discoverer, wks Service[T], opts ...connect.ClientOption) (T, error) {
	var res T

	svc, err := d.Discover(ctx, wks.Name)
	if err != nil {
		return res, err
	}

	if len(svc) == 0 {
		return res, fmt.Errorf("no service instances found")
	}

	i := svc[rand.IntN(len(svc))]

	return wks.Factory(cli.NewInsecureHttp2Client(), i.Address, opts...), nil
}
