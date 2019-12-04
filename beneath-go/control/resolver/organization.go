package resolver

import (
	"context"

	"github.com/beneath-core/beneath-go/control/entity"
	"github.com/beneath-core/beneath-go/control/gql"
	"github.com/beneath-core/beneath-go/core/middleware"
	"github.com/beneath-core/beneath-go/core/stripe"
	uuid "github.com/satori/go.uuid"
	"github.com/vektah/gqlparser/gqlerror"
)

// Organization returns the gql.OrganizationResolver
func (r *Resolver) Organization() gql.OrganizationResolver {
	return &organizationResolver{r}
}

type organizationResolver struct{ *Resolver }

func (r *organizationResolver) OrganizationID(ctx context.Context, obj *entity.Organization) (string, error) {
	return obj.OrganizationID.String(), nil
}

// PaymentMethod returns the gql.PaymentMethodResolver
func (r *Resolver) PaymentMethod() gql.PaymentMethodResolver {
	return &paymentMethodResolver{r}
}

type paymentMethodResolver struct{ *Resolver }

func (r *paymentMethodResolver) OrganizationID(ctx context.Context, obj *entity.PaymentMethod) (string, error) {
	return obj.OrganizationID.String(), nil
}

func (r *paymentMethodResolver) Type(ctx context.Context, obj *entity.PaymentMethod) (string, error) {
	return string(obj.Type), nil
}

func (r *paymentMethodResolver) Card(ctx context.Context, obj *entity.PaymentMethod) (*gql.Card, error) {
	return &gql.Card{
		Brand: obj.Card.Brand,
		Last4: obj.Card.Last4,
	}, nil
}

func (r *queryResolver) OrganizationByName(ctx context.Context, name string) (*entity.Organization, error) {
	organization := entity.FindOrganizationByName(ctx, name)
	if organization == nil {
		return nil, gqlerror.Errorf("Organization %s not found", name)
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organization.OrganizationID)
	if !perms.View {
		return nil, gqlerror.Errorf("Not allowed to view organization %s", name)
	}

	return organization, nil
}

func (r *queryResolver) GetUserOrganizationPermissions(ctx context.Context, userID uuid.UUID, organizationID uuid.UUID) (*entity.PermissionsUsersOrganizations, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return nil, gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions on organization %s", organizationID.String())
	}

	user := entity.FindUser(ctx, userID)
	if user == nil {
		return nil, gqlerror.Errorf("User %s not found", userID.String())
	}

	permissions := entity.FindPermissionsUsersOrganizations(ctx, userID, organizationID)
	if permissions == nil {
		return nil, gqlerror.Errorf("Permissions not found for user %s in organization %s", userID.String(), organizationID.String())
	}

	return permissions, nil
}

func (r *mutationResolver) CreateOrganization(ctx context.Context, name string) (*entity.Organization, error) {
	secret := middleware.GetSecret(ctx)
	if !secret.IsUser() {
		return nil, gqlerror.Errorf("Not allowed to create organization")
	}

	org, err := entity.CreateOrganizationWithUser(ctx, name, secret.GetOwnerID())
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (r *mutationResolver) AddUserToOrganization(ctx context.Context, username string, organizationID uuid.UUID, view bool, admin bool) (*entity.User, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return nil, gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions on organization %s", organizationID.String())
	}

	user := entity.FindUserByUsername(ctx, username)
	if user == nil {
		return nil, gqlerror.Errorf("No user found with that username")
	}

	err := organization.AddUser(ctx, user.UserID, view, admin)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return user, nil
}

func (r *mutationResolver) RemoveUserFromOrganization(ctx context.Context, userID uuid.UUID, organizationID uuid.UUID) (bool, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return false, gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return false, gqlerror.Errorf("Not allowed to perform admin functions in organization %s", organizationID.String())
	}

	if len(organization.Users) < 2 {
		return false, gqlerror.Errorf("Can't remove last member of organization")
	}

	err := organization.RemoveUser(ctx, userID)
	if err != nil {
		return false, gqlerror.Errorf(err.Error())
	}

	return true, nil
}

func (r *mutationResolver) UpdateUserOrganizationPermissions(ctx context.Context, userID uuid.UUID, organizationID uuid.UUID, view *bool, admin *bool) (*entity.PermissionsUsersOrganizations, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return nil, gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions in organization %s", organizationID.String())
	}

	user := entity.FindUser(ctx, userID)
	if user == nil {
		return nil, gqlerror.Errorf("User %s not found", userID.String())
	}

	permissions := entity.FindPermissionsUsersOrganizations(ctx, userID, organizationID)
	if permissions == nil {
		return nil, gqlerror.Errorf("Permissions for not found for organization %s and user %s", organizationID.String(), userID.String())
	}

	// TODO: change this to organization.UpdateUserPermissions
	// permissions, err := permissions.Update(ctx, view, admin)
	// if err != nil {
	// 	return nil, gqlerror.Errorf("Failed to update permissions")
	// }

	return permissions, nil
}

func (r *mutationResolver) UpdateUserOrganizationQuotas(ctx context.Context, userID uuid.UUID, organizationID uuid.UUID, readQuota *int, writeQuota *int) (*entity.User, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return nil, gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions in organization %s", organizationID.String())
	}

	user := entity.FindUser(ctx, userID)
	if user == nil {
		return nil, gqlerror.Errorf("User %s not found", userID.String())
	}

	// TODO: change this to organization.UpdateUserQuotas
	// user, err := user.UpdateQuotas(ctx, readQuota, writeQuota)
	// if err != nil {
	// 	return nil, gqlerror.Errorf("Failed to update the user's quotas")
	// }

	return user, nil
}

func (r *mutationResolver) UpdateBillingPlan(ctx context.Context, organizationID uuid.UUID, billingPlanID uuid.UUID) (*entity.Organization, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return nil, gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions in organization %s", organizationID.String())
	}

	billingPlan := entity.FindBillingPlan(ctx, billingPlanID)
	if billingPlan == nil {
		return nil, gqlerror.Errorf("Billing plan %s not found", billingPlanID.String())
	}

	organization, err := organization.UpdateBillingPlanID(ctx, billingPlanID)
	if err != nil {
		return nil, gqlerror.Errorf("Failed to update the organization's billing plan")
	}

	return organization, nil
}

func (r *mutationResolver) CreateStripeSetupIntent(ctx context.Context, organizationID uuid.UUID, billingPlanID uuid.UUID) (string, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return "", gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return "", gqlerror.Errorf("Not allowed to perform admin functions in organization %s", organizationID.String())
	}

	setupIntent := stripe.CreateSetupIntent(organizationID, billingPlanID)
	if setupIntent == nil {
		return "", gqlerror.Errorf("Unable to create setup intent")
	}

	return setupIntent.ClientSecret, nil
}

func (r *queryResolver) GetCurrentPaymentMethod(ctx context.Context, organizationID uuid.UUID) (*entity.PaymentMethod, error) {
	organization := entity.FindOrganization(ctx, organizationID)
	if organization == nil {
		return nil, gqlerror.Errorf("Organization %s not found", organizationID.String())
	}

	secret := middleware.GetSecret(ctx)
	perms := secret.OrganizationPermissions(ctx, organizationID)
	if !perms.Admin {
		return nil, gqlerror.Errorf("Not allowed to perform admin functions in organization %s", organizationID.String())
	}

	paymentMethod := &entity.PaymentMethod{}
	if organization.PaymentMethod == entity.PaymentMethodCard {
		paymentMethodID := stripe.GetCustomerRecentPaymentMethodID(organization.StripeCustomerID, string(entity.PaymentMethodCard))
		stripePaymentMethod := stripe.RetrievePaymentMethod(paymentMethodID)
		paymentMethod = &entity.PaymentMethod{
			OrganizationID: organizationID,
			Type:           entity.PaymentMethodCard,
			Card: &entity.Card{
				Brand: string(stripePaymentMethod.Card.Brand),
				Last4: stripePaymentMethod.Card.Last4,
			},
		}
	} else if organization.PaymentMethod == entity.PaymentMethodWire {
		paymentMethod = &entity.PaymentMethod{
			OrganizationID: organizationID,
			Type:           entity.PaymentMethodWire,
			Card:           &entity.Card{},
		}
	} else {
		return nil, gqlerror.Errorf("Organization does not have a payment method")
	}

	return paymentMethod, nil
}
