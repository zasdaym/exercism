class HighScores
  attr_reader :scores

  def initialize(scores)
    @scores = scores
    @sorted_scores = scores.sort.reverse
  end

  def personal_best
    @sorted_scores.first
  end

  def latest
    scores[-1]
  end

  def latest_is_personal_best?
    latest == personal_best
  end

  def personal_top_three
    @sorted_scores.first(3)
  end
end
