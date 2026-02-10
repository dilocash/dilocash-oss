/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import type { NextConfig } from "next";

const path = require('path');

const nextConfig: NextConfig = {
  /* config options here */
  cacheComponents: true,
  turbopack : {
    root : path.join(__dirname, '../..')
  }
};

export default nextConfig;
