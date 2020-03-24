class Matrix
  attr_reader :rows, :columns

  def initialize(matrix)
    rows = matrix.split("\n")
    @rows = rows.map { |row| row.split(" ").map { |c| c.to_i } }
    @columns = @rows.transpose
  end
end
