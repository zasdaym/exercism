class Acronym
  def self.abbreviate(phrase)
    letters = phrase.split(/[- ]/).collect do |part|
      part[0]
    end
    letters.join.upcase
  end
end
