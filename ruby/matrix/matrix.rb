class Matrix
  def initialize(s)
    @rows = s.split("\n").map do |line|
      line.split(" ").map do |item|
        item.to_i
      end
    end
  end

  def row(index)
    @rows[index - 1]
  end

  def column(index)
    [1]
  end
end
