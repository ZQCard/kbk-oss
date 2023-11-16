package data

import (
	"context"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"repo.example.com/oss/internal/conf"
	"repo.example.com/oss/pkg/middleware/requestInfo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRedisCmd,
	NewMysqlCmd,
	NewRedisClient,
	NewFileRepo,
	NewOssClient,
)

// Data .
type Data struct {
	cfg    *conf.Bootstrap
	logger *log.Helper
	db     *gorm.DB
	rdb    *redis.Client
	oss    *OssClient
}

func NewData(cfg *conf.Bootstrap, db *gorm.DB, redisCli *redis.Client, logger log.Logger, oss *OssClient) (*Data, func(), error) {
	logs := log.NewHelper(log.With(logger, "module", "charging-station-file/data"))
	cleanup := func() {
		logs.Info("closing the data resources")
	}

	return &Data{
		logger: logs,
		cfg:    cfg,
		db:     db,
		rdb:    redisCli,
		oss:    oss,
	}, cleanup, nil
}

func getDomain(ctx context.Context) string {
	domain := ctx.Value(requestInfo.DomainKey)
	return domain.(string)
}

func getDbWithDomain(ctx context.Context, db *gorm.DB) *gorm.DB {
	domain := ctx.Value(requestInfo.DomainKey)
	if domain != nil {
		db = db.Where("domain = ?", domain)
	}
	return db
}

func NewRedisClient(conf *conf.Data) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
		DB:           int(conf.Redis.Db),
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	logs := log.NewHelper(log.With(logger, "module", "serviceName/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	err := client.Ping().Err()
	if err != nil {
		// redis非强制连接
		logs.Warnf("redis connect error: %v", err)
	}
	return client
}

func NewMysqlCmd(conf *conf.Bootstrap, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "serviceName/data/mysql"))
	db, err := gorm.Open(mysql.Open(conf.Data.Database.Source), &gorm.Config{})
	if err != nil {
		// 数据库非强制连接
		logs.Warnf("mysql connect error: %v", err)
	}
	// 如果是开发环境 打印sql
	if conf.Env.Mode == "dev" {
		db = db.Debug()
	}
	return db
}

type OssClient struct {
	*oss.Client
}

func NewOssClient(conf *conf.Bootstrap) (*OssClient, error) {
	client, err := oss.New(conf.Oss.EndPoint, conf.Oss.AccessKey, conf.Oss.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &OssClient{client}, nil
}
