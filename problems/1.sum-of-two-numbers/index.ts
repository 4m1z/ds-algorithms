// problem number one .
//give an array of integers and return the indeces of the
//two numbers that add up to the given target

type TSolution = [number, number] | null;
let testArray: number[] = [1, 2, 3, 7, 9];

function calSolution(arr: number[], targetValue: number): TSolution {
  for (let j = 0; j < arr?.length; ++j) {
    let BT = targetValue - arr[j];
    for (let i = 0; i < arr?.length; ++i) {
      if (arr[i] === BT) {
        return [j, i];
      }
    }
  }
  return null;
}

function calSolution3(arr: number[], targetValue: number): TSolution {
  let numArrMap: any = {};
  for (let i = 0; i < arr?.length; ++i) {
    let current = numArrMap[arr[i]];
    if (current >= 0) {
      return [current, i];
    } else {
      const numToFind = targetValue - arr[i];
      numArrMap[numToFind] = i;
    }
  }
  return null;
}

function calSolution4(arr: number[], targetValue: number): TSolution {
  const entrySet = new Set<number>();
  for (let i = 0; i < arr?.length; ++i) {
    entrySet.add(arr[i]);
    const numberThatWeWant = targetValue - arr[i];
    if (entrySet.has(numberThatWeWant)) {
      return [numberThatWeWant, arr[i]];
    }
  }
  return null;
}

const valuex = calSolution4([1, 2, 3, 4, 6, 9], 11);

function calSolution2(arr: number[], targetValue: number): TSolution {
  let numArrMap = new Map();
  for (let i = 0; i < arr?.length; ++i) {
    if (numArrMap.get(arr[i]) >= 0) {
      return [numArrMap.get(arr[i]), i];
    } else {
      numArrMap.set(targetValue - arr[i], i);
    }
  }
  return null;
}

// currently working on this !!!!
type TNewSolution<T> = [T, T, T] | null;

function calSolutionForThreeIndexes(
  arr: number[],
  T: number
): TNewSolution<number> {
  const firstMap = new Map();
  for (let i = 0; i < arr.length; ++i) {
    if (firstMap.get(arr[i] >= 0)) {
      return [firstMap.get(arr[i]), i, 2];
    }

    let sumOfTwoOthers = T - arr[i];
    firstMap.set(sumOfTwoOthers, i);
  }
  return null;
}

const value = calSolutionForThreeIndexes([1, 2, 3, 4, 6, 9], 11);
console.log(value);
// calculate 3 nbumber be the Goal

////// tests
console.log("solution 1 ....");
console.log(calSolution([1, 2, 3, 4, 6, 9], 11));
console.log(calSolution([1, 2, 3, 4, 5], 25));
console.log(calSolution([1, 6], 7));
console.log(calSolution([1, 6], 11));
console.log(calSolution([5], 5));
console.log(calSolution([], 5));
console.log("solution 2 ....");
console.log(calSolution2([1, 2, 3, 4, 6, 9], 11));
console.log(calSolution2([1, 2, 3, 4, 5], 25));
console.log(calSolution2([1, 6], 7));
console.log(calSolution2([1, 6], 11));
console.log(calSolution2([5], 5));
console.log(calSolution2([], 5));