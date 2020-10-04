package cab

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const picPath = "/raw/assets/cabs/"
const picExt = "jpg"

// GetCabByID gets cab by id
func (cr *Repo) GetCabByID(ID string) (*Cab, error) {
	c := new(Cab)
	err := cr.DB.Model(c).Where("id = ?", ID).Select()
	if err != nil {
		return nil, gqlerror.Errorf("cab is not available")
	}

	c.Pic = picPath + c.Pic + "." + picExt

	return c, nil
}
