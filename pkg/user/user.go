package user

import "github.com/lib/pq"

type User struct {
	VkID   int32
	Groups pq.Int64Array `gorm:"type:integer[]"`
}
