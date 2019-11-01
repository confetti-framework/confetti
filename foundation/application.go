package foundation

import (
	appInterface "laravelgo/interface/application"
)

type ApplicationStruct struct {
	Container           ContainerStruct
	HasBeenBootstrapped bool
}

func Application() ApplicationStruct {
	application := ApplicationStruct{Container: Container()}

	application.Container.Instance((*appInterface.Container)(nil), application)
	return application
}
