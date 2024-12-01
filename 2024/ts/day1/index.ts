/**
 * @license MIT
 * @author gdjohn4s
 * @created 2024-12-01
 */

import { readFileSync } from "fs";

function readFileLineByLine(filePath: string): string[] {
  const fileContent = readFileSync(filePath, "utf-8");
  return fileContent.split(/\r?\n/);
}

const lines = readFileLineByLine("./input.txt");
const mapLists = () => {
  const [firstList, secondList] = lines.slice(0, -1).reduce(
    (acc: [number[], number[]], line) => {
      const [first, second] = line.split("  ");
      acc[0].push(Number(first));
      acc[1].push(Number(second));
      return acc;
    },
    [[], []],
  );

  return [firstList, secondList];
};

const firstResult = () => {
  const [firstList, secondList] = mapLists();

  const sortedFirst = firstList.sort();
  const sortedSecond = secondList.sort();

  const distance = sortedFirst.reduce(
    (sum, num, i) => sum + Math.abs(num - sortedSecond[i]),
    0,
  );

  console.log("First: ", distance);
};

const secondResult = () => {
  const [firstList, secondList] = mapLists();

  const sortedFirst = firstList.sort();
  const sortedSecond = secondList.sort();

  const distance = sortedFirst.reduce((sum, num) => {
    const count = sortedSecond.filter((el) => el === num).length;
    return sum + num * count;
  }, 0);

  console.log("Second: ", distance);
};

firstResult();
secondResult();
