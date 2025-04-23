module Blackjack
  def self.parse_card(card)
    case card
    when "ace" then 11
    when "two" then 2
    when "three" then 3
    when "four" then 4
    when "five" then 5
    when "six" then 6
    when "seven" then 7
    when "eight" then 8
    when "nine" then 9
    when "ten" then 10
    when "jack" then 10
    when "queen" then 10
    when "king" then 10
    else 0
    end
  end

  def self.card_range(card1, card2)
    total = parse_card(card1) + parse_card(card2)
    case total
    when 21 then "blackjack"
    when 17..20 then "high"
    when 12..16 then "mid"
    else "low"
    end
  end

  def self.first_turn(card1, card2, dealer_card)
    range = card_range(card1, card2)
    dealer_value = parse_card(dealer_card)

    if card1 == "ace" && card2 == "ace"
      "P"
    elsif range == "blackjack"
      if dealer_value < 10
        "W"
      else
        "S"
      end
    elsif range == "high"
      "S"
    elsif range == "mid" && dealer_value < 7
      "S"
    else
      "H"
    end
  end
end
