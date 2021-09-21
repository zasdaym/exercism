# frozen_string_literal: true

class LogLineParser
  def initialize(line)
    parts = line.split(":")
    raise "Invalid log line" if parts.size != 2
    @log_level = parts[0].tr("[]", "").swapcase
    @message = parts[1].strip
  end

  attr_reader :message, :log_level

  def reformat
    "#{message} (#{log_level})"
  end
end
