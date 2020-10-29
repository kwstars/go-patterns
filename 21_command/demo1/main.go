package main

func main() {
	tv := &tv{}

	on := &OnCommand{device: tv}

	off := &OffCommand{device: tv}

	bOn := &Botton{command: on}
	bOn.press()

	bOff := &Botton{command: off}
	bOff.press()
}
