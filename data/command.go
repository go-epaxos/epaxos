package data

import (
	"bytes"
)

type Command []byte

type Commands []Command

func (self Command) Compare(other Command) int {
	return bytes.Compare(self, other)
}

func (c Commands) GetCopy() Commands {
	cmds := make(Commands, len(c))
	for i := range cmds {
		cmds[i] = make(Command, len(c[i]))
		copy(cmds[i], c[i])
	}
	return cmds
}
