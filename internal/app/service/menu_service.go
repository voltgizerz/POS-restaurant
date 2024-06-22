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
	authService    ports.IAuth
	menuRepository ports.IMenuRepository
}

func NewMenuService(i interactor.MenuService) *MenuService {
	return &MenuService{
		authService:    i.AuthService,
		menuRepository: i.MenuRepository,
	}
}

func (s *MenuService) RegisterMenu(ctx context.Context, menuData entity.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.RegisterMenu")
	defer span.Finish()

	menuConvert := entity.MenuOrm{
		Name:      menuData.Name,
		Thumbnail: menuData.Thumbnail,
		UserId:    menuData.UserId,
		Price:     menuData.Price,
	}

	idData, err := s.menuRepository.AddMenu(ctx, &menuConvert)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": menuData.UserId,
			"error":   err,
		}).Error("[MenuService] error on AddMenu")
		return 0, err
	}

	return idData, nil
}

func (s *MenuService) GetMenu(ctx context.Context, idUser int64) ([]*entity.MenuResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.GetMenu")
	defer span.Finish()

	result, err := s.menuRepository.FetchMenuById(ctx, idUser)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": idUser,
			"error":   err,
		}).Error("[MenuService] error on GetMenu")
		return nil, err
	}

	convert_menu := []*entity.MenuResponse{}
	for _, data := range result {
		menu_data := entity.MenuResponse{
			ID:        data.ID,
			Name:      data.Name,
			Thumbnail: data.Thumbnail,
			Price:     data.Price,
			IsActive:  data.IsActive,
		}
		convert_menu = append(convert_menu, &menu_data)
	}

	return convert_menu, nil
}

func (s *MenuService) DeleteMenuID(ctx context.Context, idUser int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.DeleteMenuID")
	defer span.Finish()

	_, err := s.menuRepository.DeleteMenuByMenuID(ctx, idUser)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": idUser,
			"error":   err,
		}).Error("[MenuService] error on DeleteMenuID")
		return 0, err
	}

	return 1, nil
}

func (s *MenuService) DeleteMenuBatchUserID(ctx context.Context, idUser int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.DeleteMenuBatchUserID")
	defer span.Finish()

	_, err := s.menuRepository.DeleteMenuBatchUser(ctx, idUser)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"user_id": idUser,
			"error":   err,
		}).Error("[MenuService] error on DeleteMenBatchUserID")
		return 0, err
	}

	return 1, nil
}

func (s *MenuService) UpdateMenuID(ctx context.Context, menuData *entity.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.MenuService.UpdateMenuID")
	defer span.Finish()

	menuOrm := entity.MenuOrm{
		ID:        menuData.ID,
		Name:      menuData.Name,
		Thumbnail: menuData.Thumbnail,
		Price:     menuData.Price,
		UserId:    menuData.UserId,
		IsActive:  menuData.IsActive,
	}

	result, err := s.menuRepository.UpdateMenuByMenuID(ctx, &menuOrm)
	if err != nil {
		logger.LogStdErr.WithFields(logrus.Fields{
			"menu_id": menuData.ID,
			"error":   err,
		}).Error("[MenuService] error on UpdateMenuID")
		return 0, err
	}

	return result, nil
}
