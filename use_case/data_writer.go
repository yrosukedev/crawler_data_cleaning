package use_case

type DataWriter interface {
	Write(record []Field) error
}
