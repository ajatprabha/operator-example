package endpoints

import (
	"fmt"
	"github.com/ajatprabha/operator-example/api/v1alpha1"
	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kubelog "sigs.k8s.io/controller-runtime/pkg/log"
)

type DarkroomEndpoint struct {
	client client.Client
}

func NewDarkroomEndpoint(client client.Client) *DarkroomEndpoint {
	return &DarkroomEndpoint{client: client}
}

func (de *DarkroomEndpoint) SetupWithWS(ws *restful.WebService) {
	ws.Route(ws.GET("darkrooms").To(de.list).
		Doc("List of Darkrooms").
		Returns(200, "OK", &List{}))

	ws.Route(ws.POST("darkrooms").To(de.create).
		Doc("Create a new darkroom").
		Reads(&Darkroom{}).
		Returns(200, "OK", &Darkroom{}).
		Returns(400, "Bad Request", nil))
}

func (de *DarkroomEndpoint) list(request *restful.Request, response *restful.Response) {
	dl := new(v1alpha1.DarkroomList)
	err := de.client.List(request.Request.Context(), dl, &client.ListOptions{})
	if err != nil {
		writeError(response, 404, Error{
			Title:   "Error",
			Details: fmt.Sprintf("Could not retrieve list: %s", err),
		})
	} else {
		l := From.List(dl)
		if err := response.WriteAsJson(l); err != nil {
			writeError(response, 404, Error{
				Title:   "Error",
				Details: "Could not list resources",
			})
		}
	}
}

func (de *DarkroomEndpoint) create(request *restful.Request, response *restful.Response) {
	d := new(Darkroom)
	err := request.ReadEntity(d)
	if err != nil {
		writeError(response, 400, Error{
			Title:   "Bad Request",
			Details: "Could not read entity",
		})
		return
	}

	if err := d.Validate(); err != nil {
		writeError(response, 400, Error{
			Title:   "Validation error",
			Details: err.Error(),
		})
		return
	}

	obj := &v1alpha1.Darkroom{
		ObjectMeta: v1.ObjectMeta{Name: d.Name, Namespace: "default"},
		Spec: v1alpha1.DarkroomSpec{Version: d.Version, Source: v1alpha1.Source{
			Type: v1alpha1.Type(d.Source.Type),
		}, SubDomains: d.SubDomains},
	}

	switch obj.Spec.Source.Type {
	case v1alpha1.WebFolder:
		obj.Spec.Source.BaseURL = d.Source.BaseURL
	case v1alpha1.S3:
		obj.Spec.Source.AccessKey = d.Source.AccessKey
		obj.Spec.Source.SecretKey = d.Source.SecretKey
		obj.Spec.Source.Region = d.Source.Region
		obj.Spec.Source.PathPrefix = d.Source.PathPrefix
	}

	err = de.client.Create(request.Request.Context(), obj, &client.CreateOptions{})
	if err != nil {
		writeError(response, 400, Error{
			Title:   "Error",
			Details: fmt.Sprintf("Could not create object: %s", err),
		})
	} else {
		d := From.Object(obj)
		if err := response.WriteAsJson(d); err != nil {
			writeError(response, 422, Error{
				Title:   "Error",
				Details: "Could not write response",
			})
		}
	}
}

func writeError(response *restful.Response, httpStatus int, err Error) {
	if err := response.WriteHeaderAndJson(httpStatus, err, "application/json"); err != nil {
		kubelog.Log.Error(err, "Could not write the error response")
	}
}
