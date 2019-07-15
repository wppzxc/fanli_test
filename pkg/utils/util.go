package utils

import (
	"github.com/spf13/pflag"
	"github.com/wpp/fanli_test/pkg/types"
	"strings"
)

func GetMsg(result types.ItemResult) string {
	str := ""
	for _, i := range result.Data {
		str = str + i.ExtendDocument + "\n"
	}
	return str
}

// WordSepNormalizeFunc changes all flags that contain "_" separators
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}
