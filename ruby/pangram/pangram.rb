class Pangram
  def self.pangram?(sentence)
    sentence.tr('^A-Za-z', '').split('').uniq.sort.join('').downcase == 'abcdefghijklmnopqrstuvwxyz'
  end
end
