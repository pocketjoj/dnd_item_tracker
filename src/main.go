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
	Item      interface{} `json:"item"`
	ItemGroup interface{} `json:"itemGroup"`
}

// type Item struct {
// 	Name               string `json:"name"`
// 	Source             string `json:"source"`
// 	Page               string `json:"page"`
// 	Rarity             string `json:"rarity"`
// 	Attunement         string `json:"reqAttune"`
// 	AttunementInfo     string `json:"reqAttuneTags"`
// 	Wondrous           string `json:"wondrous"`
// 	BonusSpellAttack   string `json:"bonusSpellAttack"`
// 	BonusSpellSaveDC   int64  `json:"bonusSpellSaveDc"`
// 	Entries            string `json:"entries"`
// 	Weight             int64  `json:"weight"`
// 	Focus              string `json:"focus"`
// 	BaseItem           string `json:"baseItem"`
// 	Type               string `json:"type"`
// 	WeaponCategory     string `json:"weaponCategory"`
// 	Property           string `json:"property"`
// 	Damage1            string `json:"dmg1"`
// 	DamangeType        string `json:"dmgType"`
// 	BonusWeapon        string `json:"bonusWeapon"`
// 	Weapon             string `json:"weapon"`
// 	GrantsProficiency  string `json:"grantsProficiency"`
// 	Tier               string `json:"tier"`
// 	Loottables         string `json:"lootTables"`
// 	SRD                string `json:"srd"`
// 	Value              string `json:"value"`
// 	Recharge           string `json:"recharge"`
// 	Charges            string `json:"charges"`
// 	Tattoo             string `json:"tattoo"`
// 	Resist             string `json:"resist"`
// 	Detail1            string `json:"detail1"`
// 	Hasrefs            string `json:"hasRefs"`
// 	Crew               string `json:"crew"`
// 	VehicleAC          string `json:"vehAc"`
// 	VehicleHP          string `json:"vehHp"`
// 	VehicleSpeed       string `json:"vehSpeed"`
// 	CapacityPassenger  string `json:"capPassenger"`
// 	CapacityCargo      string `json:"capCargo"`
// 	ConditionImmunity  string `json:"conditionImmune"`
// 	Damage2            string `json:"dmg2"`
// 	Attachedspells     string `json:"attachedSpells"`
// 	Hasfluffimages     string `json:"hasFluffImages"`
// 	Additionalsources  string `json:"additionalSources"`
// 	Additionalentries  string `json:"additionalEntries"`
// 	Scftype            string `json:"scfType"`
// 	Ability            string `json:"ability"`
// 	AC                 string `json:"ac"`
// 	Range              string `json:"range"`
// 	Strength           string `json:"strength"`
// 	Stealth            string `json:"stealth"`
// 	Vulnerable         string `json:"vulnerable"`
// 	Curse              string `json:"curse"`
// 	BonusAC            string `json:"bonusAc"`
// 	Poison             string `json:"poison"`
// 	Poisontypes        string `json:"poisonTypes"`
// 	Immune             string `json:"immune"`
// 	Sentient           string `json:"sentient"`
// 	Containercapacity  string `json:"containerCapacity"`
// 	Packcontents       string `json:"packContents"`
// 	Atomicpackcontents string `json:"atomicPackContents"`
// 	Bonusweaponattack  string `json:"bonusWeaponAttack"`
// 	Bonussavingthrow   string `json:"bonusSavingThrow"`
// 	Staff              string `json:"staff"`
// 	Axe                string `json:"axe"`
// 	Age                string `json:"age"`
// 	Bonusweapondamage  string `json:"bonusWeaponDamage"`
// 	Carryingcapacity   string `json:"carryingCapacity"`
// 	Speed              string `json:"speed"`
// 	Copy               string `json:"_copy"`
// 	Othersources       string `json:"otherSources"`
// 	Sword              string `json:"sword"`
// 	Ammotype           string `json:"ammoType"`
// 	Vehdmgthresh       string `json:"vehDmgThresh"`
// 	Reqattunealt       string `json:"reqAttuneAlt"`
// 	BonusProficiency   int    `json:"bonusProficiencyBonus"`
// 	Dexteritymax       string `json:"dexterityMax"`
// 	Crewmin            string `json:"crewMin"`
// 	Crewmax            string `json:"crewMax"`
// 	Travelcost         string `json:"travelCost"`
// 	Shippingcost       string `json:"shippingCost"`
// 	Bonusabilitycheck  string `json:"bonusAbilityCheck"`
// 	Weightnote         string `json:"weightNote"`
// }

// type ItemList struct {
// 	Collection []Item
// }

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
