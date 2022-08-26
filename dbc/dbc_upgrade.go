package dbc

import (
	"context"
	"fmt"
	"time"

	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc/upgrade"
	"github.com/tikafog/of/dbc/upgrade/migrate"
	"github.com/tikafog/of/dbc/upgrade/schema"
	"github.com/tikafog/of/utils"
)

func openUpgrade[T *upgrade.Client](name of.Name, path string, o *Option) (T, error) {
	dbPath, exist, err := utils.OpenDSN(utils.DSNTypeSqlite3, path, name.String())
	if err != nil {
		return nil, err
	}
	fmt.Println("openBootNode database", "path", dbPath, "exist", exist)
	cli, err := openUpgradeDatabase(dbPath, o.debug)
	if err != nil {
		return nil, err
	}
	if o.timeout < MinTimeoutSec {
		o.timeout = MinTimeoutSec
	}
	ctx, ccf := context.WithTimeout(context.TODO(), time.Duration(o.timeout)*time.Second)
	defer ccf()
	if err := createOrInitUpgrade(ctx, cli, exist); err != nil {
		return nil, err
	}
	return cli, nil
}

func createOrInitUpgrade(ctx context.Context, cli *upgrade.Client, exist bool) error {
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
		if upgrade.IsNotFound(err) {
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

func openUpgradeDatabase(path string, debug bool) (*upgrade.Client, error) {
	var options []upgrade.Option

	if debug {
		options = append(options, upgrade.Debug())
	}

	client, err := upgrade.Open("sqlite3", path, options...)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to database: %v", err)
	}

	return client, nil
}
