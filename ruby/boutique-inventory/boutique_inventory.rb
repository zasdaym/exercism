class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    names = @items.map do |item|
      item[:name]
    end

    names.sort
  end

  def cheap
    @items.filter do |item|
      item[:price] < 30
    end
  end

  def out_of_stock
    @items.filter do |item|
      item[:quantity_by_size] == {}
    end
  end

  def stock_for_item(name)
    filtered = @items.filter do |item|
      item[:name] == name
    end

    filtered[0][:quantity_by_size]
  end

  def total_stock
    quantities = @items.map do |item|
      item[:quantity_by_size].reduce(0) { |result, (_, value)| result += value }
    end

    quantities.sum
  end
end
