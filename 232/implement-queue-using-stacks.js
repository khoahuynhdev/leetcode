


/**
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */

 class MyQueue {
  _buffer;
  _tail = 0;
  _head = 0;
  _over = false;

  constructor() {
    this._buffer = Array(101).fill(0)
  }

  get length() {
    const num = this._over ? 1 : 0
    return num  * this._buffer.length + this._tail - this._head
  }

  empty() {
      return this.length === 0
  }

  peek() {
      return this._buffer[this._head]
  }

  pop() {
    if (this._head !== this._tail || this._over) {
      const val = this._buffer[this._head++]
      if (this._head === this._buffer.length) {
        this._head = 0
        this._over = false
      }
      return val
    }
  }

  push(val) {
    if (this._tail === this._head && this._over) {
      if (this._head++ === this._buffer.length) {
        this._head = 0
        this._over = false
      }
    }
    this._buffer[this._tail++] = val
    if (this._tail === this._buffer.length) {
      this._tail = 0
      this._over = true
    }
    return this.length
  }
}
