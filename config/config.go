package config

type Config interface {
	Parse(any) error
}
