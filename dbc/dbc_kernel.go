package dbc

import (
	"context"

	"github.com/tikafog/of"
	"github.com/tikafog/of/utils"

	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/kernel/migrate"
	"github.com/tikafog/of/dbc/kernel/schema"
)

func openKernel(name of.Name, path string, o *Option) (*kernel.Client, error) {
	dbPath, exist, err := utils.OpenDSN(utils.DSNTypeSqlite3, path, name.String(), o.debug)
	if err != nil {
		return nil, err
	}

	log.Debug("open kernel database", "path", dbPath, "exist", exist)

	cli, err := openKernelDatabase(dbPath, o.Debug())
	if err != nil {
		return nil, err
	}
	ctx, ccf := context.WithTimeout(context.TODO(), o.Timeout())
	defer ccf()
	if err := createOrInitKernel(ctx, cli, exist); err != nil {
		return nil, err
	}
	cli.Use(MutatorFunc)
	return cli, nil
}

func createOrInitKernel(ctx context.Context, cli *kernel.Client, exist bool) error {
	if !exist {
		log.Warn(" kernel not exist, run migrating")
		err := cli.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithForeignKeys(false),
		)
		if err != nil {
			return Errorf("failed creating schema resources: %v", err)
		}
		_, err = cli.Version.Create().Save(ctx)
		if err != nil {
			return Errorf("create version failed:%v", err)
		}
		return nil
	}
	log.Debug("kernel exist")
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
			return Errorf("failed creating schema resources: %v", serr)
		}
		if kernel.IsNotFound(err) {
			_, err = cli.Version.Create().Save(ctx)
			if err != nil {
				return Errorf("create version failed:%v", err)
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
			return Errorf("failed creating schema resources: %v", err)
		}
		_, err = cli.Version.Update().SetCurrent(schema.CurrentVersion).Save(ctx)
		if err != nil {
			return Errorf("update version failed:%v", err)
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
		return nil, Errorf("failed opening connection to database: %v", err)
	}

	return client, nil
}
