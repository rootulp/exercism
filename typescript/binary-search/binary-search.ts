/**
 * Find performs a binary search for `item` in `array` and returns its index if
 * found. `find` throws an Error if value if item is not in the array. `array`
 * must be provided in sorted order.
 */
export function find(array: number[], item: number): number {
  return binarySearch(array, item, 0, array.length - 1)
}

function binarySearch(array: number[], item: number, left: number, right: number): number {
  if(right < left) {
    throw new Error('Value not in array')
  }
  if (right === left && array[right] !== item) {
    throw new Error('Value not in array')
  }

  const middle = Math.floor((right - left) / 2) + left
  const middleVal = array[middle]
  if (item === middleVal) {
    return middle
  }
  if (item > middleVal) {
    return binarySearch(array, item, middle + 1, right)
  }
  if (item < middleVal) {
    return binarySearch(array, item, left, middle - 1)
  }
  throw new Error("unreachable")
}
