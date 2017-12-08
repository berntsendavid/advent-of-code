const instructions = [
  'R2',
  'L3',
  'R2',
  'R4',
  'L2',
  'L1',
  'R2',
  'R4',
  'R1',
  'L4',
  'L5',
  'R5',
  'R5',
  'R2',
  'R2',
  'R1',
  'L2',
  'L3',
  'L2',
  'L1',
  'R3',
  'L5',
  'R187',
  'R1',
  'R4',
  'L1',
  'R5',
  'L3',
  'L4',
  'R50',
  'L4',
  'R2',
  'R70',
  'L3',
  'L2',
  'R4',
  'R3',
  'R194',
  'L3',
  'L4',
  'L4',
  'L3',
  'L4',
  'R4',
  'R5',
  'L1',
  'L5',
  'L4',
  'R1',
  'L2',
  'R4',
  'L5',
  'L3',
  'R4',
  'L5',
  'L5',
  'R5',
  'R3',
  'R5',
  'L2',
  'L4',
  'R4',
  'L1',
  'R3',
  'R1',
  'L1',
  'L2',
  'R2',
  'R2',
  'L3',
  'R3',
  'R2',
  'R5',
  'R2',
  'R5',
  'L3',
  'R2',
  'L5',
  'R1',
  'R2',
  'R2',
  'L4',
  'L5',
  'L1',
  'L4',
  'R4',
  'R3',
  'R1',
  'R2',
  'L1',
  'L2',
  'R4',
  'R5',
  'L2',
  'R3',
  'L4',
  'L5',
  'L5',
  'L4',
  'R4',
  'L2',
  'R1',
  'R1',
  'L2',
  'L3',
  'L2',
  'R2',
  'L4',
  'R3',
  'R2',
  'L1',
  'L3',
  'L2',
  'L4',
  'L4',
  'R2',
  'L3',
  'L3',
  'R2',
  'L4',
  'L3',
  'R4',
  'R3',
  'L2',
  'L1',
  'L4',
  'R4',
  'R2',
  'L4',
  'L4',
  'L5',
  'L1',
  'R2',
  'L5',
  'L2',
  'L3',
  'R2',
  'L2'
];
class dir {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
  move(dx, dy, dist) {
    this.x += dx * Math.abs(dist);
    this.y += dy * Math.abs(dist);
  }
}
const north = new dir(0, 1);
const east = new dir(1, 0);
const south = new dir(0, -1);
const west = new dir(-1, 0);

const directions = [north, east, south, west];

function getDir(str) {
  return str[0];
}
function getDist(str) {
  return parseInt(str.split(/L|R/)[1]);
}
function turn(str) {
  return str === 'R' ? 1 : -1;
}

const pos = new dir(0, 0);
let d_index = 0;

const endPos = instructions.reduce((acc, curr) => {
  d_index = (d_index + turn(getDir(curr)) + 4) % 4;
  const dist = getDist(curr);
  acc.move(directions[d_index].x, directions[d_index].y, dist);
  return acc;
}, pos);

console.log(Math.abs(endPos.x) + Math.abs(endPos.y));
