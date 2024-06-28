package service

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type MenuService struct {
	menuRepository ports.IMenuRepository
}

func NewMenuService(i interactor.MenuService) *MenuService {
	return &MenuService{
		menuRepository: i.MenuRepository,
	}
}

func (s *MenuService) RegisterMenu(ctx context.Context, menuData entity.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.RegisterMenu")
	defer span.Finish()

	menuConvert := entity.MenuORM{
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

	result, err := s.menuRepository.FetchMenuById(ctx, userID)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("[MenuService] error on GetMenu")
		return nil, err
	}

	convertMenu := []*entity.MenuResponse{}
	for _, data := range result {
		menuData := entity.MenuResponse{
			ID:        data.ID,
			Name:      data.Name,
			Thumbnail: data.Thumbnail,
			Price:     data.Price,
			IsActive:  data.IsActive,
		}
		convertMenu = append(convertMenu, &menuData)
	}

	return convertMenu, nil
}

func (s *MenuService) UpdateActiveMenuID(ctx context.Context, userID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.UpdateActiveMenuID")
	defer span.Finish()

	_, err := s.menuRepository.UpdateActiveMenuByMenuID(ctx, userID)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("[MenuService] error on UpdateActiveMenuByMenuID")
		return 0, err
	}

	return 1, nil
}

func (s *MenuService) UpdateActiveMenuBatchUserID(ctx context.Context, userID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.UpdateActiveMenuBatchUserID")
	defer span.Finish()

	_, err := s.menuRepository.UpdateActiveMenuBatchUser(ctx, userID)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("[MenuService] error on UpdateActiveMenuBatchUser")
		return 0, err
	}

	return 1, nil
}

func (s *MenuService) UpdateMenuID(ctx context.Context, menuData entity.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.UpdateMenuID")
	defer span.Finish()

	menuOrm := entity.MenuORM{
		ID:        menuData.ID,
		Name:      menuData.Name,
		Thumbnail: menuData.Thumbnail,
		Price:     menuData.Price,
		UserID:    menuData.UserID,
		IsActive:  menuData.IsActive,
	}

	result, err := s.menuRepository.UpdateMenuByMenuID(ctx, &menuOrm)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"menu_id": menuData.ID,
			"error":   err,
		}).Error("[MenuService] error on UpdateMenuByMenuID")
		return 0, err
	}

	return result, nil
}
