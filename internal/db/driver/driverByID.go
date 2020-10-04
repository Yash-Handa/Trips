package driver

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const picPath = "/raw/assets/drivers/"
const picExt = "jpg"

// GetDriverByID gets driver by id
func (dr *Repo) GetDriverByID(ID string) (*Driver, error) {
	d := new(Driver)
	err := dr.DB.Model(d).Where("id = ?", ID).Select()
	if err != nil {
		return nil, gqlerror.Errorf("cab is not available")
	}

	d.Pic = picPath + d.Pic + "." + picExt

	return d, nil
}
