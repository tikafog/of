package dbc

import (
	"context"
	"fmt"
	"time"

	"github.com/tikafog/of/utils"

	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/media/migrate"
	"github.com/tikafog/of/dbc/media/schema"
)

func openMedia[T *media.Client](name, path string, o *Option) (T, error) {
	dbPath, exist, err := utils.OpenDSN(utils.DSNTypeSqlite3, path, name)
	if err != nil {
		return nil, err
	}
	fmt.Println("open database", "path", dbPath, "exist", exist)
	cli, err := openMediaDatabase(dbPath, true)
	if err != nil {
		return nil, err
	}
	if o.timeout < MinTimeoutSec {
		o.timeout = MinTimeoutSec
	}
	ctx, ccf := context.WithTimeout(context.TODO(), time.Duration(o.timeout)*time.Second)
	defer ccf()
	if err := createOrInitMedia(ctx, cli, exist); err != nil {
		return nil, err
	}
	return cli, nil
}

func createOrInitMedia(ctx context.Context, cli *media.Client, exist bool) error {
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
		if media.IsNotFound(err) {
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

func openMediaDatabase(path string, debug bool) (*media.Client, error) {
	var options []media.Option

	if debug {
		options = append(options, media.Debug())
	}

	client, err := media.Open("sqlite3", path, options...)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to database: %v", err)
	}

	return client, nil
}
