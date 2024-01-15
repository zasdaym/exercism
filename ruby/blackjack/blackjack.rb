module Blackjack
  def self.parse_card(card)
    case card
    when 'ace'
      11
    when 'two'
      2
    when 'three'
      3
    when 'four'
      4
    when 'five'
      5
    when 'six'
      6
    when 'seven'
      7
    when 'eight'
      8
    when 'nine'
      9
    when 'ten', 'jack', 'queen', 'king'
      10
    else
      0
    end
  end

  def self.card_range(card1, card2)
    score = parse_card(card1) + parse_card(card2)
    case
    when score > 20
      'blackjack'
    when score > 16
      'high'
    when score > 11
      'mid'
    when score > 3
      'low'
    end
  end

  def self.first_turn(card1, card2, dealer_card)
    case
    when card1 == 'ace' && card2 == 'ace'
      'P'
    when card_range(card1, card2) == 'blackjack' && parse_card(dealer_card) < 10
      'W'
    when card_range(card1, card2) == 'blackjack'
      'S'
    when card_range(card1, card2) == 'high'
      'S'
    when card_range(card1, card2) == 'mid' && parse_card(dealer_card) < 7
      'S'
    else
      'H'
    end
  end
end
