class Matrix
  attr_reader :rows, :columns

  def initialize(lines)
    @rows = lines.split("\n").map { |numbers| numbers.split.map { |number| number.to_i } }
    @columns = rows.transpose
  end
end
