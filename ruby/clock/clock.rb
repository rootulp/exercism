class Clock

  def self.at(hours, mins=0)
    Clock.new(hours, mins)
  end

  attr_reader :hours, :mins
  def initialize(hours, mins)
    @hours = hours
    @mins = mins
  end

  def to_s
    format(@hours) + ":" + format(@mins)
  end

  def + (time_added)
    @mins += time_added
    fixup
  end

  def - (time_subtracted)
    @mins -= time_subtracted
    fixup
  end

  def == (clock)
    self.hours != clock.hours || self.mins != clock.mins ? false : true
  end

  private

  def format(time)
    time = time.to_s
    time.length == 1 ? time.prepend("0") : time
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