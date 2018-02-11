require 'date'
require 'time'

# Gigasecond
class Gigasecond
  def self.from(date)
    (date.to_time + 10**9).to_date
  end
end
