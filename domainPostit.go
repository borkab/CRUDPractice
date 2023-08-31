package postit

import (
	"context"
	"testing"

	"github.com/adamluzsi/testcase/assert"
)

// Note is the domain entity
// belso kor
type Note struct {
	Title string
	Body  string
	ID    NoteID //az Id-t a repository birtokolja
}

type NoteID string

// a NoteRepository egy Note tarolo
// https://github.com/adamluzsi/frameless/blob/main/ports/crud/crud.go

type NoteRepository interface { //role interface: https://www.martinfowler.com/bliki/RoleInterface.html

	//contextus jelzi h mi a helyzet ezzel a keressel, es a *Note mondja meg a funkcionak h mit hozzon letre
	Create(context.Context, *Note) error

	//FindByID will found you a note by a Note ID.
	//if the entity is not found, the second return value will be false.
	//so the compiler checks wether you used the found var and fails if you let it unchecked.
	//when you has a ptr of a Note here this would not be happen, and you can maybe get a nil ptr
	FindByID(context.Context, NoteID) (Note, bool, error)
	DeleteByID(context.Context, NoteID) error
	FindAll(context.Context) ([]Note, error)
}

func NoteRepositoryContract(t *testing.T, subject NoteRepository) {
	t.Run("the repo is not empty", func(t *testing.T) {
		note := Note{
			Title: "dolphins",
			Body:  "thx for the fishes",
		}

		ctx := context.Background()

		assert.NoError(t, subject.Create(ctx, &note)) //megneztuk h hiba nelkul vegig futott-e a Create func
		assert.NotEmpty(t, note.ID, "we expect that a successful create call supply an ID in the note value")

		gotNote, found, err := subject.FindByID(ctx, note.ID)
		assert.NoError(t, err) //megnezzuk h hiba nelkul vegigfutott e a find func
		assert.True(t, found, "we expect to find the note using the ID that we created with the Create func")
		assert.Equal(t, gotNote, note) //megnezzuk h a megtalalt note megegyezik-e a letrehozott note-al

		err = subject.DeleteByID(ctx, note.ID)
		assert.NoError(t, err)                       //megnezzuk h hiba nelkul vegig futott e a delete func
		_, found, _ = subject.FindByID(ctx, note.ID) //megnezzuk h a note-unk Id-jat torles utan megtalalja e
		assert.False(t, found, "we expect that found will be false after deleting our note")

	})

	t.Run("find all the existing notes", func(t *testing.T) {

		note := Note{
			Title: "dolphins",
			Body:  "thx for the fishes",
		}

		note2 := Note{
			Title: "allways",
			Body:  "take a towel with you",
		}

		ctx := context.Background()

		var myNotes []Note                     //csinalunk egy note-okbol allo listat
		myNotes = append(myNotes, note, note2) //beletesszuk ebbe a listaba a 2 note-unkat

		subject.Create(ctx, &note)  //letrahozzuk a korabban torolt jegyzetet
		subject.Create(ctx, &note2) //letrehozunk egy masik jegyzetet is

		gotNotes, err := subject.FindAll(ctx) //megkeressuk az osszes letrehozott jegyzetunket es belerakjuk egy listaba
		assert.NoError(t, err)                //megnezzuk h hiba nelkul lefutott e a kereses

		//Lists are always ordered, but we don't want to expect that from our findAll method,
		// so we have to check the length and the elements in our lists
		assert.Equal(t, len(gotNotes), len(myNotes)) //megnezzuk h a megtalalt note-okbol allo lista hossza megegyezik e a 2 note-bol allo lista hosszaval
		assert.ContainExactly(t, gotNotes, myNotes)  //megnezzuk h a myNotes lista osszes eleme megtalalhato-e a gotNotes listaban
	})

}
