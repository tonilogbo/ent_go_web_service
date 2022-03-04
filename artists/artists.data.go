package artists

import (
	"context"
	"entrds/database"
	"entrds/ent"
	"log"
	"time"
)

const ArtistNotFound = "ent: artist not found"

func getArtistList() ([]*ent.Artist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	artists, err := database.Client.Artist.Query().All(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return artists, nil
}

func insertArtist(a ent.Artist) (*ent.Artist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	artist, err := database.Client.Artist.Create().SetName(a.Name).SetGender(a.Gender).SetDateOfBirth(a.DateOfBirth).Save(ctx)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func getArtist(id int) (*ent.Artist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	a, err := database.Client.Artist.Get(ctx, id)
	if err != nil && err.Error() == ArtistNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return a, nil
}

func updateArtist(id int, a ent.Artist) (*ent.Artist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	artist, err := database.Client.Artist.UpdateOneID(id).SetName(a.Name).SetDateOfBirth(a.DateOfBirth).SetGender(a.Gender).Save(ctx)
	if err != nil {
		if err.Error() == ArtistNotFound {
			return nil, nil
		}
		return nil, err
	}
	return artist, nil
}

func deleteArtist(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := database.Client.Artist.DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil
}