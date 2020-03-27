class Series
  def initialize(series)
    @series = series.chars
  end

  def slices(length)
    if @series.length < length
      raise ArgumentError.new("Slice length exceeds series length")
    end

    result = []
    @series.each_cons(length) { |slice| result.append(slice.join) }
    result
  end
end
