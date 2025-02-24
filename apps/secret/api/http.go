package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/mcube/app"
)

var (
	h = &handler{}
)

type handler struct {
	service secret.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(secret.AppName)
	h.service = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return secret.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.CreateSecret).
		Doc("create a secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(secret.CreateSecretRequest{}).
		Writes(secret.Secret{}))

	ws.Route(ws.GET("/").To(h.QuerySecret).
		Doc("get all secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(secret.QuerySecretRequest{}).
		Writes(secret.SecretSet{}).
		Returns(200, "OK", secret.SecretSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeSecret).
		Doc("describe an secret").
		Param(ws.PathParameter("id", "identifier of the secret").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(secret.Secret{}).
		Returns(200, "OK", secret.Secret{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteSecret).
		Doc("delete a secret").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Param(ws.PathParameter("id", "identifier of the secret").DataType("string")))
}

func init() {
	app.RegistryRESTfulApp(h)
}
