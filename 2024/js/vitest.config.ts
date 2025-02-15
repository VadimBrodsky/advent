import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    includeSource: ["day_**/*.{js,ts}"],
  },
  define: {
    "import.meta.vitest": "undefined",
  },
});
