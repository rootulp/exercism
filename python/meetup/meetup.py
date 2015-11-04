from datetime import date
from calendar import Calendar

def meetup_day(year, month, weekday, schedule):
    candidates = [date for date in Calendar().itermonthdates(year, month)
                    if date.month == month and \
                    date.strftime('%A') == weekday ]

    if schedule == 'teenth':
        return find_teenth(candidates)
    else:
        return find(candidates, schedule)

def find_teenth(candidates):
    for date in candidates:
        if date.day >= 13 and date.day <= 19:
            return date

def find(candidates, schedule):
    index = -1 if schedule == 'last' else int(schedule[0]) - 1
    return candidates[index]

