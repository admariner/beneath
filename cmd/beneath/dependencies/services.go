package dependencies

import (
	"time"

	"github.com/beneath-hq/beneath/cmd/beneath/cli"
	"github.com/beneath-hq/beneath/services/data"
	"github.com/beneath-hq/beneath/services/middleware"
	"github.com/beneath-hq/beneath/services/organization"
	"github.com/beneath-hq/beneath/services/permissions"
	"github.com/beneath-hq/beneath/services/project"
	"github.com/beneath-hq/beneath/services/secret"
	"github.com/beneath-hq/beneath/services/service"
	"github.com/beneath-hq/beneath/services/stream"
	"github.com/beneath-hq/beneath/services/usage"
	"github.com/beneath-hq/beneath/services/user"
)

// TO ADD A NEW SERVICE:
// 1. Add it as a member to AllServices
// 2. Add it to NewAllServices (two places!)
// 3. Add it as a dependency in init()

// AllServices is a convenience wrapper that initializes all services
type AllServices struct {
	Data         *data.Service
	Usage        *usage.Service
	Middleware   *middleware.Service
	Organization *organization.Service
	Permissions  *permissions.Service
	Project      *project.Service
	Secret       *secret.Service
	Service      *service.Service
	Stream       *stream.Service
	User         *user.Service
}

// NewAllServices creates a new AllServices
func NewAllServices(
	data *data.Service,
	usage *usage.Service,
	middleware *middleware.Service,
	organization *organization.Service,
	permissions *permissions.Service,
	project *project.Service,
	secret *secret.Service,
	service *service.Service,
	stream *stream.Service,
	user *user.Service,
) *AllServices {
	return &AllServices{
		Data:         data,
		Usage:        usage,
		Middleware:   middleware,
		Organization: organization,
		Permissions:  permissions,
		Project:      project,
		Secret:       secret,
		Service:      service,
		Stream:       stream,
		User:         user,
	}
}

func init() {
	cli.AddDependency(NewAllServices)
	cli.AddDependency(data.New)
	cli.AddDependency(usage.New)
	cli.AddDependency(middleware.New)
	cli.AddDependency(organization.New)
	cli.AddDependency(permissions.New)
	cli.AddDependency(project.New)
	cli.AddDependency(secret.New)
	cli.AddDependency(service.New)
	cli.AddDependency(stream.New)
	cli.AddDependency(user.New)

	// the usage service takes some extra options
	cli.AddDependency(func() *usage.Options {
		return &usage.Options{
			CacheSize:      2500,
			CommitInterval: 30 * time.Second,
		}
	})
}
