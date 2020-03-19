class Pangram
  def self.pangram?(sentence)
    sentence.upcase.gsub(/[^A-Z]/, "").chars.sort.join.squeeze == "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  end
end
