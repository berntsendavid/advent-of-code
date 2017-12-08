// interface dir {
//   readonly x: number;
//   readonly y: number;
// }

// const north = new dir{0, 1};

// const directions = [{},{},{},{},]
// const intstructions = ['R1', 'L2', 'R1'];

class dir {
  x: number;
  y: number;
  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }
}

const north = new dir(0, 1);
