# ETL
class ETL
  def self.transform(old)
    results = {}
    old.each do |point_val, letters|
      letters.each do |letter|
        results[letter.downcase] = point_val
      end
    end
    results
  end
end
