/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */
import type { Metadata } from 'next';
import MainClient from './main/main-client';

export const metadata: Metadata = {
  title: 'Home',
  description: 'Manage your cash and expenses easily with dilocash.',
};

export default function Home() {
  return <MainClient />;
}
