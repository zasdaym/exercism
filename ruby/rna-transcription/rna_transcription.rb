class Complement
  def self.of_dna(dna)
    result = dna.split("").map { |d| transform_dna(d) }
    result.join("")
  end
end

def transform_dna(dna)
  codes = {
    'G' => 'C',
    'C' => 'G',
    'T' => 'A',
    'A' => 'U'
  }
  
  if codes.has_key? dna
    return codes[dna]
  else
    return ""
  end
end
