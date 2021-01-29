package demo

import (
	"github.com/bitwormhole/go-wormhole-core/application"
)

// Foo the client
type Foo struct {
	// context int  `id:"bill" class:"man"`
	inst *Bar `id:"1" class:"x y z"  initMethod:"start" destroyMethod:"stop" `
}

func (inst *Foo) Inject(ctx application.Context) error {

	// inst.target.
	return nil
}
