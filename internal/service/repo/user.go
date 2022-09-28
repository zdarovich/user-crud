package repo

import (
	"binance-order-matcher/internal/model"
	"binance-order-matcher/internal/service"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	db *gorm.DB
}

func (u *userRepo) Save(ctx context.Context, user *model.User) error {
	return u.db.Create(user).Error
}
func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	return u.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"country": user.Country}),
	}).Create(&user).Error
}

func (u *userRepo) Get(ctx context.Context, page, limit int, filter model.User) ([]*model.User, error) {
	var result []*model.User
	offset := page * limit

	tx := u.db.Where(filter).Limit(limit).Offset(offset).Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (u *userRepo) Delete(ctx context.Context, user *model.User) error {
	return u.db.Delete(user).Error
}

func NewUserRepo(db *gorm.DB) service.UserRepo {
	return &userRepo{db}
}
