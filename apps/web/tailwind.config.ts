import sharedConfig from "@dilocash/ui/tailwind.config";
import type { Config } from "tailwindcss";

const config: Config = {
  ...sharedConfig,
  content: [
    ...sharedConfig.content as string[],
    "./app/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "../../packages/ui/components/**/*.{js,ts,jsx,tsx}",
    "../../packages/ui/src/**/*.{js,ts,jsx,tsx}"
  ],
};

export default config;