package static

import (
	"distivity/config"
	"distivity/types"
)

func GetConfig() types.Config {
	configVariables := config.GetVariables()
	return configVariables
}
