package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	ossV1 "github.com/ZQCard/kbk-oss/api/oss/v1"
	"github.com/ZQCard/kbk-oss/internal/biz"
)

type FileService struct {
	ossV1.UnimplementedOSSServer
	fileUsecase *biz.FileUsecase
	log         *log.Helper
}

func NewFileService(fileUsecase *biz.FileUsecase, logger log.Logger) *FileService {
	return &FileService{
		log:         log.NewHelper(log.With(logger, "module", "service/FileService")),
		fileUsecase: fileUsecase,
	}
}

func (s *FileService) GetOssStsToken(ctx context.Context, req *emptypb.Empty) (*ossV1.OssStsTokenResponse, error) {
	stsResponse, err := s.fileUsecase.GetOssStsToken(ctx)
	if err != nil {
		return nil, err
	}
	response := &ossV1.OssStsTokenResponse{}
	response.AccessKey = stsResponse.AccessKey
	response.AccessSecret = stsResponse.AccessSecret
	response.Expiration = stsResponse.Expiration
	response.SecurityToken = stsResponse.SecurityToken
	response.EndPoint = stsResponse.EndPoint
	response.BucketName = stsResponse.BucketName
	response.Region = stsResponse.Region
	response.Url = stsResponse.Url
	return response, nil
}

func (s *FileService) UploadFile(ctx context.Context, req *ossV1.UploadFileRequest) (*ossV1.UploadFileResponse, error) {
	url, err := s.fileUsecase.UploadFile(ctx, req.FileName, req.FileType, req.Content)
	if err != nil {
		return nil, err
	}
	response := &ossV1.UploadFileResponse{}
	response.Url = url
	return response, err
}
