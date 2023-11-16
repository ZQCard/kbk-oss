package domain

import (
	"github.com/jinzhu/copier"
	ossV1 "repo.example.com/oss/api/oss/v1"
)

type OssStsToken struct {
	AccessKey     string
	AccessSecret  string
	SecurityToken string
	Expiration    string
	BucketName    string
	Region        string
	Url           string
	EndPoint      string
}

// ToPb 将domain结构体转换为pb结构体
func (oss *OssStsToken) ToPb() *ossV1.OssStsTokenResponse {
	pb := &ossV1.OssStsTokenResponse{}
	copier.Copy(pb, oss)
	return pb
}

// ToDomain 将pb结构体转换为domain结构体
func (oss *OssStsToken) ToDomain(data interface{}) *OssStsToken {
	copier.Copy(oss, data)
	return oss
}

// ToDomainOssStsToken 将pb结构体转换为domain包下的OssStsToken结构体
func ToDomainOssStsToken(data interface{}) *OssStsToken {
	oss := &OssStsToken{}
	copier.Copy(oss, data)
	return oss
}
