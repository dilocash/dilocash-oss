/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import { useState, useEffect } from 'react';
// TODO test this
export const useIsOnline = () => {

  // Check if window is defined (client-side)
  const maybeWindow = typeof window === "undefined" ? null : window;
  const [isOnline, setIsOnline] = useState(
    maybeWindow ? maybeWindow.navigator.onLine : true
  );

  useEffect(() => {
    // This effect runs only on the client side
    if (!maybeWindow) return;
    const handleStatusChange = () => setIsOnline(maybeWindow.navigator.onLine);

    maybeWindow.addEventListener('online', handleStatusChange);
    maybeWindow.addEventListener('offline', handleStatusChange);

    return () => {
      maybeWindow.removeEventListener('online', handleStatusChange);
      maybeWindow.removeEventListener('offline', handleStatusChange);
    };
  }, [maybeWindow]);

  return isOnline;
};