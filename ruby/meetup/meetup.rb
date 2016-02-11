require 'date'

# Not sure if monkey patching Date is a good practice here
class Date
  def first?
    day >= 1 && day <= 7
  end

  def second?
    day >= 8 && day <= 14
  end

  def third?
    day >= 15 && day <= 21
  end

  def fourth?
    day >= 22 && day <= 28
  end

  def teenth?
    day >= 13 && day <= 19
  end

  def last?
    day >= 23 && day <= 31
  end
end

# Meetup
class Meetup
  attr_accessor :month, :year, :weekday, :schedule
  def initialize(month, year)
    @month = month
    @year = year
  end

  def day(weekday, schedule)
    @weekday = weekday
    @schedule = schedule
    find_meetup_date
  end

  private

  # Start with last date in month for 'last' case
  def find_meetup_date
    Date.new(year, month).next_month.prev_day
        .downto(Date.new(year, month)) do |date|
      return date if correct_date?(date)
    end
  end

  def correct_date?(date)
    correct_weekday?(date) && correct_schedule?(date)
  end

  def correct_weekday?(date)
    date.send(weekday.to_s + '?')
  end

  def correct_schedule?(date)
    date.send(schedule.to_s + '?')
  end
end
