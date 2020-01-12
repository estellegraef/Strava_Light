/*
 * 2848869
 * 8089098
 * 3861852
 */

package pages

type Page struct {
	title            string
	customStyleSheet string
	jsFile           string
	headLine         string
}

//Provides the static strings for the index page
func NewIndex() Page {
	return Page{
		title:            "Strava",
		customStyleSheet: "../assets/css/items.css",
		jsFile:           "../assets/js/custom.js",
		headLine:         "Deine Aktivitäten",
	}
}

//Provides the static strings for the search page
func NewSearch() Page {
	return Page{
		title:            "Suche",
		customStyleSheet: "../assets/css/items.css",
		jsFile:           "../assets/js/custom.js",
		headLine:         "Suche",
	}
}

//Provides the static strings for the upload page
func NewUpload() Page {
	return Page{
		title:    "Hinzufügen",
		headLine: "Anlegen einer Aktivität",
	}
}

//Provides the static strings for the detail page - takes the sportType as Headline
func NewDetail(heading string) Page {
	return Page{
		title:            "Detail",
		customStyleSheet: "../assets/css/detail.css",
		jsFile:           "../assets/js/custom.js",
		headLine:         heading,
	}
}

//Provides the static strings for the edit page
func NewEdit() Page {
	return Page{
		title:    "Bearbeiten",
		headLine: "Bearbeiten einer Aktivität",
	}
}

//Returns the title of a given page
func (a Page) GetTitle() string {
	return a.title
}

//Returns the Stylesheet path of a given page
func (a Page) GetCustomStyleSheet() string {
	return a.customStyleSheet
}

//Checks if a given page has a custom stylesheet
func (a Page) HasCustomStyleSheet() bool {
	return len(a.customStyleSheet) != 0
}

//Checks if a given page has a custom JS File
func (a Page) HasJsFile() bool {
	return len(a.jsFile) != 0
}

//Returns the JS Filepath of a given page
func (a Page) GetJsFile() string {
	return a.jsFile
}

//Returns the headline of a given page
func (a Page) GetHeadline() string {
	return a.headLine
}
