function CircularBuffer(maxSize) {

  var readIndex = 0;
  var writeIndex = 0;
  var buffer = new Array(maxSize);

  function read() {
    if (!buffer[readIndex]) { throw new BufferEmptyException(); }
    var data = buffer[readIndex];
    buffer[readIndex] = null;
    readIndex = increment(readIndex);
    return data;
  }

  function write(data) {
    if (!data) { return; }
    buffer[writeIndex] = data;
    writeIndex = increment(writeIndex);
  }

  function clear() {
    readIndex = 0;
    writeIndex = 0;
    buffer = new Array(maxSize);
  }

  function increment(index) {
    index = (index + 1) % maxSize;
  }

  return {
    read: read,
    write: write,
    clear: clear
  }

};

function BufferEmptyException() {
  this.name = "BufferEmptyException";
  this.message = "Buffer is empty.";
};

function BufferFullException() {
  this.name = "BufferFullException";
  this.message = "Buffer is full.";
};

module.exports = {
  circularBuffer: function(capacity) {
    return new CircularBuffer(capacity);
  },

  bufferEmptyException: function() {
    return new BufferEmptyException();
  },

  bufferFullException: function() {
    return new BufferFullException();
  }
};
