package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"

	ElectricalRank = "elecRank"

	AccessoryRank = "acceRank"
)

// 视频点击次数
func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
