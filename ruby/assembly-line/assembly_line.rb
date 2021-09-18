class AssemblyLine
  private

  CARS_PER_HOUR = 221

  attr_reader :speed

  def initialize(speed)
    @speed = speed
  end

  def success_rate
    return 0.77 if speed >= 10
    return 0.8 if speed >= 9
    return 0.9 if speed >= 5
    1
  end

  public

  def production_rate_per_hour
    CARS_PER_HOUR * speed * success_rate
  end

  def working_items_per_minute
    production_rate_per_hour.to_i / 60
  end
end
