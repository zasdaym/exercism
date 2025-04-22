class AssemblyLine
  CARS_PER_HOUR = 221

  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    CARS_PER_HOUR * @speed * success_rate
  end

  def working_items_per_minute
    production_rate_per_hour.to_i / 60
  end

  private

  def success_rate
    if @speed > 9
      0.77
    elsif @speed > 8
      0.8
    elsif @speed > 4
      0.9
    else
      1.0
    end
  end
end
