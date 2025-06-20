package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/middleware/auth"
)

type TenantService struct {
	adminV1.TenantServiceHTTPServer

	log *log.Helper

	repo *data.TenantRepo
}

func NewTenantService(logger log.Logger, repo *data.TenantRepo) *TenantService {
	l := log.NewHelper(log.With(logger, "module", "tenant/service/admin-service"))
	return &TenantService{
		log:  l,
		repo: repo,
	}
}

func (s *TenantService) List(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListTenantResponse, error) {
	return s.repo.List(ctx, req)
}

func (s *TenantService) Get(ctx context.Context, req *userV1.GetTenantRequest) (*userV1.Tenant, error) {
	return s.repo.Get(ctx, req)
}

func (s *TenantService) Create(ctx context.Context, req *userV1.CreateTenantRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.CreateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TenantService) Update(ctx context.Context, req *userV1.UpdateTenantRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	req.Data.UpdateBy = trans.Ptr(operator.UserId)

	if err = s.repo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *TenantService) Delete(ctx context.Context, req *userV1.DeleteTenantRequest) (*emptypb.Empty, error) {
	if err := s.repo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
