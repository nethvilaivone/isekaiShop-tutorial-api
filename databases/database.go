
package databases

import "gorm.io/gorm"

type Databases interface {
	ConnectedGetting() *gorm.DB
}
