package main

/*
import (
	"fmt"
)
*/
func make_single_card(value int) Cards {
	values := []int{value}
	return make_cards_from_values(values)
}

func make_pair_cards(value int) Cards {
	values := []int{value, value}
	return make_cards_from_values(values)
}

func make_triple_cards(value int) Cards {
	values := []int{value, value, value}
	return make_cards_from_values(values)
}

func make_quadruple_cards(value int) Cards {
	values := []int{value, value, value, value}
	return make_cards_from_values(values)
}

func is_continous(cards Cards) bool {
	values := cards.values

	for i := 1; i < len(values); i++ {
		if values[i]-values[i-1] != 1 {
			return false
		}
	}

	return true
}

func get_pattern(cards Cards) Pattern {
	var single_amount int
	var pair_amount int
	var triple_amount int
	var quadruple_amount int

	var singles Cards
	var pairs Cards
	var triples Cards
	var quadruples Cards

	for _, value := range VALUES {
		if cards.contain(make_quadruple_cards(value)) {
			quadruple_amount += 1
			quadruples.add(make_single_card(value))
		} else if cards.contain(make_triple_cards(value)) {
			triple_amount += 1
			triples.add(make_single_card(value))
		} else if cards.contain(make_pair_cards(value)) {
			pair_amount += 1
			pairs.add(make_single_card(value))
		} else if cards.contain(make_single_card(value)) {
			single_amount += 1
			singles.add(make_single_card(value))
		}
	}

	has_pattern := false
	boss_card_value := 0

	if single_amount == 1 &&
		pair_amount == 0 &&
		triple_amount == 0 &&
		quadruple_amount == 0 {
		has_pattern = true
		boss_card_value = singles.values[singles.size()-1]

	}
	if single_amount == 0 &&
		pair_amount == 1 &&
		triple_amount == 0 &&
		quadruple_amount == 0 {
		has_pattern = true
		boss_card_value = pairs.values[pairs.size()-1]
	}
	if single_amount == 0 &&
		pair_amount == 0 &&
		triple_amount == 0 &&
		quadruple_amount == 1 {
		has_pattern = true
		boss_card_value = quadruples.values[quadruples.size()-1]
	}
	if single_amount >= 5 &&
		pair_amount == 0 &&
		triple_amount == 0 &&
		quadruple_amount == 0 &&
		is_continous(singles) {
		has_pattern = true
		boss_card_value = singles.values[singles.size()-1]
	}
	if single_amount == 0 &&
		pair_amount >= 3 &&
		triple_amount == 0 &&
		quadruple_amount == 0 &&
		is_continous(pairs) {
		has_pattern = true
		boss_card_value = pairs.values[pairs.size()-1]
	}
	if single_amount == 0 &&
		pair_amount == 0 &&
		triple_amount > 0 &&
		quadruple_amount == 0 &&
		is_continous(triples) {
		has_pattern = true
		boss_card_value = triples.values[triples.size()-1]
	}
	if single_amount > 0 &&
		pair_amount == 0 &&
		triple_amount > 0 &&
		quadruple_amount == 0 &&
		is_continous(triples) &&
		single_amount == triple_amount {
		has_pattern = true
		boss_card_value = triples.values[triples.size()-1]
	}
	if single_amount == 0 &&
		pair_amount > 0 &&
		triple_amount > 0 &&
		quadruple_amount == 0 &&
		is_continous(triples) &&
		single_amount == triple_amount {
		has_pattern = true
		boss_card_value = triples.values[triples.size()-1]
	}
	if single_amount == 2 &&
		pair_amount == 0 &&
		triple_amount == 0 &&
		quadruple_amount == 1 {
		has_pattern = true
		boss_card_value = quadruples.values[quadruples.size()-1]
	}
	if single_amount == 0 &&
		pair_amount == 2 &&
		triple_amount == 0 &&
		quadruple_amount == 1 {
		has_pattern = true
		boss_card_value = quadruples.values[quadruples.size()-1]
	}
	// nuke
	if single_amount == 2 &&
		pair_amount == 0 &&
		triple_amount == 0 &&
		quadruple_amount == 0 &&
		cards.values[0] == 16 &&
		cards.values[1] == 18 {
		has_pattern = true
		boss_card_value = singles.values[singles.size()-1]
	}

	var pattern Pattern
	if has_pattern {
		pattern.single_amount = single_amount
		pattern.pair_amount = pair_amount
		pattern.triple_amount = triple_amount
		pattern.quadruple_amount = quadruple_amount

		pattern.boss_card_value = boss_card_value
	}
	return pattern
}

type Pattern struct {
	single_amount    int
	pair_amount      int
	triple_amount    int
	quadruple_amount int

	boss_card_value int
}

func (pattern *Pattern) not_a_pattern() bool {
	if pattern.single_amount+pattern.pair_amount+pattern.triple_amount+pattern.quadruple_amount == 0 {
		return true
	} else {
		return false
	}
}

func (pattern *Pattern) is_bomb() bool {
	if pattern.single_amount == 0 &&
		pattern.pair_amount == 0 &&
		pattern.triple_amount == 0 &&
		pattern.quadruple_amount == 1 {
		return true
	} else {
		return false
	}
}
func (pattern *Pattern) is_nuke() bool {
	if pattern.single_amount == 2 {
		return true
	} else {
		return false
	}
}

func (pattern *Pattern) same_pattern(another_pattern Pattern) bool {
	if pattern.single_amount == another_pattern.single_amount &&
		pattern.pair_amount == another_pattern.pair_amount &&
		pattern.triple_amount == another_pattern.triple_amount &&
		pattern.quadruple_amount == another_pattern.quadruple_amount {
		return true
	}
	return false
}

func (pattern *Pattern) bigger_than(another_pattern Pattern) bool {
	if pattern.same_pattern(another_pattern) &&
		pattern.boss_card_value > another_pattern.boss_card_value {
		return true
	} else if pattern.is_bomb() && !another_pattern.is_nuke() {
		return true
	} else if pattern.is_nuke() {
		return true
	} else {
		return false
	}
}
