// poker project main.go
package main

import (
	"fmt"
)

func print(info string) {
	fmt.Println(info)
}

func wait_for_input() string {
	input_string := ""

	fmt.Scanln(&input_string)

	return input_string
}

func main() {

NEW_GAME:
	print("***********")
	print("Game Start!")
	print("***********")

	deck := make_cards_from_values(DECK)
	deck.shuffle()

	john := deck.slice(0, 17)
	mary := deck.slice(17, 34)
	lord := deck.slice(34, 54)

	john.sort()
	mary.sort()
	lord.sort()

	var continuous_pass_times = 0
	var current_cards *Cards = nil
	var current_player *Cards = nil

	current_player = &lord

	for true {
		print("Cards in hand:")
		current_player.show()

		if current_cards != nil {
			print("Cards to beat:")
			current_cards.show()
		} else {
			print("Play what ever you want:")
		}

		for true {
			input_string := wait_for_input()
			if len(input_string) == 0 {
				if current_cards == nil {
					print("Invalid play: can't pass, try again!")
					continue
				} else {
					print("Pass")
					continuous_pass_times += 1
					if continuous_pass_times == 2 {
						continuous_pass_times = 0
						current_cards = nil
					}
					break
				}
			}

			card_names := string_to_names(input_string)
			if card_names == nil {
				print("Invalid input: try again!")
				continue
			}

			input_cards := make_cards_from_names(card_names)
			if !current_player.contain(input_cards) {
				print("Invalid play: play cards in your hand, try again!")
				continue
			}

			input_cards_pattern := get_pattern(input_cards)
			if input_cards_pattern.not_a_pattern() {
				print("Invalid play: not a pattern, try again!")
				continue
			}

			if current_cards != nil {
				current_cards_pattern := get_pattern(*current_cards)
				if !input_cards_pattern.bigger_than(current_cards_pattern) {
					print("Invalid play: play bigger than last player, try again!")
					continue
				}
			}

			current_player.remove(input_cards)
			if current_player.size() == 0 {
				print("You Win!")
				goto NEW_GAME
			}

			current_cards = &input_cards
			continuous_pass_times = 0
			break
		}

		if current_player == &lord {
			current_player = &john
		} else if current_player == &john {
			current_player = &mary
		} else if current_player == &mary {
			current_player = &lord
		} else {
			print("Error: change player failed, game exits!")
		}
		print("")
	}
}
