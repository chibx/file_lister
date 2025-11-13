package structs

type Config struct {
	Sleep      int    // Time in Milliseconds
	Output     string // Without extension
	EntryPoint string
	// If files should be included, -1 = all levels, 0 = none, 1 = depth one, and so on....
	//
	// This depends on the depth level of course
	IncludeFiles []int
	// How many folders deep
	//
	// Defaults to -1 for all. 0 returns nothing or just files if the IncludeFiles option is true
	MaxDepth int
	DumpAs   string
	Cwd      string
	// MaxConcurrency int16 // Unimplemented
}

type FileTree struct {
	Name    string     `json:"name"` // Folder name
	Folders []FileTree `json:"dirs,omitempty"`
	Files   []string   `json:"files,omitempty"`
}
