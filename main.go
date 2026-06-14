package main

import "fmt"

// File represents a file in the repository
type File struct {
	Name             string
	HasStagedChanges bool
}

// FilesPanel represents the UI panel for files
type FilesPanel struct {
	SelectedIdx int
}

// Model represents the GUI model state
type Model struct {
	Files []*File
}

// Panels represents the GUI panels state
type Panels struct {
	Files FilesPanel
}

// State represents the GUI state
type State struct {
	Model  Model
	Panels Panels
}

// Gui represents the main GUI controller
type Gui struct {
	State        State
	StagingCache map[string]string // Tracks staged hunks/status by file name
}

// NewGui creates a new Gui instance
func NewGui() *Gui {
	return &Gui{
		State: State{
			Model: Model{
				Files: []*File{},
			},
			Panels: Panels{
				Files: FilesPanel{
					SelectedIdx: -1,
				},
			},
		},
		StagingCache: make(map[string]string),
	}
}

// RefreshFiles updates the file list while preserving selection and staging state
func (gui *Gui) RefreshFiles(newFiles []*File) {
	// 1. Store the currently selected file's path
	var selectedFilePath string
	if gui.State.Panels.Files.SelectedIdx >= 0 && gui.State.Panels.Files.SelectedIdx < len(gui.State.Model.Files) {
		selectedFilePath = gui.State.Model.Files[gui.State.Panels.Files.SelectedIdx].Name
	}

	// 2. Update the files list
	gui.State.Model.Files = newFiles

	// 3. Search for selectedFilePath in the new list to preserve selection
	found := false
	if selectedFilePath != "" {
		for i, file := range newFiles {
			if file.Name == selectedFilePath {
				gui.State.Panels.Files.SelectedIdx = i
				found = true
				break
			}
		}
	}

	// If the previously selected file is no longer in the list, reset selection
	if !found {
		if len(newFiles) > 0 {
			gui.State.Panels.Files.SelectedIdx = 0
		} else {
			gui.State.Panels.Files.SelectedIdx = -1
		}
	}
}

func main() { 
	fmt.Println("Hello, Bounty Hunter!")

	// Initialize GUI
	gui := NewGui()

	// Setup initial state
	gui.State.Model.Files = []*File{
		{Name: "file1.txt", HasStagedChanges: false},
		{Name: "file2.txt", HasStagedChanges: true},
		{Name: "file3.txt", HasStagedChanges: false},
	}
	gui.State.Panels.Files.SelectedIdx = 1 // file2.txt is selected
	gui.StagingCache["file2.txt"] = "hunk1:staged"

	fmt.Printf("Initial Selection: Index %d (%s)\n", gui.State.Panels.Files.SelectedIdx, gui.State.Model.Files[gui.State.Panels.Files.SelectedIdx].Name)

	// Simulate external refresh where file list changes but file2.txt is still present
	newFiles := []*File{
		{Name: "file3.txt", HasStagedChanges: false},
		{Name: "file2.txt", HasStagedChanges: true},
		{Name: "file4.txt", HasStagedChanges: false},
	}

	gui.RefreshFiles(newFiles)

	fmt.Printf("Post-Refresh Selection: Index %d (%s)\n", gui.State.Panels.Files.SelectedIdx, gui.State.Model.Files[gui.State.Panels.Files.SelectedIdx].Name)
	fmt.Printf("Staging Cache for file2.txt: %s\n", gui.StagingCache["file2.txt"])
}
