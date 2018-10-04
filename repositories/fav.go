package repositories

import (
	"github.com/OdaDaisuke/emo-lyrics-api/interfaces"
	"github.com/OdaDaisuke/emo-lyrics-api/models"
	"github.com/jinzhu/gorm"
)

type FavRepo struct {
	dbCtx *gorm.DB
}

func NewFavRepo(dbCtx *gorm.DB) *FavRepo {
	return &FavRepo{
		dbCtx: dbCtx,
	}
}

func (r *FavRepo) PostFav(params *interfaces.PostFavParams) (*models.Fav, error) {
	fav := &models.Fav{
		LyricID: params.LyricId,
		UserID:  params.UserId,
	}
	if err := r.dbCtx.Where(fav).Last(fav).Error; err != nil {
		return nil, err
	}
	return fav, nil
}

func (r *FavRepo) UnFav(params *interfaces.UnFavParams) (*models.Fav, error) {
	fav := &models.Fav{
		UserID:  params.UserId,
		LyricID: params.LyricId,
	}
	if err := r.dbCtx.Where(fav).Delete(fav).Error; err != nil {
		return nil, err
	}
	return fav, nil
}

func (r *FavRepo) GetMyFavList(params *interfaces.GetFavListParams) ([]*models.Fav, error) {
	fav := []*models.Fav{}
	where := &models.Fav{
		UserID: params.UserId,
	}

	if err := r.dbCtx.Model(fav).Where(where).Find(fav).Error; err != nil {
		return nil, err
	}
	return fav, nil
}
