package use_case

type DataReader interface {
	Read() (record []Field, err error)
}
