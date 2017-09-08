function transform(legacy_data) {
  var new_data = {};

  for (var key in legacy_data) {
    if (legacy_data.hasOwnProperty(key)) {
      add(Number(key), legacy_data[key]);
    }
  }

  function add(val, char_arr) {
    for (var i = 0; i < char_arr.length; i++) {
      new_data[char_arr[i].toLowerCase()] = val;
    }
  }

  return new_data;
}

module.exports = transform;
