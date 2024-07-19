package persistence

import (
	"context"

	"github.com/fatih/structs"
	"github.com/navidrome/navidrome/db"
	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/model/request"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("ArtistRepository", func() {
	var repo model.ArtistRepository

	BeforeEach(func() {
		ctx := log.NewContext(context.TODO())
		ctx = request.WithUser(ctx, model.User{ID: "userid"})
		repo = NewArtistRepository(ctx, NewDBXBuilder(db.Db()))
	})

	Describe("Count", func() {
		It("returns the number of artists in the DB", func() {
			Expect(repo.CountAll()).To(Equal(int64(2)))
		})
	})

	Describe("Exists", func() {
		It("returns true for an artist that is in the DB", func() {
			Expect(repo.Exists("3")).To(BeTrue())
		})
		It("returns false for an artist that is in the DB", func() {
			Expect(repo.Exists("666")).To(BeFalse())
		})
	})

	Describe("Get", func() {
		It("saves and retrieves data", func() {
			Expect(repo.Get("2")).To(Equal(&artistKraftwerk))
		})
	})

	Describe("GetIndex", func() {
		It("returns the index", func() {
			idx, err := repo.GetIndex()
			Expect(err).To(BeNil())
			Expect(idx).To(Equal(model.ArtistIndexes{
				{
					ID: "B",
					Artists: model.Artists{
						artistBeatles,
					},
				},
				{
					ID: "K",
					Artists: model.Artists{
						artistKraftwerk,
					},
				},
			}))
		})
	})

	Describe("dbArtist mapping", func() {
		var a *model.Artist
		BeforeEach(func() {
			a = &model.Artist{ID: "1", Name: "Van Halen", SimilarArtists: []model.Artist{
				{ID: "2", Name: "AC/DC"}, {ID: "-1", Name: "Test;With:Sep,Chars"},
			}}
		})
		It("maps fields", func() {
			dba := &dbArtist{Artist: a}
			m := structs.Map(dba)
			Expect(dba.PostMapArgs(m)).To(Succeed())
			Expect(m).To(HaveKeyWithValue("similar_artists", "2:AC%2FDC;-1:Test%3BWith%3ASep%2CChars"))

			other := dbArtist{SimilarArtists: m["similar_artists"].(string), Artist: &model.Artist{
				ID: "1", Name: "Van Halen",
			}}
			Expect(other.PostScan()).To(Succeed())

			actual := other.Artist
			Expect(*actual).To(MatchFields(IgnoreExtras, Fields{
				"ID":   Equal(a.ID),
				"Name": Equal(a.Name),
			}))
			Expect(actual.SimilarArtists).To(HaveLen(2))
			Expect(actual.SimilarArtists[0].ID).To(Equal("2"))
			Expect(actual.SimilarArtists[0].Name).To(Equal("AC/DC"))
			Expect(actual.SimilarArtists[1].ID).To(Equal("-1"))
			Expect(actual.SimilarArtists[1].Name).To(Equal("Test;With:Sep,Chars"))
		})
	})
})
