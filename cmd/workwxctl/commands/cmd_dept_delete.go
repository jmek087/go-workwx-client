package commands

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

func cmdDeptDelete(c *cli.Context) error {
	cfg := mustGetConfig(c)

	deptID := c.Int64(flagDeptID)
	if deptID == 0 {
		return errors.New("dept-id is required")
	}

	app := cfg.MakeWorkwxApp()
	err := app.DeleteDept(deptID)
	if err != nil {
		fmt.Printf("error = %+v\n", err)
		return err
	}

	fmt.Printf("deleted dept id = %d\n", deptID)
	return nil
}
