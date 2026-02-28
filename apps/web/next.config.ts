/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import type { NextConfig } from "next";
import { withGluestackUI } from '@gluestack/ui-next-adapter';
const path = require('path');

const nextConfig: NextConfig = {
  /* config options here */
  cacheComponents: true,
  reactStrictMode: true,

  transpilePackages: [
    "@dilocash/ui",
    "@dilocash/database",
    "@gluestack-ui/core",
    "@gluestack-ui/utils",
    "@gluestack/ui-next-adapter",
    "react-native-web",
    "nativewind",
    "react-native-css-interop",
  ],
};

export default withGluestackUI(nextConfig);