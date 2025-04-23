module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    id = city[0, 4].upcase
    :"#{id}"
  end

  def self.get_terminal(ship_identifier)
    cargo = ship_identifier[0, 3].upcase

    if cargo == "OIL" || cargo == "GAS"
      :A
    else
      :B
    end
  end
end
