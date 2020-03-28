class Series
  def initialize(series)
    @series = series.chars
  end

  def slices(length)
    if @series.length < length
      raise ArgumentError.new("Slice length exceeds series length")
    end

    @series.each_cons(length).map(&:join)
  end
end
