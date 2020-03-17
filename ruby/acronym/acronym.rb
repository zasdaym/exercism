class Acronym
  def self.abbreviate(name)
    name.gsub(/[-_ ]+/, ' ').split(" ").map { |name| name[0].chr }.join.upcase
  end
end
