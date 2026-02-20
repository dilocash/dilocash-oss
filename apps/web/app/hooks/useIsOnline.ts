import { useState, useEffect } from 'react';
// TODO test this
export const useIsOnline = () => {
  const [isOnline, setIsOnline] = useState(
    typeof window !== 'undefined' ? window.navigator.onLine : true
  );

  useEffect(() => {
    const handleStatusChange = () => setIsOnline(window.navigator.onLine);

    window.addEventListener('online', handleStatusChange);
    window.addEventListener('offline', handleStatusChange);

    return () => {
      window.removeEventListener('online', handleStatusChange);
      window.removeEventListener('offline', handleStatusChange);
    };
  }, []);

  return isOnline;
};