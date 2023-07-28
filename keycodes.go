package promptui

import "github.com/chzyer/readline"

// These runes are used to identify the commands entered by the user in the command prompt. They map
// to specific actions of promptui in prompt mode and can be remapped if necessary.
var (
	// KeyEnter is the default key for submission/selection.
	KeyEnter rune = readline.CharEnter

	// KeyCtrlH is the key for deleting input text.
	KeyCtrlH rune = readline.CharCtrlH

	// KeyPrev is the default key to go up during selection.
	KeyPrev        rune = readline.CharPrev
	KeyPrevDisplay      = "↑"

	// KeyNext is the default key to go down during selection.
	KeyNext        rune = readline.CharNext
	KeyNextDisplay      = "↓"

	// KeyBackward is the default key to page up during selection.
	KeyBackward        rune = readline.CharBackward
	KeyBackwardDisplay      = "←"

	// KeyForward is the default key to page down during selection.
	KeyForward        rune = readline.CharForward
	KeyForwardDisplay      = "→"
)

const (
	// KeyRefresh is the key to refresh the current status
	KeyRefresh = '\x01'

	// KeyCtrlE is the key to erase the search keywords
	KeyCtrlE = '\x05'

	// KeySoftEnter is the key to lock the search keywords
	KeySoftEnter = '\x1e'
)
