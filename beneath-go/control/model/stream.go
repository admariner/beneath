package model

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-pg/pg"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"

	"github.com/beneath-core/beneath-go/control/db"
	"github.com/beneath-core/beneath-go/core/schema"
)

// Stream represents a collection of data
type Stream struct {
	StreamID                uuid.UUID `sql:",pk,type:uuid,default:uuid_generate_v4()"`
	Name                    string    `sql:",notnull",validate:"required,gte=1,lte=40"` // not unique because of (project_id, user_id) index
	Description             string    `validate:"omitempty,lte=255"`
	Schema                  string    `sql:",notnull",validate:"required"`
	AvroSchema              string    `sql:",type:json,notnull",validate:"required"`
	CanonicalAvroSchema     string    `sql:",type:json,notnull",validate:"required"`
	KeyFields               []string  `sql:",notnull",validate:"required,gte=1"`
	External                bool      `sql:",notnull"`
	Batch                   bool      `sql:",notnull"`
	Manual                  bool      `sql:",notnull"`
	ProjectID               uuid.UUID `sql:"on_delete:RESTRICT,notnull,type:uuid"`
	Project                 *Project
	StreamInstances         []*StreamInstance
	CurrentStreamInstanceID *uuid.UUID `sql:"on_delete:SET NULL,type:uuid"`
	CurrentStreamInstance   *StreamInstance
	CreatedOn               time.Time `sql:",default:now()"`
	UpdatedOn               time.Time `sql:",default:now()"`
}

var (
	// used for validation
	streamNameRegex *regexp.Regexp
)

func init() {
	// configure validation
	streamNameRegex = regexp.MustCompile("^[_a-z][_\\-a-z0-9]*$")
	GetValidator().RegisterStructValidation(streamValidation, Stream{})
}

// custom stream validation
func streamValidation(sl validator.StructLevel) {
	s := sl.Current().Interface().(Stream)

	if !streamNameRegex.MatchString(s.Name) {
		sl.ReportError(s.Name, "Name", "", "alphanumericorunderscore", "")
	}
}

// FindStream finds a stream
func FindStream(streamID uuid.UUID) *Stream {
	stream := &Stream{
		StreamID: streamID,
	}
	err := db.DB.Model(stream).WherePK().Column("stream.*", "Project", "CurrentStreamInstance").Select()
	if !AssertFoundOne(err) {
		return nil
	}
	return stream
}

// FindStreamByNameAndProject finds a stream
func FindStreamByNameAndProject(name string, projectName string) *Stream {
	stream := &Stream{}
	err := db.DB.Model(stream).
		Column("stream.*", "Project", "CurrentStreamInstance").
		Where("lower(stream.name) = lower(?)", name).
		Where("lower(project.name) = lower(?)", projectName).
		Select()
	if !AssertFoundOne(err) {
		return nil
	}
	return stream
}

// FindInstanceIDByNameAndProject returns the current instance ID of the stream
func FindInstanceIDByNameAndProject(name string, projectName string) uuid.UUID {
	return getInstanceCache().get(name, projectName)
}

// FindCachedStreamByCurrentInstanceID returns select info about the instance's stream
func FindCachedStreamByCurrentInstanceID(instanceID uuid.UUID) *CachedStream {
	return getStreamCache().get(instanceID)
}

// UpdateDetails updates a stream (only exposes fields where updates are permitted)
func (s *Stream) UpdateDetails(description *string, manual *bool) error {
	if description != nil {
		s.Description = *description
	}
	if manual != nil {
		s.Manual = *manual
	}

	// validate
	err := GetValidator().Struct(s)
	if err != nil {
		return err
	}

	// update
	_, err = db.DB.Model(s).Column("description", "manual").WherePK().Update()
	return err
}

// CompileAndCreate compiles the schema, derives name and avro schemas and inserts
// the stream into the database
func (s *Stream) CompileAndCreate() error {
	// compile schema
	compiler := schema.NewCompiler(s.Schema)
	err := compiler.Compile()
	if err != nil {
		return fmt.Errorf("Error compiling schema: %s", err.Error())
	}
	streamDef := compiler.GetStream()

	// get avro schemas
	avro, err := streamDef.BuildAvroSchema()
	if err != nil {
		return fmt.Errorf("Error compiling schema: %s", err.Error())
	}

	canonicalAvro, err := streamDef.BuildCanonicalAvroSchema()
	if err != nil {
		return fmt.Errorf("Error compiling schema: %s", err.Error())
	}

	// set missing stream fields
	s.Name = streamDef.Name
	s.KeyFields = streamDef.KeyFields
	s.AvroSchema = avro
	s.CanonicalAvroSchema = canonicalAvro

	// validate
	err = GetValidator().Struct(s)
	if err != nil {
		return err
	}

	// create stream (and a new stream instance ID if not batch)
	err = db.DB.RunInTransaction(func(tx *pg.Tx) error {
		// insert stream
		_, err := tx.Model(s).Insert()
		if err != nil {
			return err
		}

		// create and set stream instance if not batch
		if !s.Batch {
			// create stream instance
			si, err := CreateStreamInstanceWithTx(tx, s.StreamID)
			if err != nil {
				return err
			}

			// update stream with stream instance ID
			s.CurrentStreamInstanceID = &si.StreamInstanceID
			_, err = tx.Model(s).WherePK().Update()
			if err != nil {
				return err
			}
		}

		// done
		return nil
	})
	if err != nil {
		return err
	}

	// done
	return nil
}

/**
 * TODO: In the future
 * - belongs to model
 * - dependencies (models)
 */
