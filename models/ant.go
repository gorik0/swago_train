package models

type AntCollony struct {
	District string `json:"District"`
	Quantity int `json:"quaQtity"`


}
//ColonyList example
type ColonyList struct {
	PopulationList []int `json:"population_list" example:"00000"`

}

var Ants = []*AntCollony{
	&AntCollony{
	District: "north",
	Quantity: 100,
},
&AntCollony{
	District: "south",
	Quantity: 10,
},
&AntCollony{
	District: "west",
	Quantity: 50,
},
&AntCollony{
	District: "moon_located",
	Quantity: 10009,
},
}