package dingtalk

type DingTalkMarkdown struct {
	MsgType  string    `json:"msgtype"`
	At       *At       `json:"at"`
	Markdown *Markdown `json:"markdown"`
}

type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func NewDingTalkMarkdown() DingTalkMarkdown {
	return DingTalkMarkdown{
		MsgType: "markdown",
		At: &At{
			IsAtAll:   false,
			AtMobiles: []string{},
		},
		Markdown: &Markdown{
			Title: "",
			Text:  "",
		},
	}
}

func (md *DingTalkMarkdown) SetTitle(title string) {
	md.Markdown.Title = title
}

func (md *DingTalkMarkdown) SetText(text string) {
	md.Markdown.Text = text
}

func (md *DingTalkMarkdown) SetIsAtAll(isAtAll bool) {
	md.At.IsAtAll = isAtAll
}

func (md *DingTalkMarkdown) SetAtMobiles(atMobiles []string) {
	md.At.AtMobiles = atMobiles
}
