package driver

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetDriverByID gets driver by id
func (dr *Repo) GetDriverByID(ID string) (*Driver, error) {
	d := new(Driver)
	err := dr.DB.Model(d).Where("id = ?", ID).Select()
	if err != nil {
		return nil, gqlerror.Errorf("cab is not available")
	}

	return d, nil
}
