/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import type { NextConfig } from "next";
import { withGluestackUI } from '@gluestack/ui-next-adapter';

const nextConfig: NextConfig = {
  reactStrictMode: true,
  transpilePackages: [
    "@dilocash/ui",
    "@dilocash/database",
    "@gluestack-ui/core",
    "@gluestack-ui/utils",
    "@gluestack/ui-next-adapter",
    "react-native-css-interop"
  ],
  webpack: (config) => {
    config.resolve.alias = {
      ...(config.resolve.alias || {}),
      'react-native$': 'react-native-web',
    };
    config.resolve.extensions = [
      '.web.js',
      '.web.jsx',
      '.web.ts',
      '.web.tsx',
      ...config.resolve.extensions,
    ];
    return config;
  },
};

export default withGluestackUI(nextConfig);