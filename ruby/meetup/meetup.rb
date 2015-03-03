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
    if schedule == :last
      return find_last(weekday, schedule)
    end
    which = 1
    curr = Date.new(year, month)
    while curr.month == month
      if curr.wday == WEEKDAYS[weekday]
        if which == SCHEDULE[schedule]
          return curr
        else
          last = curr
        end
        which += 1
      end
      curr = curr.next
    end
  end

  def find_last(weekday, schedule)
    curr = Date.new(year, month).next_month.prev_day
    while curr.month == month
      if curr.wday == WEEKDAYS[weekday]
        return curr
      end
      curr = curr.prev_day
    end
  end
end
