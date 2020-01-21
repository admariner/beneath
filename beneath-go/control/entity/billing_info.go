package entity

import (
	"context"
	"time"

	"github.com/beneath-core/beneath-go/db"
	uuid "github.com/satori/go.uuid"
)

// BillingInfo encapsulates an organization's billing information
type BillingInfo struct {
	OrganizationID uuid.UUID `sql:",pk,type:uuid,default:uuid_generate_v4()"`
	BillingPlanID  uuid.UUID `sql:"on_delete:RESTRICT,notnull,type:uuid"`
	BillingPlan    *BillingPlan
	PaymentsDriver PaymentsDriver // StripeCustomerID is currently the only thing that goes in the payload
	DriverPayload  map[string]interface{}
	CreatedOn      time.Time `sql:",default:now()"`
	UpdatedOn      time.Time `sql:",default:now()"`
	DeletedOn      time.Time
	Services       []*Service
	Users          []*User
}

// UpdateBillingInfo creates or updates an organization's billing info
func UpdateBillingInfo(ctx context.Context, organizationID uuid.UUID, billingPlanID uuid.UUID, paymentsDriver PaymentsDriver, driverPayload map[string]interface{}) (*BillingInfo, error) {
	prevBillingInfo := FindBillingInfo(ctx, organizationID)
	if prevBillingInfo == nil {
		panic("could not get existing billing info")
	}

	bi := &BillingInfo{}

	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // defer rollback on error

	// set fields
	bi.OrganizationID = organizationID
	bi.BillingPlanID = billingPlanID
	bi.PaymentsDriver = paymentsDriver
	bi.DriverPayload = driverPayload

	// validate
	err = GetValidator().Struct(bi)
	if err != nil {
		return nil, err
	}

	// insert
	_, err = tx.ModelContext(ctx, bi).OnConflict("(organization_id) DO UPDATE").Insert()
	if err != nil {
		return nil, err
	}

	// update the organization's users' quotas
	// this gets us access to bi.Users
	err = tx.ModelContext(ctx, bi).WherePK().Column("billing_info.*", "BillingPlan", "Services", "Users").Select()
	if !AssertFoundOne(err) {
		return nil, err
	}

	for _, u := range bi.Users {
		u.ReadQuota = bi.BillingPlan.SeatReadQuota
		u.WriteQuota = bi.BillingPlan.SeatWriteQuota
		_, err := tx.ModelContext(ctx, u).Column("read_quota", "write_quota", "updated_on").WherePK().Update()
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// TODO: make this part of the entire transaction; Q: any easy way to encompass function calls in the transaction?
	// if switching billing plans, we reconcile the next month's bill:
	// - we provide prorated credit for seats on the old billing plan for the remaining days of the current month
	// - we charge a prorated amount for the seats on the new billing plan for the remaining days of the current month
	// - the entire month's overages will be charged at the rates of the new billing plan
	if billingPlanID != prevBillingInfo.BillingPlanID {
		var userIDs []uuid.UUID
		for _, user := range bi.Users {
			userIDs = append(userIDs, user.UserID)
		}

		// give organization pro-rated credit for seats at old billing plan price
		if prevBillingInfo.BillingPlanID.String() != FreeBillingPlanID {
			err := commitProratedSeatsToBill(ctx, prevBillingInfo, userIDs, true)
			if err != nil {
				panic("could not commit prorated seat credits to bill")
			}
		}

		// charge organization the pro-rated amount for seats at new billing plan price
		err = commitProratedSeatsToBill(ctx, bi, userIDs, false)
		if err != nil {
			panic("could not commit prorated seats to bill")
		}

		// TODO
		// if downgrading from private projects, lock down organization's outstanding private projects
		// if !bi.BillingPlan.PrivateProjects && prevBillingInfo.BillingPlan.PrivateProjects {
		// 	// option1: lock projects so that they can't be viewed
		// 	// option2: clear existing project permissions, and ensure they can't be added back
		// }
	}

	return bi, nil
}

// FindBillingInfo finds an organization's billing info by organizationID
func FindBillingInfo(ctx context.Context, organizationID uuid.UUID) *BillingInfo {
	billingInfo := &BillingInfo{
		OrganizationID: organizationID,
	}
	err := db.DB.ModelContext(ctx, billingInfo).WherePK().Column("billing_info.*", "BillingPlan", "Services", "Users").Select()
	if !AssertFoundOne(err) {
		return nil
	}
	return billingInfo
}

// UpdateBillingPlanID updates an organization's billing plan ID
// func (bi *BillingInfo) UpdateBillingPlanID(ctx context.Context, billingPlanID uuid.UUID) (*BillingInfo, error) {
// 	bi.BillingPlanID = billingPlanID
// 	bi.UpdatedOn = time.Now()
// 	_, err := db.DB.ModelContext(ctx, bi).Column("billing_plan_id", "updated_on").WherePK().Update()
// 	return bi, err
// }

// UpdateStripeCustomerID updates an organization's Stripe Customer ID
// func (bi *BillingInfo) UpdateStripeCustomerID(ctx context.Context, stripeCustomerID string) error {
// 	// bi.StripeCustomerID = stripeCustomerID
// 	// bi.UpdatedOn = time.Now()
// 	// _, err := db.DB.ModelContext(ctx, bi).Column("stripe_customer_id", "updated_on").WherePK().Update()
// 	bi.DriverPayload["customer_id"] = stripeCustomerID
// 	bi.UpdatedOn = time.Now()
// 	_, err := db.DB.ModelContext(ctx, bi).Column("driver_payload", "updated_on").WherePK().Update()
// 	return err
// }

// // UpdatePaymentsDriver updates an organization's payment method type
// func (bi *BillingInfo) UpdatePaymentsDriver(ctx context.Context, paymentsDriver PaymentsDriver) error {
// 	bi.PaymentsDriver = paymentsDriver
// 	bi.UpdatedOn = time.Now()
// 	_, err := db.DB.ModelContext(ctx, bi).Column("payments_driver", "updated_on").WherePK().Update()
// 	return err
// }
