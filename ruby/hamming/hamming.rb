module Hamming
  def self.compute(first, second)
    raise ArgumentError.new('different DNA strands length') if first.size != second.size
    first.chars.zip(second.chars).select { |a, b| a != b }.count
  end
end
