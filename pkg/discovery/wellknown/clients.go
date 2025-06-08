package wellknown

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/calendar/v1/calendarv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/comment/v1/commentv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/customer/v1/customerv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1/longrunningv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/office_hours/v1/office_hoursv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/printing/v1/printingv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1/rosterv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/tasks/v1/tasksv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/treatment/v1/treatmentv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/typeserver/v1/typeserverv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
	"github.com/tierklinik-dobersberg/apis/pkg/h2utils"
)

type ConfigureClientOptions struct {
	SerivceEnvMap              map[ServiceScope]string
	HTTPClient                 *http.Client
	GlobalOptions              []connect.ClientOption
	Catalog                    discovery.Discoverer
	DisableEnvironmentOverride bool
}

func (opts ConfigureClientOptions) url(scope ServiceScope) string {
	if !opts.DisableEnvironmentOverride && opts.SerivceEnvMap != nil {
		val := os.Getenv(opts.SerivceEnvMap[scope])
		if val != "" {
			return val
		}
	}

	return fmt.Sprintf("http://%s", string(scope))
}

type Clients struct {
	CalendarService       calendarv1connect.CalendarServiceClient
	HolidayService        calendarv1connect.HolidayServiceClient
	CommentService        commentv1connect.CommentServiceClient
	CustomerService       customerv1connect.CustomerServiceClient
	CustomerImportService customerv1connect.CustomerImportServiceClient
	PatientService        customerv1connect.PatientServiceClient
	EventService          eventsv1connect.EventServiceClient
	UserService           idmv1connect.UserServiceClient
	RoleService           idmv1connect.RoleServiceClient
	NotifyService         idmv1connect.NotifyServiceClient
	AuthService           idmv1connect.AuthServiceClient
	OfficeHourService     office_hoursv1connect.OfficeHourServiceClient
	CallService           pbx3cxv1connect.CallServiceClient
	VoiceMailService      pbx3cxv1connect.VoiceMailServiceClient
	RosterService         rosterv1connect.RosterServiceClient
	WorkShiftService      rosterv1connect.WorkShiftServiceClient
	WorkTimeService       rosterv1connect.WorkTimeServiceClient
	OffTimeService        rosterv1connect.OffTimeServiceClient
	BoardService          tasksv1connect.BoardServiceClient
	TaskService           tasksv1connect.TaskServiceClient
	TypeService           typeserverv1connect.TypeResolverServiceClient
	LongRunningService    longrunningv1connect.LongRunningServiceClient
	PrintService          printingv1connect.PrintServiceClient
	TreatmentService      treatmentv1connect.TreatmentServiceClient
	SpeciesService        treatmentv1connect.SpeciesServiceClient
}

func ConfigureClients(opts ConfigureClientOptions) Clients {
	if opts.SerivceEnvMap == nil {
		opts.SerivceEnvMap = ServiceEnvMap
	}

	httpClient := opts.HTTPClient
	if httpClient == nil {
		httpClient = h2utils.WithDiscovery(opts.Catalog, h2utils.NewInsecureHttp2Client())
	}

	clients := Clients{
		CalendarService:       calendarv1connect.NewCalendarServiceClient(httpClient, opts.url(CalendarV1ServiceScope), opts.GlobalOptions...),
		HolidayService:        calendarv1connect.NewHolidayServiceClient(httpClient, opts.url(CalendarV1ServiceScope), opts.GlobalOptions...),
		CommentService:        commentv1connect.NewCommentServiceClient(httpClient, opts.url(CommentV1ServiceScope), opts.GlobalOptions...),
		CustomerService:       customerv1connect.NewCustomerServiceClient(httpClient, opts.url(CustomerV1ServiceScope), opts.GlobalOptions...),
		CustomerImportService: customerv1connect.NewCustomerImportServiceClient(httpClient, opts.url(CustomerV1ServiceScope), opts.GlobalOptions...),
		PatientService:        customerv1connect.NewPatientServiceClient(httpClient, opts.url(CustomerV1ServiceScope), opts.GlobalOptions...),
		EventService:          eventsv1connect.NewEventServiceClient(httpClient, opts.url(EventV1ServiceScope), opts.GlobalOptions...),
		UserService:           idmv1connect.NewUserServiceClient(httpClient, opts.url(IdmV1ServiceScope), opts.GlobalOptions...),
		RoleService:           idmv1connect.NewRoleServiceClient(httpClient, opts.url(IdmV1ServiceScope), opts.GlobalOptions...),
		NotifyService:         idmv1connect.NewNotifyServiceClient(httpClient, opts.url(IdmV1ServiceScope), opts.GlobalOptions...),
		AuthService:           idmv1connect.NewAuthServiceClient(httpClient, opts.url(IdmV1ServiceScope), opts.GlobalOptions...),
		OfficeHourService:     office_hoursv1connect.NewOfficeHourServiceClient(httpClient, opts.url(OfficeHourV1ServiceScope), opts.GlobalOptions...),
		CallService:           pbx3cxv1connect.NewCallServiceClient(httpClient, opts.url(Pbx3cxV1ServiceScope), opts.GlobalOptions...),
		VoiceMailService:      pbx3cxv1connect.NewVoiceMailServiceClient(httpClient, opts.url(Pbx3cxV1ServiceScope), opts.GlobalOptions...),
		RosterService:         rosterv1connect.NewRosterServiceClient(httpClient, opts.url(RosterV1ServiceScope), opts.GlobalOptions...),
		WorkShiftService:      rosterv1connect.NewWorkShiftServiceClient(httpClient, opts.url(RosterV1ServiceScope), opts.GlobalOptions...),
		WorkTimeService:       rosterv1connect.NewWorkTimeServiceClient(httpClient, opts.url(RosterV1ServiceScope), opts.GlobalOptions...),
		OffTimeService:        rosterv1connect.NewOffTimeServiceClient(httpClient, opts.url(RosterV1ServiceScope), opts.GlobalOptions...),
		BoardService:          tasksv1connect.NewBoardServiceClient(httpClient, opts.url(TaskV1ServiceScope), opts.GlobalOptions...),
		TaskService:           tasksv1connect.NewTaskServiceClient(httpClient, opts.url(TaskV1ServiceScope), opts.GlobalOptions...),
		TypeService:           typeserverv1connect.NewTypeResolverServiceClient(httpClient, opts.url(TypeV1ServiceScope), opts.GlobalOptions...),
		LongRunningService:    longrunningv1connect.NewLongRunningServiceClient(httpClient, opts.url(LongrunningV1ServiceScope), opts.GlobalOptions...),
		PrintService:          printingv1connect.NewPrintServiceClient(httpClient, opts.url(PrintV1ServiceScope), opts.GlobalOptions...),
		TreatmentService:      treatmentv1connect.NewTreatmentServiceClient(httpClient, opts.url(TreatmentV1ServiceScope), opts.GlobalOptions...),
		SpeciesService:        treatmentv1connect.NewSpeciesServiceClient(httpClient, opts.url(TreatmentV1ServiceScope), opts.GlobalOptions...),
	}

	return clients
}
