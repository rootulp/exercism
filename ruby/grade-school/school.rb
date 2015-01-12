class School

  def initialize
    @db = {}
  end

  def add(name, grade)
    if @db.has_key?(grade)
      @db[grade] << name
      @db[grade].sort!
    else
      @db[grade] = [name]
    end
  end

  def grade(grade)
    @db.fetch(grade, [])
  end

  def to_hash
    Hash[@db.sort]
  end

end