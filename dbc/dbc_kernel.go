package dbc

import (
	"context"
	"fmt"
	"time"

	"github.com/tikafog/of"
	"github.com/tikafog/of/utils"

	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/kernel/migrate"
	"github.com/tikafog/of/dbc/kernel/schema"
)

func openKernel[T *kernel.Client](name of.Name, path string, o *Option) (T, error) {
	dbPath, exist, err := utils.OpenDSN(utils.DSNTypeSqlite3, path, name.String())
	if err != nil {
		return nil, err
	}
	fmt.Println("openBootNode database", "path", dbPath, "exist", exist)
	cli, err := openKernelDatabase(dbPath, o.debug)
	if err != nil {
		return nil, err
	}
	if o.timeout < MinTimeoutSec {
		o.timeout = MinTimeoutSec
	}
	ctx, ccf := context.WithTimeout(context.TODO(), time.Duration(o.timeout)*time.Second)
	defer ccf()
	if err := createOrInitKernel(ctx, cli, exist); err != nil {
		return nil, err
	}
	return cli, nil
}

func createOrInitKernel(ctx context.Context, cli *kernel.Client, exist bool) error {
	if !exist {
		fmt.Println("create not exist")
		err := cli.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithForeignKeys(false),
		)
		if err != nil {
			return fmt.Errorf("failed creating schema resources: %v", err)
		}
		_, err = cli.Version.Create().Save(ctx)
		if err != nil {
			return fmt.Errorf("create version failed:%v", err)
		}
		return nil
	}
	fmt.Println("create exist")
	boot, err := cli.Version.Query().First(ctx)
	if err != nil {
		//if db.IsNotFound(err) {
		serr := cli.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithForeignKeys(false),
		)
		if serr != nil {
			return fmt.Errorf("failed creating schema resources: %v", serr)
		}
		if kernel.IsNotFound(err) {
			_, err = cli.Version.Create().Save(ctx)
			if err != nil {
				return fmt.Errorf("create version failed:%v", err)
			}
			return nil
		}
		return nil
	}
	if boot.Current < schema.CurrentVersion {
		err := cli.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithForeignKeys(false),
		)
		if err != nil {
			return fmt.Errorf("failed creating schema resources: %v", err)
		}
		_, err = cli.Version.Update().SetCurrent(schema.CurrentVersion).Save(ctx)
		if err != nil {
			return fmt.Errorf("update version failed:%v", err)
		}
		return nil
	}
	return nil
}

func openKernelDatabase(path string, debug bool) (*kernel.Client, error) {
	var options []kernel.Option

	if debug {
		options = append(options, kernel.Debug())
	}

	client, err := kernel.Open("sqlite3", path, options...)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to database: %v", err)
	}

	return client, nil
}
