package biz

import (
	"context"

	"github.com/ZQCard/kbk-oss/internal/domain"
	"github.com/go-kratos/kratos/v2/log"
)

type FileRepo interface {
	GetOssStsToken(ctx context.Context) (*domain.OssStsToken, error)
	UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error)
}

type FileUsecase struct {
	repo   FileRepo
	logger *log.Helper
}

func NewFileUsecase(repo FileRepo, logger log.Logger) *FileUsecase {
	return &FileUsecase{repo: repo, logger: log.NewHelper(log.With(logger, "module", "usecase/file"))}
}

func (uc *FileUsecase) GetOssStsToken(ctx context.Context) (*domain.OssStsToken, error) {
	return uc.repo.GetOssStsToken(ctx)
}

func (uc *FileUsecase) UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	return uc.repo.UploadFile(ctx, fileName, fileType, content)
}
