export function decodedValue(colors: string[]) {
  const [first, second] = colors
  return COLORS.indexOf(first) * 10 + COLORS.indexOf(second)
}

export const COLORS = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
]
