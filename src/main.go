package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Outer struct {
	Meta      []string    `json:"meta"`
	Item      []Item      `json:"item"`
	ItemGroup interface{} `json:"itemGroup"`
}

type Item struct {
	Name               string                 `json:"name"`
	Source             string                 `json:"source"`
	Page               int64                  `json:"page"`
	Rarity             string                 `json:"rarity"`
	Attunement         interface{}            `json:"reqAttune"`
	AttunementInfo     interface{}            `json:"reqAttuneTags"`
	Wondrous           bool                   `json:"wondrous"`
	BonusSpellAttack   string                 `json:"bonusSpellAttack"`
	BonusSpellSaveDC   string                 `json:"bonusSpellSaveDc"`
	Entries            interface{}            `json:"entries"`
	Weight             float64                `json:"weight"`
	Focus              interface{}            `json:"focus"`
	BaseItem           string                 `json:"baseItem"`
	Type               string                 `json:"type"`
	WeaponCategory     string                 `json:"weaponCategory"`
	Property           interface{}            `json:"property"`
	Damage1            string                 `json:"dmg1"`
	DamangeType        string                 `json:"dmgType"`
	BonusWeapon        string                 `json:"bonusWeapon"`
	Weapon             bool                   `json:"weapon"`
	GrantsProficiency  bool                   `json:"grantsProficiency"`
	Tier               string                 `json:"tier"`
	Loottables         interface{}            `json:"lootTables"`
	SRD                interface{}            `json:"srd"`
	Value              float64                `json:"value"`
	Recharge           string                 `json:"recharge"`
	Charges            int64                  `json:"charges"`
	Tattoo             bool                   `json:"tattoo"`
	Resist             interface{}            `json:"resist"`
	Detail1            string                 `json:"detail1"`
	Hasrefs            bool                   `json:"hasRefs"`
	Crew               int64                  `json:"crew"`
	VehicleAC          int64                  `json:"vehAc"`
	VehicleHP          int64                  `json:"vehHp"`
	VehicleSpeed       float64                `json:"vehSpeed"`
	CapacityPassenger  int64                  `json:"capPassenger"`
	CapacityCargo      float64                `json:"capCargo"`
	ConditionImmunity  interface{}            `json:"conditionImmune"`
	Damage2            string                 `json:"dmg2"`
	Attachedspells     interface{}            `json:"attachedSpells"`
	Hasfluffimages     bool                   `json:"hasFluffImages"`
	Additionalsources  interface{}            `json:"additionalSources"`
	Additionalentries  interface{}            `json:"additionalEntries"`
	Scftype            string                 `json:"scfType"`
	Ability            map[string]interface{} `json:"ability"`
	AC                 int64                  `json:"ac"`
	Range              string                 `json:"range"`
	Strength           string                 `json:"strength"`
	Stealth            bool                   `json:"stealth"`
	Vulnerable         interface{}            `json:"vulnerable"`
	Curse              bool                   `json:"curse"`
	BonusAC            string                 `json:"bonusAc"`
	Poison             bool                   `json:"poison"`
	Poisontypes        interface{}            `json:"poisonTypes"`
	Immune             interface{}            `json:"immune"`
	Sentient           bool                   `json:"sentient"`
	Containercapacity  map[string]interface{} `json:"containerCapacity"`
	Packcontents       interface{}            `json:"packContents"`
	Atomicpackcontents interface{}            `json:"atomicPackContents"`
	Bonusweaponattack  string                 `json:"bonusWeaponAttack"`
	Bonussavingthrow   string                 `json:"bonusSavingThrow"`
	Staff              bool                   `json:"staff"`
	Axe                bool                   `json:"axe"`
	Age                string                 `json:"age"`
	Bonusweapondamage  string                 `json:"bonusWeaponDamage"`
	Carryingcapacity   int64                  `json:"carryingCapacity"`
	Speed              int64                  `json:"speed"`
	Copy               map[string]interface{} `json:"_copy"`
	Othersources       interface{}            `json:"otherSources"`
	Sword              bool                   `json:"sword"`
	Ammotype           string                 `json:"ammoType"`
	Vehdmgthresh       int64                  `json:"vehDmgThresh"`
	Reqattunealt       string                 `json:"reqAttuneAlt"`
	BonusProficiency   string                 `json:"bonusProficiencyBonus"`
	Dexteritymax       string                 `json:"dexterityMax"`
	Crewmin            int64                  `json:"crewMin"`
	Crewmax            int64                  `json:"crewMax"`
	Travelcost         int64                  `json:"travelCost"`
	Shippingcost       int64                  `json:"shippingCost"`
	Bonusabilitycheck  string                 `json:"bonusAbilityCheck"`
	Weightnote         string                 `json:"weightNote"`
}

func main() {
	log.SetFlags(log.Lshortfile)

	data, err := os.Open("src/items.json")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	var Items Outer

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &Items)
	if err != nil {
		log.Fatal(err)
	}

	// Writing data to a file to test what I'm getting... Resulting output is a map.
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	str := fmt.Sprintf("%v", Items.Item)
	_, err2 := f.WriteString(string(str))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
