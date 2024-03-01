defmodule CollatzConjecture do
  @doc """
  calc/1 takes an integer and returns the number of steps required to get the
  number to 1 when following the rules:
    - if number is odd, multiply with 3 and add 1
    - if number is even, divide by 2
  """
  @spec calc(input :: pos_integer()) :: non_neg_integer()
  def calc(number) when is_number(number) and number > 0 do
    cond do
      number === 1 -> 0
      rem(trunc(number), 2) === 1 -> 1 + calc(number * 3 + 1)
      rem(trunc(number), 2) === 0 -> 1 + calc(div(number, 2))
    end
  end
end
