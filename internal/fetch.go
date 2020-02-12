package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func fetchUri(uri string) (resp *http.Response, err error) {
	resp, err = http.Get(uri)

	if err != nil {
		return
	}

	if resp.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf(`failed to fetch data from %s`, uri))
	}
	return
}

func (t *Tracker) trackAcolytes() (err error) {
	var (
		world map[string]interface{}
	)

	resp, err := fetchUri(WorldStateUrl)

	if err = json.NewDecoder(resp.Body).Decode(&world); err != nil {
		return err
	}
	defer resp.Body.Close()

	for _, a := range world["PersistentEnemies"].([]interface{}) {
		at := a.(map[string]interface{})["AgentType"].(string)

		n := getAcolyteNameFromType(at)
		if n == "" {
			return errors.New(fmt.Sprintf(`unknown Stalker agent "%s"`, n))
		}

		ldt := a.(map[string]interface{})["LastDiscoveredTime"].(map[string]interface{})["$date"].(map[string]interface{})["$numberLong"].(string)
		ms, err := StrMillisToTime(ldt)
		if err != nil {
			return err
		}
		t.acolytes[n].LastDiscoveredTime = ms.Unix()

		ldl := a.(map[string]interface{})["LastDiscoveredLocation"].(string)
		t.acolytes[n].LastDiscoveredLocation = t.nodes[ldl]

		t.acolytes[n].Discovered = a.(map[string]interface{})["Discovered"].(bool)
		t.acolytes[n].HealthPercent = a.(map[string]interface{})["HealthPercent"].(float64)
	}
	return err
}

func getAcolyteNameFromType(t string) string {
	switch t {
	case "/Lotus/Types/Enemies/Acolytes/StrikerAcolyteAgent":
		return "Angst"
	case "/Lotus/Types/Enemies/Acolytes/HeavyAcolyteAgent":
		return "Malice"
	case "/Lotus/Types/Enemies/Acolytes/RogueAcolyteAgent":
		return "Mania"
	case "/Lotus/Types/Enemies/Acolytes/AreaCasterAcolyteAgent":
		return "Misery"
	case "/Lotus/Types/Enemies/Acolytes/ControlAcolyteAgent":
		return "Torment"
	case "/Lotus/Types/Enemies/Acolytes/DuellistAcolyteAgent":
		return "Violence"
	default:
		return ""
	}
}
