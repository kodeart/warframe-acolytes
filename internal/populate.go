package internal

import (
	"encoding/json"
	"time"
)

func (t *Tracker) initAcolytes() {
	t.acolytes["Angst"] = &Acolyte{
		Name:      "Angst",
		AgentType: "/Lotus/Types/Enemies/Acolytes/StrikerAcolyteAgent",
		Mods: map[string]float64{
			"Body Count":            51.52,
			"Repeater Clip":         22.22,
			"Spring-Loaded Chamber": 22.22,
			"Pressurized Magazine":  4.04,
		},
	}

	t.acolytes["Malice"] = &Acolyte{
		Name:      "Malice",
		AgentType: "/Lotus/Types/Enemies/Acolytes/HeavyAcolyteAgent",
		Mods: map[string]float64{
			"Focused Defence":     51.52,
			"Guided Ordnance":     22.22,
			"Targeting Subsystem": 22.22,
			"Narrow Barrel":       4.04,
		},
	}

	t.acolytes["Mania"] = &Acolyte{
		Name:      "Mania",
		AgentType: "/Lotus/Types/Enemies/Acolytes/RogueAcolyteAgent",
		Mods: map[string]float64{
			"Catalyzer Link":     51.52,
			"Embedded Catalyzer": 22.22,
			"Weeping Wounds":     22.22,
			"Nano-Applicator":    4.04,
		},
	}

	t.acolytes["Misery"] = &Acolyte{
		Name:      "Misery",
		AgentType: "/Lotus/Types/Enemies/Acolytes/AreaCasterAcolyteAgent",
		Mods: map[string]float64{
			"Focused Defense":       25.38,
			"Body Count":            8.57,
			"Catalyzer Link":        8.57,
			"Hydraulic Crosshairs":  8.57,
			"Shrapnel Shot":         8.57,
			"Bladed Rounds":         3.70,
			"Blood Rush":            3.70,
			"Embedded Catalyzer":    3.70,
			"Guided Ordnance":       3.70,
			"Laser Sight":           3.70,
			"Repeater Clip":         3.70,
			"Sharpened Bullets":     3.70,
			"Spring-Loaded Chamber": 3.70,
			"Targeting Subsystem":   3.70,
			"Weeping Wounds":        3.70,
			"Argon Scope":           0.67,
			"Maiming Strike":        0.67,
			"Nano-Applicator":       0.67,
			"Narrow Barrel":         0.67,
			"Pressurized Magazine":  0.67,
		},
	}

	t.acolytes["Torment"] = &Acolyte{
		Name:      "Torment",
		AgentType: "/Lotus/Types/Enemies/Acolytes/ControlAcolyteAgent",
		Mods: map[string]float64{
			"Hydraulic Crosshairs": 51.52,
			"Blood Rush":           22.22,
			"Laser Sight":          22.22,
			"Argon Scope":          4.04,
		},
	}

	t.acolytes["Violence"] = &Acolyte{
		Name:      "Violence",
		AgentType: "/Lotus/Types/Enemies/Acolytes/DuellistAcolyteAgent",
		Mods: map[string]float64{
			"Shrapnel Shot":     51.52,
			"Bladed Rounds":     22.22,
			"Sharpened Bullets": 22.22,
			"Maiming Strike":    4.04,
		},
	}
}

func (a *Acolyte) Notified() bool {
	if !a.Discovered {
		a.notified = false
		return true
	}

	if a.notified {
		return true
	}

	now := time.Now().Unix()
	end := a.LastDiscoveredTime + 30

	if (end < now && !a.notified) || (end > now) {
		a.notified = true
		return false
	}

	return true
}

func (t *Tracker) loadNodes() error {
	resp, err := fetchUri(MissionNodes)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(&t.nodes)
}
