package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ogataka50/ent-test/ent"
	"github.com/ogataka50/ent-test/ent/item"
	"github.com/ogataka50/ent-test/ent/itemdescription"
	"github.com/ogataka50/ent-test/ent/migrate"
	"github.com/pioz/faker"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

const (
	dataSource = "postgres://enttest:enttest@localhost:5432/enttest?sslmode=disable"
)

func main() {
	client := connectDB()
	defer client.Close()

	ctx := context.Background()

	//////////////////
	// create item
	log.Println("============= create item =============")
	i1, err := client.Item.Create().
		SetName(faker.StringWithSize(10)).
		Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating item: %v", err)
	}
	i2, err := client.Item.Create().
		SetName(faker.StringWithSize(10)).
		Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating item: %v", err)
	}
	log.Println("============= item was created: ", i1, i2)

	// create item description
	log.Println("============= create item description =============")
	d, err := client.ItemDescription.Create().
		SetDescription(faker.StringWithSize(30)).
		SetOwner(i1).
		Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating item description: %v", err)
	}
	log.Println("============= item description was created: ", d)

	// create item variation
	log.Println("============= create item variation =============")
	v1, err := client.ItemVariation.Create().
		SetVariantName("variation : " + faker.StringWithSize(5)).
		SetOriginItemID(i1.ID).
		Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating item variation1: %v", err)
	}
	log.Println("============= item description1 was created: ", v1)

	v2, err := client.ItemVariation.Create().
		SetVariantName("variation : " + faker.StringWithSize(5)).
		SetOriginItemID(i1.ID).
		Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating item variation2: %v", err)
	}

	log.Println("============= item variation2 was created: ", v2)

	// create item group
	log.Println("============= create item group =============")
	g, err := client.ItemGroup.Create().
		SetName("group: "+faker.StringWithSize(5)).
		AddGroupItem(i1, i2).
		Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating item group: %v", err)
	}
	log.Println("============= item group was created: ", g)

	///////////////////
	// select
	log.Println("============= find items =============")
	items, err := client.Item.Query().
		Where(
			item.Name(i1.Name),
		).
		All(ctx)
	if err != nil {
		fmt.Errorf("failed find items: %v", err)
	}

	log.Println("============= find items: ", items)

	// select item description
	log.Println("============= find item description =============")
	itemDescriptions, err := client.Item.Query().
		Where(
			item.Name(i1.Name),
		).
		QueryItemDescription().
		All(ctx)
	if err != nil {
		fmt.Errorf("failed find item description: %v", err)
	}

	log.Println("============= find item description: ", itemDescriptions)

	// select item variation
	log.Println("============= find item variation =============")
	itemVariation, err := client.Item.Query().
		Where(
			item.Name(i1.Name),
		).
		QueryItemVariation().
		All(ctx)
	if err != nil {
		fmt.Errorf("failed find item variation: %v", err)
	}

	log.Println("============= find item variation: ", itemVariation)

	// select item variation
	log.Println("============= find owner item by description =============")
	itemByOwner, err := client.ItemDescription.Query().
		Where(
			itemdescription.Description(d.Description),
		).
		QueryOwner().
		All(ctx)
	if err != nil {
		fmt.Errorf("failed find owner item: %v", err)
	}

	log.Println("============= find owner item  by description: ", itemByOwner)

	// select item variation
	log.Println("============= find item group by item =============")
	GroupByItem, err := client.Item.Query().
		Where(
			item.Name(i1.Name),
		).
		QueryItemGroup().
		All(ctx)
	if err != nil {
		fmt.Errorf("failed find group by item: %v", err)
	}

	log.Println("============= find group by item: ", GroupByItem)
}

func connectDB() *ent.Client {
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	loggerAdapter := zerologadapter.New(zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}))
	db = sqldblogger.OpenDriver(dataSource, db.Driver(), loggerAdapter)

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
