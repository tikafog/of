package dbc

import (
	"context"

	"github.com/tikafog/of"
	"github.com/tikafog/of/utils"

	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/bootnode/migrate"
	"github.com/tikafog/of/dbc/bootnode/schema"
)

func openBootNode(name of.Name, path string, o *Option) (*bootnode.Client, error) {
	dbPath, exist, err := utils.OpenDSN(utils.DSNTypeSqlite3, path, name.String(), o.debug)
	if err != nil {
		return nil, err
	}

	log.Debug("open bootnode database", "path", dbPath, "exist", exist)

	cli, err := openBootNodeDatabase(dbPath, o.Debug())
	if err != nil {
		return nil, err
	}
	ctx, ccf := context.WithTimeout(context.TODO(), o.Timeout())
	defer ccf()
	if err := createOrInitBootNode(ctx, cli, exist); err != nil {
		return nil, err
	}
	cli.Use(MutatorFunc)
	return cli, nil
}

func createOrInitBootNode(ctx context.Context, cli *bootnode.Client, exist bool) error {
	log.Warn("bootnode init", "file_exist", exist)
	if !exist {
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
		if bootnode.IsNotFound(err) {
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

func openBootNodeDatabase(path string, debug bool) (*bootnode.Client, error) {
	var options []bootnode.Option

	if debug {
		options = append(options, bootnode.Debug(), bootnode.Log(bootnodeDebug))
	}

	client, err := bootnode.Open("sqlite3", path, options...)
	if err != nil {
		return nil, Errorf("failed opening connection to database: %v", err)
	}

	return client, nil
}
