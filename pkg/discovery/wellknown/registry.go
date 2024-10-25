package wellknown

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"

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

func (s *Service[T]) Register(ctx context.Context, discoverer discovery.Discoverer, addr string, opts ...RegisterOption) error {
	return Register(ctx, s, discoverer, addr, opts...)
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
	// tkd/calendar/v1
	CalendarService = Create("tkd.calendar.v1.CalendarService", calendarv1connect.NewCalendarServiceClient)
	HolidayService  = Create("tkd.calendar.v1.HolidayService", calendarv1connect.NewHolidayServiceClient)

	// tkd/comment/v1
	CommentService = Create("tkd.comment.v1.CommentService", commentv1connect.NewCommentServiceClient)

	// tkd/customer/v1
	CustomerService       = Create("tkd.customer.v1.CustomerService", customerv1connect.NewCustomerServiceClient)
	CustomerImportService = Create("tkd.customer.v1.CustomerImportService", customerv1connect.NewCustomerImportServiceClient)

	// tkd/events/v1
	EventService = Create("tkd.events.v1.EventService", eventsv1connect.NewEventServiceClient)

	// tkd/idm/v1
	UserService   = Create("tkd.idm.v1.UserService", idmv1connect.NewUserServiceClient)
	RoleService   = Create("tkd.idm.v1.RoleService", idmv1connect.NewRoleServiceClient)
	NotifyService = Create("tkd.idm.v1.NotifyService", idmv1connect.NewNotifyServiceClient)
	AuthService   = Create("tkd.idm.v1.AuthService", idmv1connect.NewAuthServiceClient)

	// tkd/office_hours/v1
	OfficeHoursService = Create("tkd.office_hours.v1.OfficeHourService", office_hoursv1connect.NewOfficeHourServiceClient)

	// tkd/pbx3cx/v1
	CallService      = Create("tkd.pbx3cx.v1.CallService", pbx3cxv1connect.NewCallServiceClient)
	VoiceMailService = Create("tkd.pbx3cx.v1.VoiceMailService", pbx3cxv1connect.NewVoiceMailServiceClient)

	// tkd/roster/v1
	RosterService    = Create("tkd.roster.v1.RosterService", rosterv1connect.NewRosterServiceClient)
	WorkShiftService = Create("tkd.roster.v1.WorkShiftService", rosterv1connect.NewWorkShiftServiceClient)
	WorkTimeService  = Create("tkd.roster.v1.WorkTimeService", rosterv1connect.NewWorkTimeServiceClient)
	OffTimeService   = Create("tkd.roster.v1.OffTimeService", rosterv1connect.NewOffTimeServiceClient)

	// tkd/tasks/v1
	BoardService = Create("tkd.tasks.v1.BoardService", tasksv1connect.NewBoardServiceClient)
	TaskService  = Create("tkd.tasks.v1.TaskService", tasksv1connect.NewTaskServiceClient)

	// tkd/typeserver/v1
	TypeService = Create("tkd.typeserver.v1.TypeResolverService", typeserverv1connect.NewTypeResolverServiceClient)
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

type RegisterOption func(instance *discovery.ServiceInstance)

func Register[T any](ctx context.Context, wk *Service[T], discoverer discovery.Discoverer, addr string, opts ...RegisterOption) error {
	instance := &discovery.ServiceInstance{
		Name:    wk.Name,
		Address: addr,
	}

	for _, opt := range opts {
		opt(instance)
	}

	if instance.Instance == "" {
		if i := os.Getenv("INSTANCE_ID"); i != "" {
			instance.Instance = i
		} else {
			hostname, err := os.Hostname()
			if err != nil {
				return fmt.Errorf("failed to get hostname: %w", err)
			}

			instance.Instance = hostname
		}
	}

	return discovery.Register(ctx, discoverer, *instance)
}

func WithInstanceID(id string) RegisterOption {
	return func(instance *discovery.ServiceInstance) {
		instance.Instance = id
	}
}

func WithMeta(meta map[string]string) RegisterOption {
	return func(instance *discovery.ServiceInstance) {
		instance.Meta = meta
	}
}
