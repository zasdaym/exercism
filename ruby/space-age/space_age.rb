class SpaceAge
  def initialize(seconds)
    @seconds = seconds
  end

  PLANETS = { 
    mercury: 0.2408467,
    venus: 0.61519726,
    mars: 1.8808158,
    jupiter: 11.862615,
    saturn: 29.447498,
    uranus: 84.016846,
    neptune: 164.79132
  }

  def on_earth
    @seconds / 60 / 60 / 24 / 365.25
  end

  def on_mercury
    on_earth / PLANETS[:mercury]
  end

  def on_venus
    on_earth / PLANETS[:venus]
  end
  
  def on_mars
    on_earth / PLANETS[:mars]
  end
  
  def on_jupiter
    on_earth / PLANETS[:jupiter]
  end

  def on_saturn
    on_earth / PLANETS[:saturn]
  end

  def on_uranus
    on_earth / PLANETS[:uranus]
  end

  def on_neptune
    on_earth / PLANETS[:neptune]
  end
end
