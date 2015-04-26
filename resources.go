package main
import "time"
const (
	Local = iota
	Remote = iota
)
const (
	Asset = iota
	Page = iota
)


type SiteResource struct {
	ResourceType uint
	ResourceLocation uint
	Ref string
	Delta time.Duration 
}
