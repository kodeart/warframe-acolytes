package internal

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/happierall/l"
)

const (
	WorldStateUrl = "http://content.warframe.com/dynamic/worldState.php"
	MissionNodes  = "https://raw.githubusercontent.com/WFCD/warframe-worldstate-data/master/data/solNodes.json"
)

var (
	banner = "Warframe Acolytes Tracker"
)

// NewTracker creates a new instant of Tracker.
func NewTracker(refresh uint, silent, notify bool) (*Tracker, error) {
	t := new(Tracker)
	cwd, _ := os.Getwd()
	t.cwd = cwd
	t.silent = silent
	t.notify = notify

	t.acolytes = make(map[string]*Acolyte)
	t.nodes = make(map[string]Node)

	d := time.Duration(refresh)
	t.timer = new(Timer)
	t.timer.duration = d
	t.timer.end = time.Now().Add(d * time.Second).Round(0)

	t.initAcolytes()
	err := t.loadNodes()

	return t, err
}

// Run the tracker.
func (t *Tracker) Run() {
	if err := ui.Init(); err != nil {
		l.Errorf(`failed to initialize the termui: %v`, err)
	}
	defer ui.Close()

	t.trackAcolytes()

	// ui
	ban := t.newBanner()
	tbl := t.newTable()

	draw := func(ch <-chan time.Time) {
		t.loop(tbl, ch)
		ban.Text = fmt.Sprintf("%s in %v", banner, t.timer.status)
		ui.Render(ban, tbl)
	}

	draw(nil)
	ev := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	for {
		select {
		case e := <-ev:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			draw(ticker)
		}
	}
}

func (t *Tracker) loop(tbl *widgets.Table, ch <-chan time.Time) {
	t.timer.status = "- refreshing..."
	now := time.Now()

	if !now.After(t.timer.end) {
		t.timer.left = time.Duration(math.Floor(t.timer.end.Sub(now).Seconds())) * time.Second
		t.timer.status = fmt.Sprintf("%v", t.timer.left)
	} else {
		t.timer.end = now.Add(t.timer.duration * time.Second).Round(0)
		t.trackAcolytes()
		t.refreshTable(tbl)
	}
}

func (t *Tracker) newBanner() *widgets.Paragraph {
	b := widgets.NewParagraph()
	b.Text = banner
	b.Border = false
	b.SetRect(0, 0, 75, 5)
	return b
}

func (t *Tracker) newTable() *widgets.Table {
	tbl := widgets.NewTable()
	tbl.ColumnWidths = []int{10, 30, 30, 10}
	tbl.BorderStyle = ui.NewStyle(ui.ColorRed)
	tbl.RowSeparator = true
	tbl.FillRow = true
	tbl.SetRect(0, 2, 85, 17)

	tbl.Rows = append(tbl.Rows, []string{"Name", "Location", "Mission type", "Health"})
	tbl.RowStyles[0] = ui.NewStyle(ui.ColorClear, ui.ColorClear, ui.ModifierBold)

	for n, _ := range AcolyteNames {
		tbl.Rows = append(tbl.Rows, []string{n, "--", "--", fmt.Sprintf("%.2f %%", 0.00)})
	}
	t.refreshTable(tbl)
	return tbl
}

func (t *Tracker) refreshTable(tbl *widgets.Table) {
	for n, i := range AcolyteNames {
		a := t.acolytes[n]
		lo := "--"
		mt := "--"

		tbl.RowStyles[i] = ui.NewStyle(ui.ColorClear, ui.ColorClear, ui.ModifierClear)

		if a.Discovered {
			lo = a.LastDiscoveredLocation.Name
			mt = fmt.Sprintf("%s (%s)", a.LastDiscoveredLocation.MissionType, a.LastDiscoveredLocation.Enemy)
			tbl.RowStyles[i] = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
		}

		if !a.Notified() {
			ix := strings.LastIndex(a.AgentType, "/")
			icon := t.cwd + "/images" + a.AgentType[ix:] + ".png"

			if t.notify {
				if !t.silent {
					beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
				}
				beeep.Notify(a.Name, fmt.Sprintf("@%s - %s (%s)",
					lo, a.LastDiscoveredLocation.MissionType, a.LastDiscoveredLocation.Enemy), icon)
			}

		}

		tbl.Rows[i] = []string{a.Name, lo, mt, fmt.Sprintf("%.2f %%", a.HealthPercent*100)}
	}
}
