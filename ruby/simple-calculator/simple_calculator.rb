class SimpleCalculator
  class UnsupportedOperation < StandardError
  end

  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  def self.calculate(a, b, operation)
    if !ALLOWED_OPERATIONS.include?(operation)
      raise UnsupportedOperation
    end

    if a.class == String or b.class == String
      raise ArgumentError
    end

    case operation
    when "+"
      result = a + b
    when "/"
      begin
        result = a / b
      rescue ZeroDivisionError => e
        return "Division by zero is not allowed."
      end
    when "*"
      result = a * b
    end

    "#{a} #{operation} #{b} = #{result}"
  end
end
