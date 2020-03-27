class SumOfMultiples
  def initialize(*factors)
    @factors = factors
  end

  def to(limit)
    1.upto(limit - 1).select { |n| multiple? n }.sum
  end

  def multiple?(number)
    @factors.find { |factor| number % factor == 0 }
  end
end
