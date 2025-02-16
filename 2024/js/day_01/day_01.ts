import { readFile } from "node:fs/promises";

type List = number[];

function parseLists(input: string, columns = 2): List[] {
  let out = Array.from({ length: columns }).map(() => []);

  input
    .trim()
    .split(/\s+/)
    .forEach((num, index) => {
      out[index % columns].push(Number.parseInt(num));
    });

  return out;
}

function calculateTotalDisance(input: List[]): number {
  return input
    .map((list) => list.sort())
    .reduce<number>((acc, list, listIndex, sortedLists) => {
      list.forEach((number, index) => {
        const nextList = sortedLists[listIndex + 1];

        if (nextList) {
          acc += Math.abs(number - nextList[index]);
        }
      });
      return acc;
    }, 0);
}

if (import.meta.vitest) {
  const { it, expect, describe } = import.meta.vitest;

  describe("day 01", () => {
    describe("part 1", () => {
      it("sample input", () => {
        let sampleInput = `
          3   4
          4   3
          2   5
          1   3
          3   9
          3   3
        `;

        let [list1, list2] = parseLists(sampleInput);
        let disance = calculateTotalDisance([list1, list2]);

        expect(disance).toBe(11);
      });

      it("full input", async () => {
        let filePath = new URL("./input.txt", import.meta.url);
        let input = await readFile(filePath, { encoding: "utf8" });

        let lists = parseLists(input, 2);
        let distance = calculateTotalDisance(lists);

        expect(distance).toEqual(1882714);
      });
    });
  });
}
