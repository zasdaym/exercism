# frozen_string_literal: true

class Series
  attr_reader :digits

  def initialize(digits)
    @digits = digits
  end

  def slices(length)
    raise ArgumentError, 'slice length cannot be bigger than digits length' if length > digits.length

    digits.slice(0..digits.length - length).chars.map.with_index { |_, i| digits.slice(i, length) }
  end
end
