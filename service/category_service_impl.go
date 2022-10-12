package service

import (
	"context"
	"database/sql"
	"main/helper"
	"main/model/domain"
	"main/model/web"
	"main/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}


func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
  tx, err := service.DB.Begin()
  helper.PanicIfError(err)

  category := domain.Category{
    Name: request.Name,
  }

  category = service.CategoryRepository.Save(ctx, tx, category)

  return helper.CategoryToResponse(category)
  
  defer helper.CommitOrRollback(tx)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
  tx, err := service.DB.Begin()
  helper.PanicIfError(err)

  defer helper.CommitOrRollback(tx)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
  tx, err := service.DB.Begin()
  helper.PanicIfError(err)

  defer helper.CommitOrRollback(tx)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
  tx, err := service.DB.Begin()
  helper.PanicIfError(err)

  defer helper.CommitOrRollback(tx)
}

func (service *CategoryServiceImpl) FindByAll(ctx context.Context) []web.CategoryResponse {
  tx, err := service.DB.Begin()
  helper.PanicIfError(err)

  defer helper.CommitOrRollback(tx)
}

