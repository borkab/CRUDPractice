package postit

import "context"

// Note is the domain entity
// belso kor
type Note struct {
	Title string
	Body  string
	ID    NoteID
}

type NoteID string

// a NoteRepository egy Note tarolo
// https://github.com/adamluzsi/frameless/blob/main/ports/crud/crud.go

type NoteRepository interface {
	Create(context.Context, *Note) error //contextus jelzi h mi a helyzet ezzel a keressel, es a *Note mondja meg a funkcionak h mit hozzon letre
	FindByID(context.Context, NoteID) (*Note, error)
	DeleteByID(context.Context, NoteID) error
}
