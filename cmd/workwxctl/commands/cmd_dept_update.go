package commands

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/jmek087/go-workwx-client/v2"
)

func cmdDeptUpdate(c *cli.Context) error {
	cfg := mustGetConfig(c)

	deptID := c.Int64(flagDeptID)
	if deptID == 0 {
		return errors.New("dept-id is required")
	}

	deptInfo := &workwx.DeptInfo{
		ID:       deptID,
		Name:     c.String(flagName),
		NameEn:   c.String(flagNameEn),
		ParentID: c.Int64(flagParentID),
		Order:    uint32(c.Uint(flagOrder)),
	}

	app := cfg.MakeWorkwxApp()
	err := app.UpdateDept(deptInfo)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("updated dept id = %d\n", deptID)
	return nil
}
