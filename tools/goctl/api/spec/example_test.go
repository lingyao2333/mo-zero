package spec_test

import (
	"fmt"
	"github.com/lingyao2333/mo-zero/tools/goctl/api/spec"
)

func ExampleMember_GetEnumOptions() {
	member := spec.Member{
		Tag: `json:"foo,options=foo|bar|options|123"`,
	}
	fmt.Println(member.GetEnumOptions())
	// Output:
	// [foo bar options 123]
}
