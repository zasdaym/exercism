class Series
  def initialize(s)
    @s = s.chars
  end

  def slices(n)
    if @s.length < n
      raise ArgumentError.new("Number exceeds series length")
    end

    result = []
    (@s.length - n + 1).times { |i| result.append(@s[i, n].join) }
    result
  end
end
