package resolver

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/vektah/gqlparser/gqlerror"

	"gitlab.com/beneath-hq/beneath/control/entity"
	"gitlab.com/beneath-hq/beneath/control/gql"
	"gitlab.com/beneath-hq/beneath/internal/middleware"
)

// Project returns the gql.ProjectResolver
func (r *Resolver) Project() gql.ProjectResolver {
	return &projectResolver{r}
}

type projectResolver struct{ *Resolver }

func (r *projectResolver) ProjectID(ctx context.Context, obj *entity.Project) (string, error) {
	return obj.ProjectID.String(), nil
}

func (r *queryResolver) ExploreProjects(ctx context.Context) ([]*entity.Project, error) {
	return entity.FindProjects(ctx), nil
}

func (r *queryResolver) ProjectByOrganizationAndName(ctx context.Context, organizationName string, projectName string) (*entity.Project, error) {
	project := entity.FindProjectByOrganizationAndName(ctx, organizationName, projectName)
	if project == nil {
		return nil, gqlerror.Errorf("Project %s/%s not found", organizationName, projectName)
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.ProjectPermissions(ctx, project.ProjectID, project.Public)
	if !perms.View {
		return nil, gqlerror.Errorf("Not allowed to read project %s/%s", organizationName, projectName)
	}

	return projectWithPermissions(project, perms), nil
}

func (r *queryResolver) ProjectByID(ctx context.Context, projectID uuid.UUID) (*entity.Project, error) {
	project := entity.FindProject(ctx, projectID)
	if project == nil {
		return nil, gqlerror.Errorf("Project %s not found", projectID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.ProjectPermissions(ctx, projectID, project.Public)
	if !perms.View {
		return nil, gqlerror.Errorf("Not allowed to read project %s", projectID.String())
	}

	return projectWithPermissions(project, perms), nil
}

func (r *queryResolver) ProjectMembers(ctx context.Context, projectID uuid.UUID) ([]*entity.ProjectMember, error) {
	secret := middleware.GetSecret(ctx)
	perms := secret.ProjectPermissions(ctx, projectID, false)
	if !perms.View {
		return nil, gqlerror.Errorf("You're not allowed to see the members of project %s", projectID.String())
	}

	return entity.FindProjectMembers(ctx, projectID)
}

func (r *mutationResolver) CreateProject(ctx context.Context, name string, displayName *string, organizationID uuid.UUID, public bool, site *string, description *string, photoURL *string) (*entity.Project, error) {
	secret := middleware.GetSecret(ctx)
	if !secret.IsUser() {
		return nil, gqlerror.Errorf("Only users can create projects")
	}

	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Create {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions on organization %s", organizationID.String())
	}

	bi := entity.FindBillingInfo(ctx, organizationID)
	if bi == nil {
		return nil, gqlerror.Errorf("Could not find billing info for organization %s", organizationID.String())
	}

	if !public && !bi.BillingPlan.PrivateProjects {
		return nil, gqlerror.Errorf("Your organization's billing plan does not permit private projects")
	}

	project := &entity.Project{
		Name:           name,
		DisplayName:    DereferenceString(displayName),
		OrganizationID: organizationID,
		Site:           DereferenceString(site),
		Description:    DereferenceString(description),
		PhotoURL:       DereferenceString(photoURL),
		Public:         public,
	}

	projPerms := entity.ProjectPermissions{
		View:   true,
		Create: true,
		Admin:  true,
	}

	err := project.CreateWithUser(ctx, secret.GetOwnerID(), projPerms)
	if err != nil {
		return nil, err
	}

	// refetching project to include user
	project = entity.FindProject(ctx, project.ProjectID)
	if project == nil {
		panic(fmt.Errorf("expected project with ID %s to exist", project.ProjectID.String()))
	}

	return projectWithPermissions(project, projPerms), nil
}

func (r *mutationResolver) UpdateProject(ctx context.Context, projectID uuid.UUID, displayName *string, public *bool, site *string, description *string, photoURL *string) (*entity.Project, error) {
	project := entity.FindProject(ctx, projectID)
	if project == nil {
		return nil, gqlerror.Errorf("Project %s not found", projectID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.ProjectPermissions(ctx, projectID, false)
	if !perms.Admin {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions on project %s", projectID.String())
	}

	bi := entity.FindBillingInfo(ctx, project.OrganizationID)
	if bi == nil {
		return nil, gqlerror.Errorf("Could not find billing info for organization %s", project.OrganizationID.String())
	}

	if public != nil && !*public && !bi.BillingPlan.PrivateProjects {
		return nil, gqlerror.Errorf("Your organization's billing plan does not permit private projects")
	}

	err := project.UpdateDetails(ctx, displayName, public, site, description, photoURL)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return projectWithPermissions(project, perms), nil
}

func (r *mutationResolver) DeleteProject(ctx context.Context, projectID uuid.UUID) (bool, error) {
	secret := middleware.GetSecret(ctx)
	perms := secret.ProjectPermissions(ctx, projectID, false)
	if !perms.Admin {
		return false, gqlerror.Errorf("Not allowed to perform admin functions in project %s", projectID.String())
	}

	project := entity.FindProject(ctx, projectID)
	if project == nil {
		return false, gqlerror.Errorf("Project %s not found", projectID.String())
	}

	err := project.Delete(ctx)
	if err != nil {
		return false, gqlerror.Errorf(err.Error())
	}

	return true, nil
}

func projectWithPermissions(p *entity.Project, perms entity.ProjectPermissions) *entity.Project {
	if perms.View || perms.Create || perms.Admin {
		p.Permissions = &entity.PermissionsUsersProjects{
			View:   perms.View,
			Create: perms.Create,
			Admin:  perms.Admin,
		}
	}
	return p
}
