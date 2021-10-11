module Isogram
  def self.isogram?(phrase)
    processed = phrase.tr(' -', '').downcase.chars
    processed.uniq == processed
  end
end
