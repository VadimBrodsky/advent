export function day01() {
  console.log("hello");
}

if (import.meta.vitest) {
  const { it, expect } = import.meta.vitest;

  it("day 01", () => {
    expect(true).toBe(true);
  });
}
