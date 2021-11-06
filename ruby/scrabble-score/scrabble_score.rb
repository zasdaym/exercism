class Scrabble
  def self.score(word)
    return 0 if word == nil
    word.chars.reduce(0) { |sum, char| sum + @@score_by_letter.fetch(char.upcase, 0) }
  end

  def self.populate_scores()
    compressed_scores = {
      1 => "AEIOULNRST",
      2 => "DG",
      3 => "BCMP",
      4 => "FHVWY",
      5 => "K",
      8 => "JX",
      10 => "QZ",
    }

    @@score_by_letter = {}
    compressed_scores.each { |score, letters| letters.chars.each { |letter| @@score_by_letter[letter] = score } }
  end

  attr_reader :score

  def initialize(word)
    Scrabble.populate_scores if defined?(@@score_by_letter) == nil
    @score = Scrabble.score(word)
  end
end
