//+build wireinject

package definition

import (
	"github.com/Liar233/empty-app/vendor/github.com/google/wire"
	"github.com/Liar233/empty-app/internal/app"
)

func InitApp() app.Application {
	wire.Build(app.NewApplication)

	return app.Application{}
}
