require 'date'

class Meetup
  attr_accessor :month, :year
  WEEKDAYS = {
    :sunday => 0,
    :monday => 1,
    :tuesday => 2,
    :wednesday => 3,
    :thursday => 4,
    :friday => 5,
    :saturday => 6
  }
  SCHEDULE = {
    :first => 1,
    :second => 2,
    :third => 3,
    :fourth => 4
  }
  def initialize(month, year)
    @month = month
    @year = year
  end

  def day(weekday, schedule)
    which = 1
    curr = Date.new(year, month)
    while curr.month == month
      if curr.wday == WEEKDAYS[weekday]
        p curr
        if which == SCHEDULE[schedule]
          return curr
        else
          last = curr
        end
        which += 1
      end
      curr = curr.next
    end

    if schedule == :last
      return last
    end

  end
end
