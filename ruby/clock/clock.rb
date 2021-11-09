class Clock
  private

  def initialize(hour: 0, minute: 0)
    @minute = (minute % 60) + (minute / 60 + hour) % 24 * 60
  end

  public

  attr_reader :minute

  def to_s
    "#{'%02d' % (minute / 60)}:#{'%02d' % (minute % 60)}"
  end

  def +(other)
    Clock.new(hour: 0, minute: minute + other.minute)
  end

  def -(other)
    Clock.new(hour: 0, minute: minute - other.minute)
  end

  def ==(other)
    minute == other.minute
  end
end
