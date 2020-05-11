package entity

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/cache/v7"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/beneath-hq/beneath/internal/hub"
)

// FindInstanceIDByOrganizationProjectAndName returns the current instance ID of the stream
func FindInstanceIDByOrganizationProjectAndName(ctx context.Context, organizationName string, projectName string, streamName string) uuid.UUID {
	return getInstanceCache().get(ctx, organizationName, projectName, streamName)
}

// instanceCache is a Redis and LRU based cache mapping (projectName, streamName) pairs to
// the current instance ID for that stream (table `streams`, column `current_stream_instance_id`)
type instanceCache struct {
	codec *cache.Codec
}

var (
	_instanceCache instanceCache
)

// getInstanceCache returns a global instanceCache
func getInstanceCache() instanceCache {
	if _instanceCache.codec == nil {
		_instanceCache.codec = &cache.Codec{
			Redis:     hub.Redis,
			Marshal:   _instanceCache.marshal,
			Unmarshal: _instanceCache.unmarshal,
		}
		_instanceCache.codec.UseLocalCache(_instanceCache.cacheLRUSize(), _instanceCache.cacheLRUTime())
	}

	return _instanceCache
}

func (c instanceCache) get(ctx context.Context, organizationName string, projectName string, streamName string) uuid.UUID {
	var instanceID uuid.UUID
	err := c.codec.Once(&cache.Item{
		Key:        c.redisKey(organizationName, projectName, streamName),
		Object:     &instanceID,
		Expiration: c.cacheTime(),
		Func:       c.getterFunc(ctx, organizationName, projectName, streamName),
	})

	if err != nil {
		panic(err)
	}

	return instanceID
}

func (c instanceCache) clear(ctx context.Context, organizationName string, projectName string, streamName string) {
	err := c.codec.Delete(c.redisKey(organizationName, projectName, streamName))
	if err != nil && err != cache.ErrCacheMiss {
		panic(err)
	}
}

func (c instanceCache) cacheTime() time.Duration {
	return time.Hour
}

func (c instanceCache) cacheLRUSize() int {
	return 10000
}

func (c instanceCache) cacheLRUTime() time.Duration {
	return 10 * time.Second
}

func (c instanceCache) redisKey(organizationName string, projectName string, streamName string) string {
	return fmt.Sprintf("inst:%s:%s:%s", organizationName, projectName, streamName)
}

func (c instanceCache) marshal(v interface{}) ([]byte, error) {
	instanceID := v.(uuid.UUID)
	return instanceID.Bytes(), nil
}

func (c instanceCache) unmarshal(b []byte, v interface{}) (err error) {
	instanceID := v.(*uuid.UUID)
	*instanceID, err = uuid.FromBytes(b)
	return err
}

func (c instanceCache) getterFunc(ctx context.Context, organizationName string, projectName string, streamName string) func() (interface{}, error) {
	return func() (interface{}, error) {
		res := uuid.Nil
		_, err := hub.DB.QueryContext(ctx, pg.Scan(&res), `
			select s.primary_stream_instance_id
			from streams s
			join projects p on s.project_id = p.project_id
			join organizations o on p.organization_id = o.organization_id
			where lower(s.name) = lower(?)
			and lower(p.name) = lower(?)
			and lower(o.name) = lower(?)
			and primary_stream_instance_id is not null
		`, streamName, projectName, organizationName)
		if err == pg.ErrNoRows {
			return res, nil
		}
		return res, err
	}
}
