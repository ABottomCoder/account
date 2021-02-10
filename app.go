package red_envelope

import (
	"github.com/ABottomCoder/infra"
	"github.com/ABottomCoder/infra/base"
	_ "github.com/ABottomCoder/account/apis/web"
	_ "github.com/ABottomCoder/account/core/accounts"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
