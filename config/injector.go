package config

// import (
// 	"erdemkayacomtr/controllers"
// 	"erdemkayacomtr/repositories"
// 	"erdemkayacomtr/services"

// 	"github.com/google/wire"
// )

// var db = wire.NewSet(ConnectToDB)

// var categoryRepositorySet = wire.NewSet(repositories.CategoryRepositoryInit, wire.Bind(new(repositories.CategoryRepository), new(*repositories.CategoryRepositoryImpl)))
// var categoryServiceSet = wire.NewSet(services.CategoryServiceInit, wire.Bind(new(services.CategoryService), new(*services.CategoryServiceImpl)))
// var categoryControllerSet = wire.NewSet(controllers.CategoryControllerInit, wire.Bind(new(controllers.CategoryController), new(*controllers.CategoryControllerImpl)))

// func Init() *Initialization {
// 	wire.Build(db, categoryRepositorySet, categoryServiceSet, categoryControllerSet, NewInitializer)
// 	return &Initialization{}
// }
