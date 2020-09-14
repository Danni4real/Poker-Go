package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const CARD_TYPE_AMOUNT = 15
const CARD_AMOUNT_IN_DECK = 54

var DECK = []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9, 10, 10, 10, 10, 11, 11, 11, 11, 12, 12, 12, 12, 14, 14, 14, 14, 16, 18}
var VALUES = [CARD_TYPE_AMOUNT]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 16, 18}
var NAMES = [CARD_TYPE_AMOUNT]string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2", "B", "R"}
var VALUE_NAME_MAP map[int]string
var NAME_VALUE_MAP map[string]int

func value_to_name(value int) string {
	if VALUE_NAME_MAP == nil {
		VALUE_NAME_MAP = make(map[int]string)
		for i := 0; i < CARD_TYPE_AMOUNT; i++ {
			VALUE_NAME_MAP[VALUES[i]] = NAMES[i]
		}
	}

	return VALUE_NAME_MAP[value]
}

func name_to_value(name string) int {
	if NAME_VALUE_MAP == nil {
		NAME_VALUE_MAP = make(map[string]int)
		for i := 0; i < CARD_TYPE_AMOUNT; i++ {
			NAME_VALUE_MAP[NAMES[i]] = VALUES[i]
		}
	}

	return NAME_VALUE_MAP[name]
}

func is_card_name(char byte) bool {
	for _, name := range NAMES {
		if char == name[0] && char != '1' {
			return true
		}
	}

	return false
}

func char_to_name(char byte) string {
	for _, name := range NAMES {
		if char == name[0] && char != '1' {
			return name
		}
	}

	return ""
}

func string_to_names(cards_string string) []string {
	var names []string
	cards_string_length := len(cards_string)

	for i := 0; i < cards_string_length; i++ {
		if cards_string[i] == '0' && i > 0 && cards_string[i-1] == '1' {
			// do nothing
		} else if cards_string[i] == '1' && i < cards_string_length-1 && cards_string[i+1] == '0' {
			names = append(names, "10")
		} else if is_card_name(cards_string[i]) {
			names = append(names, char_to_name(cards_string[i]))
		} else {
			return nil
		}
	}

	return names
}

func make_cards_from_values(values []int) Cards {
	var cards Cards

	for i := 0; i < len(values); i++ {
		cards.values = append(cards.values, values[i])
	}

	return cards
}

func make_cards_from_names(names []string) Cards {
	var values []int

	for _, name := range names {
		values = append(values, name_to_value(name))
	}

	return make_cards_from_values(values)
}

type Cards struct {
	values []int
}

func (cards *Cards) show() {
	for i := 0; i < len(cards.values); i++ {
		fmt.Print(value_to_name(cards.values[i]), " ")
	}
	fmt.Println()
}

func (cards *Cards) clone() Cards {
	return make_cards_from_values(cards.values)
}

func (cards *Cards) remove_value(value int) bool {

	for i := 0; i < len(cards.values); i++ {
		if cards.values[i] == value {
			cards.values = append(cards.values[:i], cards.values[i+1:]...)
			return true
		}
	}

	return false
}

//inner func, not a api, may leave cards at an inconsistent state!
func (cards *Cards) unatomic_remove(sub_cards Cards) bool {

	for _, value := range sub_cards.values {
		if cards.remove_value(value) == false {
			return false
		}
	}

	return true
}

func (cards *Cards) contain(sub_cards Cards) bool {
	cards_clone := cards.clone()

	if cards_clone.unatomic_remove(sub_cards) == true {
		return true
	}

	return false
}

// atomic remove, remove all or not remove at all
func (cards *Cards) remove(sub_cards Cards) bool {

	if cards.contain(sub_cards) {
		cards.unatomic_remove(sub_cards)
		return true
	}

	return false
}

func (cards *Cards) add(sub_cards Cards) {
	cards.values = append(cards.values, sub_cards.values...)
}

func (cards *Cards) size() int {
	return len(cards.values)
}

func (cards *Cards) sort() {
	sort.Ints(cards.values)
}

func (cards *Cards) shuffle() {
	values := []int{}
	random_index_list := rand.New(rand.NewSource(time.Now().Unix())).Perm(len(cards.values))

	for _, index := range random_index_list {
		values = append(values, cards.values[index])
	}

	cards.values = values
}

// include start, exclude end, will not change cards
func (cards *Cards) slice(start, end int) Cards {
	return make_cards_from_values(cards.values[start:end])
}

func (cards *Cards) to_string() string {
	cards_string := ""

	for _, value := range cards.values {
		cards_string += value_to_name(value)
	}

	return cards_string
}
