package data

import (
	"bytes"
	"context"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/log"

	ossV1 "github.com/ZQCard/kbk-oss/api/oss/v1"
	"github.com/ZQCard/kbk-oss/internal/biz"
	"github.com/ZQCard/kbk-oss/internal/domain"
)

type FileRepo struct {
	data *Data
	log  *log.Helper
}

func NewFileRepo(data *Data, logger log.Logger) biz.FileRepo {
	r := &FileRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/file")),
	}
	return r
}

func (f FileRepo) UploadFile(ctx context.Context, fileName string, fileType string, content []byte) (string, error) {
	client := f.data.oss

	bucket, err := client.Bucket(f.data.cfg.Oss.BucketName)
	if err != nil {
		return "", ossV1.ErrorSystemError("初始化失败").WithCause(err)
	}
	path := "uploadFile/" + strings.Trim(fileType, ".") + "/" + fileName + fileType
	err = bucket.PutObject(path, bytes.NewReader(content))
	if err != nil {
		return "", ossV1.ErrorSystemError("文件上传失败").WithCause(err)
	}
	url := "https://" + f.data.cfg.Oss.BucketName + "." + f.data.cfg.Oss.EndPoint + "/" + path
	return url, nil
}

func (f FileRepo) GetOssStsToken(ctx context.Context) (*domain.OssStsToken, error) {
	config := &openapi.Config{
		AccessKeyId:     &f.data.cfg.Oss.AccessKey,
		AccessKeySecret: &f.data.cfg.Oss.AccessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("sts.cn-hangzhou.aliyuncs.com")
	client, err := sts.NewClient(config)
	if err != nil {
		return &domain.OssStsToken{}, ossV1.ErrorSystemError("初始化失败").WithCause(err)
	}
	assumeRoleRequest := &sts.AssumeRoleRequest{
		RoleArn:         &f.data.cfg.Oss.StsRoleArn,
		RoleSessionName: &f.data.cfg.Oss.SessionName,
	}
	resp, err := client.AssumeRole(assumeRoleRequest)
	if err != nil {
		return &domain.OssStsToken{}, ossV1.ErrorSystemError("角色授权失败").WithCause(err)
	}

	response := &domain.OssStsToken{}
	response.AccessKey = *resp.Body.Credentials.AccessKeyId
	response.AccessSecret = *resp.Body.Credentials.AccessKeySecret
	response.Expiration = *resp.Body.Credentials.Expiration
	response.SecurityToken = *resp.Body.Credentials.SecurityToken
	response.EndPoint = f.data.cfg.Oss.EndPoint
	response.BucketName = f.data.cfg.Oss.BucketName
	response.Region = f.data.cfg.Oss.Region
	response.Url = "https://" + f.data.cfg.Oss.BucketName + "." + f.data.cfg.Oss.EndPoint + "/"
	return response, nil
}
