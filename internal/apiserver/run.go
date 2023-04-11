package apiserver

func Run() error {
	server, err := createAPIServer()
	if err != nil {

	}
	return server.PrepareRun().Run()
}
