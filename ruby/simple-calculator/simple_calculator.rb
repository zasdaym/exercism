# frozen_string_literal: true

class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze
  DIVIDE_BY_ZERO_MESSAGE = 'Division by zero is not allowed.'
  private_constant :ALLOWED_OPERATIONS, :DIVIDE_BY_ZERO_MESSAGE

  class UnsupportedOperation < StandardError
    def message
      'Operation is not supported.'
    end
  end

  def self.calculate(first_operand, second_operand, operation)
    raise UnsupportedOperation.new unless ALLOWED_OPERATIONS.include?(operation)
    raise ArgumentError unless first_operand.is_a?(Integer)
    raise ArgumentError unless second_operand.is_a?(Integer)

    begin
      result = first_operand.public_send(operation, second_operand)
    rescue ZeroDivisionError => e
      return DIVIDE_BY_ZERO_MESSAGE
    end

    "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
