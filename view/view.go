package view

import (
	"../providers"
	"github.com/marcusolsson/tui-go"
)

func StartView(pro []providers.Provider) {
	jobs := tui.NewTable(4, 0)

	jobsScroll := tui.NewScrollArea(jobs)
	jobsScroll.SetAutoscrollToBottom(true)

	jobsBox := tui.NewVBox(jobsScroll)
	jobsBox.SetBorder(true)

	jobsTitles := tui.NewTable(4, 1)
	jobsTitles.AppendRow(
		tui.NewPadder(1, 0, tui.NewLabel("Title")),
		tui.NewPadder(1, 0, tui.NewLabel("Company")),
		tui.NewPadder(1, 0, tui.NewLabel("Type")),
		tui.NewPadder(1, 0, tui.NewLabel("Location")),
	)
	jobsTitles.SetSizePolicy(tui.Minimum, tui.Maximum)

	view := tui.NewVBox(jobsTitles, jobsBox)

	for _, p := range pro {
		err := p.RetrieveJobs(func(job *providers.Job) {
			jobs.AppendRow(
				tui.NewPadder(1, 0, tui.NewLabel(job.Title)),
				tui.NewPadder(1, 0, tui.NewLabel(job.Company)),
				tui.NewPadder(1, 0, tui.NewLabel(job.Type)),
				tui.NewPadder(1, 0, tui.NewLabel(job.Location)),
			)
		})
		if err != nil {
			panic(err)
		}
	}

	ui, err := tui.New(view)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Q", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
