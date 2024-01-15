module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    city[0, 4].upcase.to_sym
  end

  def self.get_terminal(ship_identifier)
    material = ship_identifier[0, 3]
    ["OIL", "GAS"].include?(material) ? :A : :B
  end
end
