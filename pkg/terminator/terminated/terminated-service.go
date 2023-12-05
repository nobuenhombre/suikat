package terminated

type IGracefulShutDownService interface {
	Run()
	GracefulShutDown()
}
