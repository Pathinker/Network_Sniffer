package decoder

func (d *Decoder) colorize(s string) string {
	// ANSI green text
	return "\033[32m" + s + "\033[0m"
}
