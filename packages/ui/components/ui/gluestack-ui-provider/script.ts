/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

export const script = (mode: string) => {
  const documentElement = document.documentElement;

  function getSystemColorMode() {
    return window.matchMedia('(prefers-color-scheme: dark)').matches
      ? 'dark'
      : 'light';
  }

  try {
    const isSystem = mode === 'system';
    const theme = isSystem ? getSystemColorMode() : mode;
    documentElement.classList.remove(theme === 'light' ? 'dark' : 'light');
    documentElement.classList.add(theme);
    documentElement.style.colorScheme = theme;
  } catch (e) {
    console.error(e);
  }
};
