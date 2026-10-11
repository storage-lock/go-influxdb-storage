package influxdb_storage

import (
	"context"
	"time"

	go_iterator "github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
)

// InfluxDBStorage InfluxDB存储实现
// InfluxDB 是时序数据库，不支持事务和原子条件写入（CAS）
// 因此无法安全地用于分布式锁，本实现仅作为非 CAS 存储的参考示例
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
	return []storage.StorageCapability{
		storage.CapabilityReliableTime,
		storage.CapabilityWatch,
	}
}

func (x *InfluxDBStorage) Watch(ctx context.Context, lockId string) (<-chan *storage.LockInformation, error) {
	return storage.NewPollingLockWatcher(x, 500*time.Millisecond).Watch(ctx, lockId)
}

func (x *InfluxDBStorage) Init(ctx context.Context) error {
	return nil
}

func (x *InfluxDBStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	return storage_lock.ErrStorageCapabilityMissing
}

func (x *InfluxDBStorage) UpdateWithVersion(ctx context.Context, lockId string, expectedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	return storage_lock.ErrStorageCapabilityMissing
}

func (x *InfluxDBStorage) DeleteWithVersion(ctx context.Context, lockId string, expectedVersion storage.Version, lockInformation *storage.LockInformation) error {
	return storage_lock.ErrStorageCapabilityMissing
}

func (x *InfluxDBStorage) Get(ctx context.Context, lockId string) (string, error) {
	return "", storage_lock.ErrLockNotFound
}

func (x *InfluxDBStorage) GetTime(ctx context.Context) (time.Time, error) {
	return time.Now(), nil
}

func (x *InfluxDBStorage) Close(ctx context.Context) error {
	return nil
}

func (x *InfluxDBStorage) List(ctx context.Context) (go_iterator.Iterator[*storage.LockInformation], error) {
	return go_iterator.FromSlice([]*storage.LockInformation{}), nil
}