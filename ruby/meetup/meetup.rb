require 'date'

class Meetup
  attr_accessor :month, :year
  DAYS = {
    :sunday => 0,
    :monday => 1,
    :tuesday => 2,
    :wednesday => 3,
    :thursday => 4,
    :friday => 5,
    :saturday => 6
  }

  SCHED = {
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
    case schedule
    when :last
      return find_last(weekday, schedule)
    when :teenth
      return find_teenth(weekday, schedule)
    else
      return find_regular(weekday, schedule)
    end
  end

  private

  def find_last(weekday, schedule)
    curr = Date.new(year, month).next_month.prev_day
    while curr.month == month
      if curr.wday == DAYS[weekday]
        return curr
      end
      curr = curr.prev_day
    end
  end

  def find_teenth(weekday, schedule)
    curr = Date.new(year, month)
    while curr.month == month
      if curr.wday == DAYS[weekday] && is_teen?(curr)
        return curr
      end
      curr = curr.next_day
    end
  end

  def is_teen?(date)
    if date.day > 12 && date.day < 20
      return true
    else
      return false
    end
  end

  def find_regular(weekday, schedule)
    ordinal = 1
    curr = Date.new(year, month)
    while curr.month == month
      if curr.wday == DAYS[weekday]
        if ordinal == SCHED[schedule]
          return curr
        else
          ordinal += 1
        end
      end
      curr = curr.next
    end
  end
end
