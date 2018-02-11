# Clock
class Clock
  def self.at(hours, mins = 0)
    Clock.new(hours, mins)
  end

  attr_reader :hours, :mins
  def initialize(hours, mins)
    @hours = hours
    @mins = mins
  end

  def to_s
    format(hours) + ':' + format(mins)
  end

  def +(other)
    @mins += other
    fixup
  end

  def -(other)
    @mins -= other
    fixup
  end

  def ==(other)
    hours == other.hours && mins == other.mins
  end

  private

  def format(time)
    time.to_s.length == 1 ? time.to_s.prepend('0') : time.to_s
  end

  def fixup
    overflow_mins
    wrap_midnight
    self
  end

  def overflow_mins
    hours_added, new_mins = mins.divmod(60)
    @hours += hours_added
    @mins = new_mins
  end

  def wrap_midnight
    if @hours <= 0
      @hours += 24
    elsif @hours >= 24
      @hours -= 24
    end
  end
end
