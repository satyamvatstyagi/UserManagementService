package mtnapm

import (
	"context"
	"fmt"
	"net/http"

	"github.com/satyamvatstyagi/UserManagementService/pkg/common/env"
	"go.elastic.co/apm/v2"
)

type APM struct {
	ctx  context.Context
	span *apm.Span
}

func initServiceIntrumentation(ctx context.Context, serviceName string, serviceType string) *APM {
	instrumentObj := &APM{}
	instrumentObj.span, instrumentObj.ctx = apm.StartSpan(ctx, serviceName, serviceType)
	return instrumentObj
}

func (i *APM) SetLabel(key string, value interface{}) *APM {
	i.span.Context.SetLabel(key, value)
	return i
}

func (i *APM) SetHTTPRequest(req *http.Request) *APM {
	i.span.Context.SetHTTPRequest(req)
	return i
}

func (i *APM) SetName(name string) *APM {
	i.span.Name = name
	return i
}

func (i *APM) GetSpan() *apm.Span {
	return i.span
}

func (i *APM) SetServiceTarget(name, serviceType string) *APM {
	i.span.Context.SetServiceTarget(apm.ServiceTargetSpanContext{Name: name, Type: serviceType})
	return i
}

func (i *APM) GetContext() context.Context {
	return i.ctx
}

func (i *APM) SetDatabaseName(dbName string, queryStmt string, dbType string, user string) *APM {
	i.span.Context.SetDatabase(apm.DatabaseSpanContext{Statement: queryStmt, Type: dbType, User: user, Instance: dbName})
	return i
}

func (i *APM) SetAction(action string) *APM {
	i.span.Action = action
	return i
}
func (i *APM) SetTypeAndSubType(typeName string, subType string) *APM {
	i.span.Type = typeName
	i.span.Subtype = subType
	return i
}

func (i *APM) SetDestinationAddress(addr string, port int) *APM {
	i.span.Context.SetDestinationAddress(addr, port)
	return i
}

func (i *APM) SetDestinationService(name string, resource string) *APM {
	// i.span.Context.SetServiceTarget(apm.DestinationServiceSpanContext{Name: name, Resource: resource})
	return i
}

func InitGormAPM(ctx context.Context, dialector string, query string) *APM {
	dbName := env.EnvConfig.DatabaseName
	dbuser := env.EnvConfig.DatabaseUser
	name := truncateString(query, 24)

	spanType := fmt.Sprintf("db.%s.query", dialector)
	ins := initServiceIntrumentation(ctx, name, spanType).
		SetDatabaseName(dbName, query, dialector, dbuser).
		SetDestinationService("db", dialector).
		SetTypeAndSubType("db", dialector).
		SetAction(query)
	return ins
}

func truncateString(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	return str[:maxLength]
}
