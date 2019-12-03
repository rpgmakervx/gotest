package testifelse

func TestSwitch(x int) string {
	switch x {
	case 0:
		return "I got no money!"
	case 1:
		return "That's a good idea!"
	case 2:
		return "Come to see me on Tuesday"
	case 3, 4, 5:
		return "There are too many of them"
	default:
		return "I'll give left you by default"
	}
}
