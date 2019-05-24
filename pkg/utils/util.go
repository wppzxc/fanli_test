package utils

import "github.com/wpp/fanli_test/pkg/types"

func GetMsg(result types.ItemResult) string {
	str := ""
	for _, i := range result.Data{
		str = str + i.ExtendDocument + "\n"
	}
	return str
}
