package bootstrap

import (
	"api/common"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/weplanx/go/api"
	"github.com/weplanx/go/helper"
	"github.com/weplanx/go/passport"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var Provides = fx.Provide(
	LoadSettings,
	InitializeDatabase,
	InitializeRedis,
	InitializeCookie,
	InitializePassport,
	InitializeCipher,
	api.New,
	HttpServer,
)

// LoadSettings 初始化应用配置
func LoadSettings() (app *common.Set, err error) {
	if _, err = os.Stat("./config.yml"); os.IsNotExist(err) {
		err = errors.New("the path [./config.yml] does not have a configuration file")
		return
	}
	var b []byte
	b, err = ioutil.ReadFile("./config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &app)
	if err != nil {
		return
	}
	return
}

// InitializeDatabase 初始化MongoDB数据库
func InitializeDatabase(app *common.Set) (client *mongo.Client, db *mongo.Database, err error) {
	option := app.Database
	if client, err = mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(option.Uri),
	); err != nil {
		return
	}
	db = client.Database(option.Name)
	return
}

// InitializeRedis 初始化Redis缓存
// 配置文档 https://github.com/go-redis/redis
func InitializeRedis(app *common.Set) (client *redis.Client, err error) {
	option := app.Redis
	client = redis.NewClient(&redis.Options{
		Addr:     option.Address,
		Password: option.Password,
		DB:       option.DB,
	})
	if err = client.Ping(context.Background()).Err(); err != nil {
		return
	}
	return
}

// InitializeCookie 创建 Cookie 工具
func InitializeCookie(app *common.Set) *helper.CookieHelper {
	return helper.NewCookieHelper(app.Cookie, http.SameSiteStrictMode)
}

// InitializePassport 创建认证
func InitializePassport(app *common.Set) *passport.Passport {
	return passport.New(map[string]*passport.Auth{
		"system": {
			Key: app.Key,
			Iss: app.Name,
			Aud: []string{"admin"},
			Exp: 720,
		},
	})
}

// InitializeCipher 初始化数据加密
func InitializeCipher(app *common.Set) (*helper.CipherHelper, error) {
	return helper.NewCipherHelper(app.Key)
}

// HttpServer 启动 HTTP 服务
func HttpServer(lc fx.Lifecycle, config *common.Set) (f *fiber.App) {
	f = fiber.New()
	f.Use(logger.New())
	f.Use(recover.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.Cors, ","),
		AllowMethods:     "POST",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go f.Listen(":9000")
			return nil
		},
	})
	return
}
