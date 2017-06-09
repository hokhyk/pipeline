package restfulserver

import "net/http"
import "github.com/rancher/go-rancher/api"
import "github.com/rancher/go-rancher/client"
import "github.com/rancher/pipeline/jenkins"
import "strings"
import "github.com/rancher/pipeline/pipeline"

//Server rest api server
type Server struct {
	PipelineContext *pipeline.PipelineContext
}

//ListPipelines query List of pipelines
func (s *Server) ListPipelines(rw http.ResponseWriter, req *http.Request) error {
	apiContext := api.GetApiContext(req)
	headers := req.Header
	for header, content := range headers {
		println(header, ":[", strings.Join(content, ","), "]")
	}
	println(req.Host)
	println(req.URL)
	apiContext.Write(&client.GenericCollection{
		Data: toPipelineCollections(s.PipelineContext.ListPipelines()),
	})
	return nil
}

func (s *Server) ListPipeline(rw http.ResponseWriter, req *http.Request) error {
	apiContext := api.GetApiContext(req)
	apiContext.Write(nil)
	return nil
}

func (s *Server) ListActivities(rw http.ResponseWriter, req *http.Request) error {
	apiContext := api.GetApiContext(req)
	apiContext.Write(&client.GenericCollection{
		Data: []interface{}{
		//toActivityResource(apiContext),
		},
	})
	return nil
}

func (s *Server) CreatePipelineWithXML(rw http.ResponseWriter, req *http.Request) error {
	apiContext := api.GetApiContext(req)
	if err := jenkins.GetCSRF(); err != nil {
		return err
	}
	if err := jenkins.CreateJob("test1"); err != nil {
		return err
	}
	apiContext.Write("ok")
	return nil
}

func NewServer(pipelineContext *pipeline.PipelineContext) *Server {
	return &Server{
		PipelineContext: pipelineContext,
	}
}
