class Triangle
  def initialize(sides)
    @sides = sides.sort
  end

  def equilateral?
    @sides.uniq.size == 1 && @sides[0] != 0
  end

  def isosceles?
    @sides[1] == @sides[2] && self.is_triangle
  end

  def scalene?
    @sides.uniq.size == 3 && self.is_triangle
  end

  def is_triangle()
    @sides[0] + @sides[1] >= @sides[2]
  end
end
