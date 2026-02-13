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
  turbopack : {
    root : path.join(__dirname, '../..')
  },
  transpilePackages: [
    "@dilocash/ui",
    "@gluestack-ui/core",
    "@gluestack-ui/utils",
    "@gluestack/ui-next-adapter",
    "react-native-web",
  ],
};

export default withGluestackUI(nextConfig);