class LocomotiveEngineer
  def self.generate_list_of_wagons(*wagons)
    wagons
  end

  def self.fix_list_of_wagons(each_wagons_id, missing_wagons)
    [1] + missing_wagons + each_wagons_id[3..] + each_wagons_id[0..1]
  end

  def self.add_missing_stops(route, **stops)
    route[:stops] = stops.values
    route
  end

  def self.extend_route_information(route, more_route_information)
    route.merge(more_route_information)
  end
end
