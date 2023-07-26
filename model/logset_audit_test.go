package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestCreateAccessLogCollection(t *testing.T) {
	ctx := context.TODO()
	err := x.Db.Collection("logset_audit").Drop(ctx)
	assert.NoError(t, err)
	err = x.Db.Collection("logset_audit_fail").Drop(ctx)
	assert.NoError(t, err)
	option := options.CreateCollection().
		SetTimeSeriesOptions(
			options.TimeSeries().
				SetTimeField("timestamp").
				SetMetaField("metadata"),
		).
		SetExpireAfterSeconds(15552000)
	err = x.Db.CreateCollection(ctx, "logset_audit", option)
	assert.NoError(t, err)
	err = x.Db.CreateCollection(ctx, "logset_audit_fail", option)
	assert.NoError(t, err)
}
