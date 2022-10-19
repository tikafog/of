package dbc

import (
	"context"
	"io"
	"log"

	"github.com/tikafog/of"
	"github.com/tikafog/of/utils"

	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/media/migrate"
	"github.com/tikafog/of/dbc/media/schema"
)

func openMedia(name of.Name, path string, o *Option) (*media.Client, io.Closer, error) {
	dbPath, exist, err := utils.OpenDSN(utils.DSNTypeSqlite3, path, name.String(), o.debug)
	if err != nil {
		return nil, nil, err
	}
	if debug {
		log.Println("[DBC] open database", "path", dbPath, "exist", exist)
	}
	cli, err := openMediaDatabase(dbPath)
	if err != nil {
		return nil, nil, err
	}
	ctx, ccf := context.WithTimeout(context.TODO(), o.Timeout())
	defer ccf()
	if err := createOrInitMedia(ctx, cli, exist); err != nil {
		return nil, nil, err
	}
	cli.Use(MutatorFunc)
	return cli, cli, nil
}

func createOrInitMedia(ctx context.Context, cli *media.Client, exist bool) error {
	if !exist {
		if debug {
			log.Println("[DBC] media not exist")
		}
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
	if debug {
		log.Println("[DBC] media exist")
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
		if media.IsNotFound(err) {
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

func openMediaDatabase(path string) (*media.Client, error) {
	var options []media.Option

	if debug {
		options = append(options, media.Debug())
	}

	client, err := media.Open("sqlite3", path, options...)
	if err != nil {
		return nil, Errorf("failed opening connection to database: %v", err)
	}

	return client, nil
}
