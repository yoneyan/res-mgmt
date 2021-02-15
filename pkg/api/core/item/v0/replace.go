package v0

import (
	"github.com/yoneyan/res-mgmt/pkg/api/core/item"
)

func replace(input, replace item.Item) error {
	// Name
	if input.Name != "" {
		replace.Name = input.Name
	}
	// NOC
	if input.NOC != "" {
		replace.NOC = input.NOC
	}
	// Comment
	if input.Comment != "" {
		replace.Comment = input.Comment
	}

	// uint boolean
	// TypeID
	if input.TypeID != replace.TypeID {
		replace.TypeID = input.TypeID
	}
	// OwnerID
	if input.OwnerID != replace.OwnerID {
		replace.OwnerID = input.OwnerID
	}
	return nil
}
