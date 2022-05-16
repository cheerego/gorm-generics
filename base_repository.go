package gorm_generics

import (
	"context"
	"gorm.io/gorm"
)

const key = "ctx-gorm"

func FromContext(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	db, ok := ctx.Value(key).(*gorm.DB)
	if ok {
		return db
	}
	return defaultDB
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return BaseRepository[T]{db: db}
}

func (b *BaseRepository[T]) FindById(ctx context.Context, id int32) (*T, error) {
	var m T
	err := FromContext(ctx, b.db).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (b *BaseRepository[T]) FindByIdWithDeleted(ctx context.Context, id int32) (*T, error) {
	var m T
	err := FromContext(ctx, b.db).Unscoped().First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (b *BaseRepository[T]) FindByIds(ctx context.Context, id ...int32) ([]*T, error) {
	m := make([]*T, 0)

	err := FromContext(ctx, b.db).Find(&m, id).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (b *BaseRepository[T]) FindByIdsWithDeleted(ctx context.Context, id ...int32) ([]*T, error) {
	m := make([]*T, 0)
	err := FromContext(ctx, b.db).Unscoped().Find(&m, id).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (b *BaseRepository[T]) Insert(ctx context.Context, m *T) error {
	err := FromContext(ctx, b.db).Create(m).Error
	return err
}

func (b *BaseRepository[T]) BatchInsert(ctx context.Context, m []*T) (int64, error) {
	tx := FromContext(ctx, b.db).Create(m)
	return tx.RowsAffected, tx.Error
}

func (b *BaseRepository[T]) DeleteById(ctx context.Context, id int32) (int64, error) {
	var m T
	tx := FromContext(ctx, b.db).Delete(&m, id)
	return tx.RowsAffected, tx.Error
}

func (b *BaseRepository[T]) DeleteByIds(ctx context.Context, ids int32) (int64, error) {
	var m T
	tx := FromContext(ctx, b.db).Delete(&m, ids)
	return tx.RowsAffected, tx.Error
}

func (b *BaseRepository[T]) ForceDeleteById(ctx context.Context, id int32) (int64, error) {
	var m T

	tx := FromContext(ctx, b.db).Unscoped().Delete(&m, id)
	return tx.RowsAffected, tx.Error
}

func (b *BaseRepository[T]) ForceDeleteByIds(ctx context.Context, ids ...int32) (int64, error) {
	var m T
	tx := FromContext(ctx, b.db).Unscoped().Delete(&m, ids)
	return tx.RowsAffected, tx.Error
}
