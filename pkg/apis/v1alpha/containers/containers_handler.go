/*
Copyright 2018 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package containers

import (
	"github.com/emicklei/go-restful"

	"kubesphere.io/kubesphere/pkg/constants"
	"kubesphere.io/kubesphere/pkg/filter/route"
	"kubesphere.io/kubesphere/pkg/models"
)

func Register(ws *restful.WebService) {
	ws.Route(ws.GET("/namespaces/{namespace}/pods/{podname}/containers").To(handleContainersUnderNameSpaceAndPod).Filter(route.RouteLogging)).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/nodes/{nodename}/namespaces/{namespace}/pods/{podname}/containers").To(handleContainersUnderNodeAndNameSpaceAndPod).Filter(route.RouteLogging)).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON)
}

func handleContainersUnderNameSpaceAndPod(request *restful.Request, response *restful.Response) {
	var result constants.ResultMessage
	var resultNameSpaces []models.ResultNameSpaceForContainer
	var resultNameSpace models.ResultNameSpaceForContainer

	resultNameSpace = models.FormatContainersMetrics("", request.PathParameter("namespace"), request.PathParameter("podname"))

	resultNameSpaces = append(resultNameSpaces, resultNameSpace)

	result.Data = resultNameSpaces
	response.WriteAsJson(result)
}

func handleContainersUnderNodeAndNameSpaceAndPod(request *restful.Request, response *restful.Response) {
	var result constants.ResultMessage
	var resultNameSpaces []models.ResultNameSpaceForContainer
	var resultNameSpace models.ResultNameSpaceForContainer

	resultNameSpace = models.FormatContainersMetrics(request.PathParameter("nodename"), request.PathParameter("namespace"), request.PathParameter("podname"))

	resultNameSpaces = append(resultNameSpaces, resultNameSpace)

	result.Data = resultNameSpaces
	response.WriteAsJson(result)
}