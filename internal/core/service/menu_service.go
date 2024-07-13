package service

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/core/models"
	"github.com/voltgizerz/POS-restaurant/internal/core/ports"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type MenuService struct {
	txRepository   ports.ITxRepository
	menuRepository ports.IMenuRepository
}

func NewMenuService(i interactor.MenuService) *MenuService {
	return &MenuService{
		txRepository:   i.TxRepository,
		menuRepository: i.MenuRepository,
	}
}

func (s *MenuService) RegisterMenu(ctx context.Context, menuData entity.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.RegisterMenu")
	defer span.Finish()

	menuConvert := models.MenuORM{
		Name:      menuData.Name,
		Thumbnail: menuData.Thumbnail,
		UserID:    menuData.UserID,
		Price:     menuData.Price,
	}

	idData, err := s.menuRepository.AddMenu(ctx, menuConvert)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": menuData.UserID,
			"error":   err,
		}).Error("[MenuService] error on AddMenu")
		return 0, err
	}

	return idData, nil
}

func (s *MenuService) GetMenu(ctx context.Context, userID int64) ([]*entity.MenuResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.GetMenu")
	defer span.Finish()

	result, err := s.menuRepository.FetchMenuByID(ctx, userID)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("[MenuService] error on FetchMenuByID")
		return nil, err
	}

	menus := []*entity.MenuResponse{}
	for _, data := range result {
		dataMenu := &entity.MenuResponse{
			ID:        data.ID,
			Name:      data.Name,
			Thumbnail: data.Thumbnail,
			Price:     data.Price,
			IsActive:  data.IsActive,
		}
		menus = append(menus, dataMenu)
	}

	return menus, nil
}

func (s *MenuService) UpdateActiveMenuID(ctx context.Context, userID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.UpdateActiveMenuID")
	defer span.Finish()

	res, err := s.menuRepository.UpdateActiveMenuByMenuID(ctx, nil, userID)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("[MenuService] error on UpdateActiveMenuByMenuID")
		return 0, err
	}

	return res, nil
}

func (s *MenuService) UpdateActiveMenuBatchUserID(ctx context.Context, userID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.UpdateActiveMenuBatchUserID")
	defer span.Finish()

	res, err := s.menuRepository.UpdateActiveMenuBatchUser(ctx, nil, userID)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("[MenuService] error on UpdateActiveMenuBatchUser")
		return 0, err
	}

	return res, nil
}

func (s *MenuService) UpdateMenuID(ctx context.Context, menuData entity.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.UpdateMenuID")
	defer span.Finish()

	menuORM := models.MenuORM{
		ID:        menuData.ID,
		Name:      menuData.Name,
		Thumbnail: menuData.Thumbnail,
		Price:     menuData.Price,
		UserID:    menuData.UserID,
		IsActive:  menuData.IsActive,
	}

	result, err := s.menuRepository.UpdateMenuByMenuID(ctx, nil, &menuORM)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"menu_id": menuData.ID,
			"error":   err,
		}).Error("[MenuService] error on UpdateMenuByMenuID")
		return 0, err
	}

	return result, nil
}
