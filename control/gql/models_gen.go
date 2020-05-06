// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"time"

	"github.com/satori/go.uuid"
	"gitlab.com/beneath-hq/beneath/control/entity"
)

type Organization interface {
	IsOrganization()
}

type CreateModelInput struct {
	ProjectID           uuid.UUID   `json:"projectID"`
	Name                string      `json:"name"`
	Kind                string      `json:"kind"`
	SourceURL           *string     `json:"sourceURL"`
	Description         *string     `json:"description"`
	InputStreamIDs      []uuid.UUID `json:"inputStreamIDs"`
	OutputStreamSchemas []string    `json:"outputStreamSchemas"`
	ReadQuota           int         `json:"readQuota"`
	WriteQuota          int         `json:"writeQuota"`
}

type Metrics struct {
	EntityID     uuid.UUID `json:"entityID"`
	Period       string    `json:"period"`
	Time         time.Time `json:"time"`
	ReadOps      int       `json:"readOps"`
	ReadBytes    int       `json:"readBytes"`
	ReadRecords  int       `json:"readRecords"`
	WriteOps     int       `json:"writeOps"`
	WriteBytes   int       `json:"writeBytes"`
	WriteRecords int       `json:"writeRecords"`
}

type NewServiceSecret struct {
	Secret *entity.ServiceSecret `json:"secret"`
	Token  string                `json:"token"`
}

type NewUserSecret struct {
	Secret *entity.UserSecret `json:"secret"`
	Token  string             `json:"token"`
}

type PrivateOrganization struct {
	OrganizationID string                                `json:"organizationID"`
	Name           string                                `json:"name"`
	DisplayName    string                                `json:"displayName"`
	Description    *string                               `json:"description"`
	PhotoURL       *string                               `json:"photoURL"`
	CreatedOn      time.Time                             `json:"createdOn"`
	UpdatedOn      time.Time                             `json:"updatedOn"`
	ReadQuota      *int                                  `json:"readQuota"`
	WriteQuota     *int                                  `json:"writeQuota"`
	ReadUsage      int                                   `json:"readUsage"`
	WriteUsage     int                                   `json:"writeUsage"`
	Projects       []*entity.Project                     `json:"projects"`
	Services       []*entity.Service                     `json:"services"`
	PersonalUserID *uuid.UUID                            `json:"personalUserID"`
	PersonalUser   *entity.User                          `json:"personalUser"`
	Permissions    *entity.PermissionsUsersOrganizations `json:"permissions"`
}

func (PrivateOrganization) IsOrganization() {}

type UpdateModelInput struct {
	ModelID             uuid.UUID   `json:"modelID"`
	SourceURL           *string     `json:"sourceURL"`
	Description         *string     `json:"description"`
	InputStreamIDs      []uuid.UUID `json:"inputStreamIDs"`
	OutputStreamSchemas []string    `json:"outputStreamSchemas"`
	ReadQuota           *int        `json:"readQuota"`
	WriteQuota          *int        `json:"writeQuota"`
}
