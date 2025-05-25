package out

type FileApi interface {
	GetFile(id string) (string, error)
}
