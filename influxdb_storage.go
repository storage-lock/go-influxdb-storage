package influxdb_storage

// InfluxDB 是时序数据库，不支持事务和唯一约束
// 不适合用于分布式锁，仅作为实验性实现

import (
	"context"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
)

// InfluxDBStorage InfluxDB存储实现（实验性）
// ⚠️ InfluxDB 不支持CAS，不建议用于分布式锁
type InfluxDBStorage struct {
}

var _ storage.Storage = &InfluxDBStorage{}

func NewInfluxDBStorage(ctx context.Context) (*InfluxDBStorage, error) {
	return &InfluxDBStorage{}, nil
}

const StorageName = "influxdb-storage"

func (x *InfluxDBStorage) GetName() string {
	return StorageName
}

func (x *InfluxDBStorage) Capabilities() []storage.StorageCapability {
	// InfluxDB 不支持CAS
	return []storage.StorageCapability{
		storage.CapabilityReliableTime,
	}
}

func (x *InfluxDBStorage) Init(ctx context.Context) error {
	return nil
}

func (x *InfluxDBStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	return nil
}

func (x *InfluxDBStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	return nil
}

func (x *InfluxDBStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	return nil
}

func (x *InfluxDBStorage) Get(ctx context.Context, lockId string) (string, error) {
	return "", nil
}

func (x *InfluxDBStorage) GetTime(ctx context.Context) (time.Time, error) {
	return time.Now(), nil
}

func (x *InfluxDBStorage) Close(ctx context.Context) error {
	return nil
}

func (x *InfluxDBStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return nil, nil
}
