/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/components/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        mono: ["JetBrains Mono", "monospace"],
      },
    },
  },
  plugins: [],
};
