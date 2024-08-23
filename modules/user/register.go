package user

const (
	basePath   = "v1/users"
	detailPath = basePath + "/:id"
)

func (module *UserModule) registerHandlers() {
	module.app.Post(basePath, module.controller.handleCreate)
	module.app.Get(basePath, module.controller.handleList)
	module.app.Get(detailPath, module.controller.handleDetail)
	module.app.Put(detailPath, module.controller.handleUpdate)
	module.app.Delete(detailPath, module.controller.handleDelete)
}
