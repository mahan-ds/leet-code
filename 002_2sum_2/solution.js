let listOne = [1, 4, 7, 10];
let listTwo = [2, 3, 6, 11 ,17];

let max_length =
  listOne.length > listTwo.length ? listOne.length : listTwo.length;
let final = [];
let i = 0;
let j = 0;
let middle = (listOne.length + listTwo.length ) / 2
console.log(middle)
while (i < max_length && j < max_length) {
  if (listOne[i] < listTwo[j]) {

    if (listOne[i]) {
      final.push(listOne[i]);
      i++;
    } else {
      final.push(listTwo[j]);
      j++;
      break;
    }
  } else {
    if (listTwo[j]) {
      final.push(listTwo[j]);
      j++;
    } else {
      final.push(listOne[i]);
      i++;
      break;
    }
  }
}

console.log(final);
