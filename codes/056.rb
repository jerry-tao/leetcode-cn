# Definition for an interval.
# class Interval
#     attr_accessor :start, :end
#     def initialize(s=0, e=0)
#         @start = s
#         @end = e
#     end
# end

# @param {Interval[]} intervals
# @return {Interval[]}

def merge(intervals)
  return [] if intervals.empty?
  intervals = intervals.sort {|a, b| a.start <=> b.start }
  start, foo = intervals.first.start, intervals.first.end
  result = []
  result << intervals.first
  intervals.each do |item|
    if item.start <= result.last.end && item.end > result.last.end
      result.last.end = item.end
    end
    result << item if item.start > result.last.end
  end
  return result
end
