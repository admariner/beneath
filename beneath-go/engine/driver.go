package engine

import (
	"github.com/beneath-core/beneath-go/core/codec"
	pb "github.com/beneath-core/beneath-go/proto"
	uuid "github.com/satori/go.uuid"
)

// StreamsDriver defines the functions necessary to encapsulate Beneath's streaming data needs
type StreamsDriver interface {
	// GetMaxMessageSize returns the maximum accepted message size in bytes
	GetMaxMessageSize() int

	// QueueWriteRequest queues a write request -- concretely, it results in
	// the write request being written to Pubsub, then from there read by
	// the data processing pipeline and written to BigTable and BigQuery
	QueueWriteRequest(req *pb.WriteRecordsRequest) error

	// ReadWriteRequests triggers fn for every WriteRecordsRequest that's written with QueueWriteRequest
	ReadWriteRequests(fn func(*pb.WriteRecordsRequest) error) error

	// QueueWriteReport publishes a batch of keys + metrics to the streams driver
	QueueWriteReport(rep *pb.WriteRecordsReport) error

	// ReadWriteReports reads messages from the Metrics topic
	ReadWriteReports(fn func(*pb.WriteRecordsReport) error) error
}

// TablesDriver defines the functions necessary to encapsulate Beneath's operational datastore needs
type TablesDriver interface {
	// GetMaxKeySize returns the maximum accepted key size in bytes
	GetMaxKeySize() int

	// GetMaxDataSize returns the maximum accepted value size in bytes
	GetMaxDataSize() int

	// WriteRecord saves a record unless sequenceNumber is lower than that of a previous write
	// to the same key
	WriteRecord(instanceID uuid.UUID, key []byte, avroData []byte, sequenceNumber int64) error

	// ReadRecords reads one or multiple (not necessarily sequential) records by key and calls fn one by one
	ReadRecords(instanceID uuid.UUID, keys [][]byte, fn func(idx uint, avroData []byte, sequenceNumber int64) error) error

	// ReadRecordRange reads one or a range of records by key and calls fn one by one
	ReadRecordRange(instanceID uuid.UUID, keyRange *codec.KeyRange, limit int, fn func(avroData []byte, sequenceNumber int64) error) error
}

// WarehouseDriver defines the functions necessary to encapsulate Beneath's data archiving needs
type WarehouseDriver interface {
	// GetMaxDataSize returns the maximum accepted row size in bytes
	GetMaxDataSize() int

	// RegisterProject should be called when a new project is created to create a corresponding dataset in the warehouse
	RegisterProject(projectID uuid.UUID, public bool, name, displayName, description string) error

	// RegisterStreamInstance should be called when a new stream instance is created to create a corresponding table in the warehouse
	RegisterStreamInstance(projectID uuid.UUID, projectName string, streamID uuid.UUID, streamName string, streamDescription string, schemaJSON string, keyFields []string, instanceID uuid.UUID) error

	// WriteRecordToWarehouse is called to write a record to the data warehouse
	WriteRecord(projectName string, streamName string, instanceID uuid.UUID, key []byte, data map[string]interface{}, sequenceNumber int64) error
}
