class Squares
  attr_reader :square_of_sum, :sum_of_squares, :difference
  
  def initialize(n)
    if n == 1
      @square_of_sum = 1
      @sum_of_squares = 1
      @difference = 0
    else
      @square_of_sum = ((1 + n) * n / 2.0) ** 2
      @sum_of_squares = n * (n + 1) * (2*n + 1) / 6
      @difference = @square_of_sum - @sum_of_squares
    end
  end
end
