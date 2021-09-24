class Series
  attr_reader :digits

  def initialize(digits)
    @digits = digits
  end
  
  def slices(n)
    raise ArgumentError.new('n cannot be bigger than digits length') if n > digits.length
    digits.slice(0..digits.length-n).chars.map.with_index { |_, i| digits.slice(i, n) }
  end
end
