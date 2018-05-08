package abstractFactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	factory := HTMLFactory{}

	usYahoo := factory.CreateLink("Yahoo!", "http://www.yahoo.com")
	jaYahoo := factory.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp")

	tray := factory.CreateTray("Yahoo!")
	tray.AddItem(usYahoo)
	tray.AddItem(jaYahoo)

	page := factory.CreatePage("Title", "Author")
	page.AddContent(tray)

	output := page.Output()

	expect := `<html>
  <head>
    <title>
      Title
    </title>
  </head>
  <body>
    <h1>
      Title
    </h1>
    <ul>
      <li>
        Yahoo!
        <ul>
          <li>
            <a href="http://www.yahoo.com">
              Yahoo!
            </a>
          </li>
          <li>
            <a href="http://www.yahoo.co.jp">
              Yahoo!Japan
            </a>
          </li>
        </ul>
      </li>
    </ul>
    <hr>
    <address>
      Author
    </address>
  </body>
</html>`

	if output != expect {
		t.Errorf("Expect output to %s, but %s\n", expect, output)
	}
}
