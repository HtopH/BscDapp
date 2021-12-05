package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
