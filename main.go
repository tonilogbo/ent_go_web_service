package main

import (
	"context"
	"entrds/artists"
	"entrds/database"
	"entrds/ent"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func DateOfBirth(year int, month time.Month, day int) time.Time {
	return time.Date(year,month,day,0,0,0,0,time.UTC)
}

func CreateGraph(ctx context.Context, client *ent.Client) error {

	yvette, err := client.Artist.Create().SetName("Yvette Young").SetGender("Female").SetDateOfBirth(DateOfBirth(1991,time.June,28)).Save(ctx)
	if err != nil {
		return err
	}

	covet, err := client.Band.Create().SetName("Covet").SetYearFormed(2014).AddMembers(yvette).Save(ctx)
		if err != nil {
		return err
	}



	_, err = client.Album.Create().SetName("Technicolor").SetReleaseYear(2020).AddBand(covet).Save(ctx)
	if err != nil {
		return err
	}
	return nil

}

// func QueryUser(ctx context.Context, client *ent.Client) (*ent.Artist, error) {
// 	y, err := client.Debug().Artist.Query().Where(artist.Name("Yvette Young")).First(ctx)
// 	if err != nil {
//         return nil, fmt.Errorf("failed querying user: %w", err)
//     }
//     log.Println("user returned: ", y)
//     return y, nil

// }
const basePath = "/api"

func main() {
	database.SetupDatabase()
	defer database.Client.Close()
	// ctx := context.Background()
	// if err := client.Schema.Create(ctx); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	// if err = CreateGraph(ctx,client); err != nil {
	// 	log.Fatalf("Error in graph: %v", err)
	// }
	// if _, err := QueryUser(ctx, database.Client); err != nil {
	// 	log.Fatalf("L taken: %v", err)
	// }
	artists.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}