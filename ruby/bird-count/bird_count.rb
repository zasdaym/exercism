class BirdCount
  def self.last_week
    [0, 2, 5, 3, 7, 8, 4]
  end

  def initialize(birds_per_day)
    @birds = birds_per_day
  end

  def yesterday
    @birds[-2]
  end

  def total
    @birds.sum
  end

  def busy_days
    @birds.count do |bird|
      bird > 4
    end
  end

  def day_without_birds?
    @birds.any? do |bird|
      bird == 0
    end
  end
end
