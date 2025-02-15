type List = number[];

function parseLists(input: string, columns = 2): List[] {
  let out = Array.from({ length: columns }).map(() => []);

  const a = input
    .trim()
    .split(/\s+/)
    .forEach((num, index) => {
      out[index % columns].push(Number.parseInt(num));
    });

  return out;
}

function calculateTotalDisance(input: List[]): number {
  let acc = 0;
  let sortedLists = input.map((list) => list.sort());

  sortedLists.forEach((list, listIndex) => {
    list.forEach((number, index) => {
      const nextList = sortedLists[listIndex + 1];
      if (!nextList) {
        return;
      }
      acc += Math.abs(number - nextList[index]);
    });
  });

  return acc;
}

if (import.meta.vitest) {
  const { it, expect, describe } = import.meta.vitest;

  describe("day 01", () => {
    it("gets the correct total disance for the same input", () => {
      let sampleInput = `
      3   4
      4   3
      2   5
      1   3
      3   9
      3   3`;

      let [list1, list2] = parseLists(sampleInput);
      let disance = calculateTotalDisance([list1, list2]);

      expect(disance).toBe(11);
    });
  });
}
