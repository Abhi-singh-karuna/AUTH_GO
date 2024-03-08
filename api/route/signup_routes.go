package route

import (
	"time"

	"github.com/Abhi-singh-karuna/AUTH_GO/bootstrap"
	"github.com/Abhi-singh-karuna/AUTH_GO/api/controller"
	"github.com/Abhi-singh-karuna/AUTH_GO/domain"
	"github.com/Abhi-singh-karuna/AUTH_GO/mongo"
	"github.com/Abhi-singh-karuna/AUTH_GO/repository"
	"github.com/Abhi-singh-karuna/AUTH_GO/usecase"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
