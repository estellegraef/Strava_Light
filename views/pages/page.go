package pages

type Page struct {
	title            string
	customStyleSheet string
	headLine         string
}

func NewIndex() Page {
	return Page{
		title:            "Strava",
		customStyleSheet: "../assets/css/items.css",
		headLine:         "Deine Aktivitäten",
	}
}

func NewSearch() Page {
	return Page{
		title:            "Suche",
		customStyleSheet: "../assets/css/items.css",
		headLine:         "Suche",
	}
}

func NewUpload() Page {
	return Page{
		title:    "Hinzufügen",
		headLine: "Anlegen einer Aktivität",
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

func (a Page) GetHeadline() string {
	return a.headLine
}
