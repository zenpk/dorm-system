package util

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

func GenSnowflakeString() (string, error) {
	node, err := snowflake.NewNode(viper.GetInt64("snowflake.node"))
	if err != nil {
		return "", err
	}
	snowflakeId := node.Generate()
	return snowflakeId.Base64(), nil
}
