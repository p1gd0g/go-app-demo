package main

import (
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	app.Route("/", &index{})
	app.Run()
}

type index struct {
	app.Compo
}

func (h *index) Render() app.UI {
	return app.Div().Body(
		&topMenu{},
	)
}

type topMenu struct {
	app.Compo
}

func (h *topMenu) Render() app.UI {
	return app.Div().Body(
		app.Div().Body(
			app.A().Href("https://www.oschina.net").Body(app.Text("首页")),
			app.A().Href("https://www.oschina.net/project").Body(app.Text("开源软件")),
			app.A().Href("https://www.oschina.net/question").Body(app.Text("问答")),
			app.A().Href("https://www.oschina.net/tweets").Body(app.Text("动弹")),
			app.A().Href("https://www.oschina.net/blog").Body(app.Text("博客")),
			app.A().Href("https://www.oschina.net/translate").Body(app.Text("翻译")),
			app.A().Href("https://www.oschina.net/news").Body(app.Text("资讯")),
			app.A().Href("https://gitee.com/?utm_source=oschina&utm_medium=link-index&utm_campaign=home").Body(app.Text("码云")),
			app.A().Href("https://zb.oschina.net/projects/list.html").Body(app.Text("众包")),
			app.A().Href("https://www.oschina.net/event").Body(app.Text("活动")),
			app.A().Href("https://www.oschina.net/columns").Body(app.Text("专区")),
		),
		app.Div().TabIndex(0).Body(
			app.Text("更多"),
			app.I().TabIndex(0).Body(
				app.Div().TabIndex(-1).Body(),
			),
			app.Div().TabIndex(-1).Body(
				app.A().Href("https://www.oschina.net/event/ych").Body(
					app.Text("源创会"),
				),
				app.A().Href("https://job.oschina.net/").Body(
					app.Text("求职/招聘"),
				),
				app.A().Href("https://www.oschina.net/question/topic/masteronline").Body(
					app.Text("高手问答"),
				),
				app.A().Href("https://www.oschina.net/question/topic/osc-interview").Body(
					app.Text("开源访谈"),
				),
				app.A().Href("https://my.oschina.net/editorial-story").Body(
					app.Text("周刊"),
				),
				app.A().Href("https://www.oschina.net/company").Body(
					app.Text("公司开源导航页"),
				),
			),
		),
		app.Div().Body(),
	)
}
