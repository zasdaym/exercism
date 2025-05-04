class DndCharacter
  attr_reader \
    :charisma,
    :dexterity,
    :hitpoints,
    :intelligence,
    :strength,
    :wisdom,
    :constitution,

  def self.modifier(constitution)
    (constitution - 10) / 2
  end

  def initialize
    @constitution = ability
    @charisma = ability
    @dexterity = ability
    @intelligence = ability
    @strength = ability
    @wisdom = ability
    @hitpoints = DndCharacter.modifier(@constitution) + 10
  end

  private

  def ability
    dice = Array.new(4) { rand(1..6) }
    dice.sum - dice.min
  end
end
