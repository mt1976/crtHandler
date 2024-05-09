package language

// Page - Paging
var TxtPagingPrompt *Text = New("Choose (F)orward, (B)ack or (Q)uit")

// Support DatesTimes
var (
	OneWord            *Text = New("one")
	OneNumeric         *Text = New("1")
	Minutes            *Text = New("minutes")
	MinutesShort       *Text = New("mins")
	Hour               *Text = New("hour")
	HourShort          *Text = New("hr")
	MillisecondsShort  *Text = New("ms")
	ApplicationVersion *Text = New("StarTerm - Utilities 1.0 %s")
	ApplicationName    *Text = New("StarTerm")
	Error              *Text = New("ERROR ")
	Info               *Text = New("INFO ")
	Warning            *Text = New("WARNING ")
	Success            *Text = New("SUCCESS ")
	Hint               *Text = New("HINT ")
	Paging             *Text = New("Page %v of %v")
	MinMax             *Text = New("Min: %v Max: %v")
	ValidActions       *Text = New("valid actions [%v]")
)

var ApplicationHeader *Paragraph = NewParagraph([]string{
	"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
	"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
	"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
	"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
	"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
})

// General
var (
	Underline             *Text = New("-")
	LineConstructor       *Text = New("%s%s%s")
	MACAddressConstructor *Text = New("%v:%v:%v:%v:%v:%v")
	IPAddressConstructor  *Text = New("%v.%v.%v.%v")
	MinMaxLength          *Text = New("Text Length Min: %v Max: %v")
	Proceed               *Text = New("Proceed")
	SetPrompt             *Text = New("Please set a prompt for the page")
)

// FileChooser
var (
	FileChooserTitle        *Text      = New("File Chooser")
	FileChooserDescription  *Paragraph = NewParagraph([]string{"This menu shows the list of files available for maintenance.", "Select the file you wish to use. PLEASE BE CAREFUL!"})
	FileChooserPrompt       *Text      = New("Choose a file to use")
	FileChooserConfirmation *Text      = New("Choose (S)end or (Q)uit")
	FileChooserUserName     *Text      = New("User Name")
	FileChooserUserHome     *Text      = New("User Home")
	FileChooserDirectory    *Text      = New("Directory")
	FileChooserHeadType     *Text      = New("T")
	FileChooserHeadName     *Text      = New("Name")
	FileChooserHeadMode     *Text      = New("Mode")
	FileChooserHeadModified *Text      = New("Modified")
	FileChooserHeadSize     *Text      = New("Size")
	FileChooserNotAvailable *Text      = New("N/A")
	//head, " ", "T", "Name", "Mode", "Modified", "Size"
)

// Help Messages
var (
	HelpPageTitle        *Text = New("Help Page")
	HelpFor              *Text = New("Help for ")
	HelpPromptSinglePage *Text = New("Choose (Y)es when done")
	HelpPromptMultiPage  *Text = New("Choose (F)orward, (B)ack or (Y)es when done")
	HelpSupportedActions *Text = New("The following actions are supported:")
	HelpAutoGenerated    *Text = New("Autogenerated : ")
	//	HelpBullet           *Text = NewText("- ")
	HelpHint *Text = New("Help:")
)
