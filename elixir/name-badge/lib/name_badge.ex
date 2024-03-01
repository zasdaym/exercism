defmodule NameBadge do
  def print(id, name, department) do
    parsed_department = (department || "OWNER") |> String.upcase()

    if !id do
      "#{name} - #{parsed_department}"
    else
      "[#{id}] - #{name} - #{parsed_department}"
    end
  end
end
