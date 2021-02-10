package red_envelope

import (
	_ "github.com/ABottomCoder/account/apis/web"
	_ "github.com/ABottomCoder/account/core/accounts"
	"github.com/ABottomCoder/infra"
	"github.com/ABottomCoder/infra/base"
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
