// Model objects for appkit.
// Some of the model code for appkit is generated and can be found at
// github.com/Nonzero-Sum-Solutions/appkit/model/gen/appkit.
package model

import gmodel "github.com/Nonzero-Sum-Solutions/appkit/model/gen/appkit"

type appImpl struct {
	id         string
	version    string
	configName string
}

func (a *appImpl) ID() string {
	return a.id
}

func (a *appImpl) Version() string {
	return a.version
}

func (a *appImpl) ConfigName() string {
	return a.configName
}

func NewApp(id, version, configName string) gmodel.App {
	return &appImpl{
		id:         id,
		version:    version,
		configName: configName,
	}
}
