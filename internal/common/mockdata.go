package common

import "github.com/ralugr/filter-service/internal/model"

var MockMessage1 = model.Message{
	UID: "1234",
	Body: "![First image](path/to/image.png \"Text to show on mouseover\") <!--state: Rejected--> \n" +
		"![Second image (tea)](path/to/image.png \"Text-to-show-on-mouseover\") <!--state: Accepted   --> \n" +
		"![Third image, added](path/to/image.png       \"Text-to-show-on-mouseover\") <!--    state:    rejected-->\n" +
		"Some basic text\n" +
		"![Forth image    ](path/to/image.png \"Text, to, show, on, mouseover\") <!--state: invalid-->\n" +
		"![Fifth image](path/to/image.png \"Text to show on mouseover\") <!--state: accepted-->",
	State: model.Invalid}

var MockMessage2 = model.Message{
	UID: "3456",
	Body: "#   Starts with heading 1\n" +
		"Just a message without any images. Just a message without any images. Just a message without any images." +
		"[Google link] (\"https://google.com/\")",
	State: model.Invalid,
}
var MockMessage3 = model.Message{
	UID: "4567",
	Body: "# Starts with heading 1\n" +
		"![First image (tea)](path/to/image.png \"Text-to-show-on-mouseover\")" +
		"[Internal link] (\"/internal/path\")",
	State: model.Invalid,
}
var MockMessage4 = model.Message{
	UID: "5678",
	Body: "#### Starts with heading 4 \n" +
		"![First [image]    ](path/to/image.png \"Text, to, show, on, mouseover\") <!--state: acc   -->",
	State: model.Invalid,
}

var MockMessage5 = model.Message{
	UID: "5678",
	Body: "Starts with a paragraph\n" +
		"![]() <!--state: Rejected   -->\n" +
		"# Contains heading1",
	State: model.Invalid,
}

var MockMessage6 = model.Message{
	UID: "5678",
	Body: "# Heading\n" +
		" Paragraph \n" +
		"!(path/to/image) <!--state: queued   -->" +
		"[Google invalid link] (\"google.com/\")",
	State: model.Invalid,
}
