package commands

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/jmek087/go-workwx-client/v2"
)

func cmdDeptCreate(c *cli.Context) error {
	cfg := mustGetConfig(c)

	name := c.String(flagName)
	if name == "" {
		return errors.New("dept name is required")
	}

	parentID := c.Int64(flagParentID)
	if parentID == 0 {
		parentID = 1
	}

	order := uint32(c.Uint(flagOrder))

	deptInfo := &workwx.DeptInfo{
		Name:     name,
		NameEn:   c.String(flagNameEn),
		ParentID: parentID,
		Order:    order,
	}

	app := cfg.MakeWorkwxApp()
	deptID, err := app.CreateDept(deptInfo)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("created dept id = %d\n", deptID)
	return nil
}
