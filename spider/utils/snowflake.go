package utils

import (
	"github.com/bwmarrin/snowflake"
)

func SnowFlake() int64 {
	node, _ := snowflake.NewNode(1)
	// Generate a snowflake ID.
	return int64(node.Generate())
}
