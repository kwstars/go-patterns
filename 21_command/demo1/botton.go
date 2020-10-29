package main

type Botton struct {
	command Command
}

func (b *Botton) press() {
	b.command.Execute()
}

type OnCommand struct {
	device Device
}

func (o *OnCommand) Execute() {
	o.device.On()
}

type OffCommand struct {
	device Device
}

func (o *OffCommand) Execute() {
	o.device.Off()
}
