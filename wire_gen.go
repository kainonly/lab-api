// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"api/app"
	"api/app/pages"
	"api/app/pictures"
	"api/app/system"
	"api/app/users"
	"api/bootstrap"
	"api/common"
	"github.com/gin-gonic/gin"
	"github.com/weplanx/go/engine"
)

// Injectors from wire.go:

func App(value *common.Values) (*gin.Engine, error) {
	client, err := bootstrap.UseMongoDB(value)
	if err != nil {
		return nil, err
	}
	database := bootstrap.UseDatabase(client, value)
	redisClient, err := bootstrap.UseRedis(value)
	if err != nil {
		return nil, err
	}
	conn, err := bootstrap.UseNats(value)
	if err != nil {
		return nil, err
	}
	jetStreamContext, err := bootstrap.UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	openAPI := bootstrap.UseOpenapi(value)
	cipher, err := bootstrap.UseCipher(value)
	if err != nil {
		return nil, err
	}
	hid, err := bootstrap.UseHID(value)
	if err != nil {
		return nil, err
	}
	cosClient, err := bootstrap.UseCos(value)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		Values:      value,
		MongoClient: client,
		Db:          database,
		Redis:       redisClient,
		Nats:        conn,
		Js:          jetStreamContext,
		Open:        openAPI,
		Cipher:      cipher,
		HID:         hid,
		Cos:         cosClient,
	}
	service := &system.Service{
		Inject: inject,
	}
	passport := bootstrap.UsePassport(value)
	transfer, err := bootstrap.UseTransfer(value, jetStreamContext)
	if err != nil {
		return nil, err
	}
	middleware := &system.Middleware{
		Service:  service,
		Passport: passport,
		Transfer: transfer,
	}
	usersService := &users.Service{
		Inject: inject,
	}
	pagesService := &pages.Service{
		Inject: inject,
	}
	controller := &system.Controller{
		Service:  service,
		Users:    usersService,
		Pages:    pagesService,
		Passport: passport,
	}
	engineEngine := bootstrap.UseEngine(value, jetStreamContext)
	engineService := &engine.Service{
		Engine: engineEngine,
		Db:     database,
	}
	engineController := &engine.Controller{
		Engine:  engineEngine,
		Service: engineService,
	}
	pagesController := &pages.Controller{
		Service: pagesService,
	}
	picturesService := &pictures.Service{
		Inject: inject,
	}
	picturesController := &pictures.Controller{
		Service: picturesService,
	}
	ginEngine := app.New(value, middleware, controller, engineController, pagesController, picturesController)
	return ginEngine, nil
}
