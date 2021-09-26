class Phrase
  attr_reader :word_count

  def initialize(phrase)
    count_by_word = Hash.new(0)
    phrase.downcase.scan(/\b[\w']+\b/).each do |word|
      count_by_word[word] += 1
    end
    @word_count = count_by_word
  end
end
