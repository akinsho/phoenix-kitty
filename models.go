package main

type ProgramArgs struct {
	Vim bool
	Source string
	Filename string
}

type OSWindow struct {
	Id               int    `json:"id"`
	IsFocused        bool   `json:"is_focused"`
	PlatformWindowID int    `json:"platform_window_id"`
	Tabs             []Tab  `json:"tabs"`
	WMClass          string `json:"wm_class"`
	WMName           string `json:"wm_name"`
}

type Window struct {
	Cmdline             []string          `json:"cmdline"`
	Columns             int               `json:"columns"`
	Cwd                 string            `json:"cwd"`
	Env                 map[string]string `json:"env"`
	ID                  int               `json:"id"`
	IsFocused           bool              `json:"is_focused"`
	IsSelf              bool              `json:"is_self"`
	Lines               int               `json:"lines"`
	Pid                 int               `json:"pid"`
	Title               string            `json:"title"`
	ForegroundProcesses []struct {
		Cmdline []string `json:"cmdline"`
		Cwd     string   `json:"cwd"`
		Pid     int      `json:"pid"`
	} `json:"foreground_processes"`
}

type Tab struct {
	ActiveWindowHistory []int    `json:"active_window_history"`
	EnabledLayouts      []string `json:"enabled_layouts"`
	ID                  int      `json:"id"`
	IsFocused           bool     `json:"is_focused"`
	Layout              string   `json:"layout"`
	LayoutOpts          struct{} `json:"layout_opts"`
	LayoutState         struct{} `json:"layout_state"`
	Title               string   `json:"title"`
	Windows             []Window `json:"windows"`
}
