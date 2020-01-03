package pages

type Page struct {
	title            string
	customStyleSheet string
	jsFile           string
	headLine         string
}

func NewIndex() Page {
	return Page{
		title:            "Strava",
		customStyleSheet: "../assets/css/items.css",
		jsFile:           "../assets/js/iteyms.js",
		headLine:         "Deine Aktivitäten",
	}
}

func NewSearch() Page {
	return Page{
		title:            "Suche",
		customStyleSheet: "../assets/css/items.css",
		jsFile:           "../assets/js/items.js",
		headLine:         "Suche",
	}
}

func NewUpload() Page {
	return Page{
		title:    "Hinzufügen",
		headLine: "Anlegen einer Aktivität",
	}
}

func NewDetail(heading string) Page {
	return Page{
		title:            "Detail",
		customStyleSheet: "../assets/css/detail.css",
		headLine:         heading,
	}
}

func NewEdit() Page {
	return Page{
		title:    "Bearbeiten",
		headLine: "Bearbeiten einer Aktivität",
	}
}

func (a Page) GetTitle() string {
	return a.title
}

func (a Page) GetCustomStyleSheet() string {
	return a.customStyleSheet
}

func (a Page) HasCustomStyleSheet() bool {
	return len(a.customStyleSheet) != 0
}

func (a Page) HasJsFile() bool {
	return len(a.jsFile) != 0
}

func (a Page) GetJsFile() string {
	return a.jsFile
}

func (a Page) GetHeadline() string {
	return a.headLine
}
