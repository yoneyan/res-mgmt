package v0

import (
	"fmt"
	"github.com/yoneyan/res-mgmt/pkg/api/core/item"
)

func check(input item.Item) error {
	// check
	if input.Name == "" {
		return fmt.Errorf("no data: name")
	}
	if input.OwnerID == 0 {
		return fmt.Errorf("no data: ownerID")
	}
	if input.TypeID == 0 {
		return fmt.Errorf("no data: typeID")
	}
	if input.NOC == "" {
		return fmt.Errorf("no data: noc")
	}
	return nil
}
