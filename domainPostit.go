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

type NoteRepository interface { //role interface: https://www.martinfowler.com/bliki/RoleInterface.html

	//contextus jelzi h mi a helyzet ezzel a keressel, es a *Note mondja meg a funkcionak h mit hozzon letre
	Create(context.Context, *Note) error

	//FindByID will found you a note by a Note ID.
	//if the entity is not found, the second return value will be false.
	//so the compiler checks wether you used the found var and fails if you let it unchecked.
	//when you has a ptr of a Note here this would not be happen, and you can maybe get a nil ptr
	FindByID(context.Context, NoteID) (Note, bool, error)
	DeleteByID(context.Context, NoteID) error
}
