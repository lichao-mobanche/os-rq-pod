package routers

import (
	"github.com/cfhamlet/os-rq-pod/app/controllers"
	"github.com/cfhamlet/os-rq-pod/pkg/ginserv"
	"github.com/cfhamlet/os-rq-pod/pkg/ginserv/route"
)

// RouteQueuesCtrl TODO
func RouteQueuesCtrl(root ginserv.RouterGroup, ctrl *controllers.QueuesController) {
	g := root.Group("/queues/")
	routes := []*route.Route{
		route.New(g.POST, "/", ctrl.Queues),
		route.New(g.GET, "/info/", ctrl.QueuesInfo),
		route.New(g.GET, "/view/", ctrl.ViewQueues),
	}
	route.Bind(routes, controllers.ErrorCode)
}
