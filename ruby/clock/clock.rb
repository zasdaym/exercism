class Clock
  private

  def initialize(hour: 0, minute: 0)
    @hours = hour
    @minutes = minute
    self.normalize!
  end

  def normalize!
    @hours = (minutes / 60 + hours) % 24
    @minutes = minutes % 60
  end

  public

  attr_reader :hours, :minutes

  def to_s
    "#{'%02d' % hours}:#{'%02d' % minutes}"
  end

  def +(other)
    @hours += other.hours
    @minutes += other.minutes
    self.normalize!
    self
  end

  def -(other)
    @hours -= other.hours
    @minutes -= other.minutes
    self.normalize!
    self
  end

  def ==(other)
    hours == other.hours && minutes == other.minutes
  end
end
