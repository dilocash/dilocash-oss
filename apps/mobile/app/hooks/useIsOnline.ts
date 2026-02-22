import { useState, useEffect } from 'react';
import NetInfo from '@react-native-community/netinfo';

export const useIsOnline = () => {
  const [isOnline, setIsOnline] = useState(true);

  useEffect(() => {
    // Suscribirse a los cambios de red
    const unsubscribe = NetInfo.addEventListener(state => {
      setIsOnline(!!state.isConnected && !!state.isInternetReachable);
    });

    return () => unsubscribe();
  }, []);

  return isOnline;
};