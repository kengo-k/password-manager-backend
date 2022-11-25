package saver

import (
	"fmt"
	"os"
	"strings"

	"github.com/kengo-k/password-manager/model"
)

type FileSaver struct{}

func (s *FileSaver) Save(serializedData [][]*model.Password) {
	var sb strings.Builder
	ifNil := func(sp *string) string {
		if sp == nil {
			return ""
		} else {
			return *sp
		}
	}
	for _, passwords := range serializedData {
		head := passwords[0]
		fmt.Fprintf(&sb, "# %s: name=%s,order=%d\n", head.Category.ID, head.Category.Name, head.Category.Order)
		fmt.Fprint(&sb, "| 名称 | 説明 | ユーザ | パスワード | メール | 備考 |\n")
		fmt.Fprint(&sb, "|------|------|--------|------------|--------|------|\n")
		for _, p := range passwords {
			fmt.Fprintf(&sb, "| %s | %s | %s | %s | %s | %s |\n",
				p.Name, ifNil(p.Desc), ifNil(p.User), ifNil(p.Password), ifNil(p.Mail), ifNil(p.Note))
		}
		fmt.Fprint(&sb, "\n")
	}
	f, err := os.Create("./password.md")
	if err != nil {
		panic("failed to open file for write")
	}
	defer f.Close()
	f.Write([]byte(sb.String()))
}
