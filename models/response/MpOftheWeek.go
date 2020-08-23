package response

import "github.com/MbungeApp/mbunge-core/models/db"

type MpOftheWeek struct {
	Mp      db.MP
	Details db.MpLive
}
