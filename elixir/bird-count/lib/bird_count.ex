defmodule BirdCount do
  def today([]) do
    nil
  end

  def today(list) do
    hd(list)
  end

  def increment_day_count([]) do
    [1]
  end

  def increment_day_count([head | tail]) do
    [head + 1 | tail]
  end

  def has_day_without_birds?([]) do
    false
  end

  def has_day_without_birds?([head | tail]) do
    head === 0 or has_day_without_birds?(tail)
  end

  def total([]) do
    0
  end

  def total([head | tail]) do
    head + total(tail)
  end

  def busy_days([]) do
    0
  end

  def busy_days([head | tail]) do
    busy_day = head >= 5 && 1 || 0
    busy_day + busy_days(tail)
  end
end
