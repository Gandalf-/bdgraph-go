package bdgraph

type Option struct {
	name  string
	color string
}

var OptionCleanup Option = Option{
	name: "cleanup",
}

var OptionCircular Option = Option{
	name: "circular",
}

var OptionNext Option = Option{
	name:  "color_next",
	color: "lightskyblue",
}

var OptionComplete Option = Option{
	name:  "color_complete",
	color: "springgreen",
}

var OptionUrgent Option = Option{
	name:  "color_urgent",
	color: "crimson",
}
