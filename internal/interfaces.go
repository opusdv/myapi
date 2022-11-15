package internal

type Drive interface {
	CheckDrive() (bool, error)
}
