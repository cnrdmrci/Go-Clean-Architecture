package interface_services

type LoggerService interface {
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
