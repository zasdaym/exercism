module Luhn
  def self.valid?(digits)
    trimmed = digits.delete(' ')
    return false if trimmed.length < 2

    sum = 0
    trimmed.reverse.chars.each_with_index do |digit, i|
      number = digit.to_i
      return false if number.to_s != digit

      number *= 2 if i % 2 == 1
      number -= 9 if number > 9
      sum += number
    end

    return sum % 10 == 0
  end
end
